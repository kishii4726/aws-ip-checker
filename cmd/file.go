package cmd

import (
	"aws-ip-checker/pkg/aws/config"
	servicealb "aws-ip-checker/pkg/aws/service/alb"
	serviceec2 "aws-ip-checker/pkg/aws/service/ec2"
	servicests "aws-ip-checker/pkg/aws/service/sts"
	servicewaf "aws-ip-checker/pkg/aws/service/waf"
	"aws-ip-checker/pkg/file"
	"aws-ip-checker/pkg/table"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Args:  cobra.ExactArgs(1),
	Short: "Check if the IP address listed in the csv file is included",
	Long: `Specify the path to the csv file with IP addresses as an argument.
Only one file can be specified.
e.g. sample.csv`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			cfg           aws.Config                     = config.LoadConfig("")
			us_east_1_cfg aws.Config                     = config.LoadConfig("us-east-1")
			c_ec2         *ec2.Client                    = ec2.NewFromConfig(cfg)
			c_wafv2       *wafv2.Client                  = wafv2.NewFromConfig(cfg)
			u_c_wafv2     *wafv2.Client                  = wafv2.NewFromConfig(us_east_1_cfg)
			c_wafv1_r     *wafregional.Client            = wafregional.NewFromConfig(cfg)
			c_wafv1_c     *waf.Client                    = waf.NewFromConfig(us_east_1_cfg)
			c_elbv2       *elasticloadbalancingv2.Client = elasticloadbalancingv2.NewFromConfig(cfg)
			table         *tablewriter.Table             = table.SetTable()
			ipaddresses   [][]string                     = file.ReadCsv(args[0])
		)

		// SecurityGroup
		for _, s := range serviceec2.GetSecurityGroupIds(c_ec2) {
			for _, i := range serviceec2.CheckContainIpAddress(c_ec2, serviceec2.GetSecurityGroupRules(c_ec2, s[0]), ipaddresses[0]) {
				table.Append([]string{"SecurityGroup", i[0], s[1] + "(" + s[0] + ")", i[1]})
			}
		}

		// WAFv2
		// Regional
		for _, v := range servicewaf.V2GetIpSets(c_wafv2, "REGIONAL").IPSets {
			for _, w := range servicewaf.V2CheckContainIpAddress(c_wafv2, *&v.Id, *&v.Name, "REGIONAL", ipaddresses[0]) {
				table.Append([]string{"WAF v2", "IPSet, Regional", w[0], w[1]})
			}
		}
		// Cloudfront
		for _, v := range servicewaf.V2GetIpSets(u_c_wafv2, "CLOUDFRONT").IPSets {
			for _, w := range servicewaf.V2CheckContainIpAddress(u_c_wafv2, *&v.Id, *&v.Name, "CLOUDFRONT", ipaddresses[0]) {
				table.Append([]string{"WAF v2", "IPSet, CloudFront", w[0], w[1]})
			}
		}

		// WAFv1
		// Regional
		for _, v := range servicewaf.V1RegionalGetIpSets(c_wafv1_r).IPSets {
			for _, w := range servicewaf.V1RegionalCheckContainIpAddress(c_wafv1_r, *&v.IPSetId, *&v.Name, ipaddresses[0]) {
				table.Append([]string{"WAF Classic", "IPSet, Regional", w[0], w[1]})
			}
		}
		// Cloudfront
		for _, v := range servicewaf.V1CloudFrontGetIpSets(c_wafv1_c).IPSets {
			for _, w := range servicewaf.V1CloudFrontCheckContainIpAddress(c_wafv1_c, *&v.IPSetId, *&v.Name, ipaddresses[0]) {
				table.Append([]string{"WAF Classic", "IPSet, CloudFront", w[0], w[1]})
			}
		}

		// alb
		for _, lb := range servicealb.GetLoadBalancers(c_elbv2).LoadBalancers {
			for _, li := range servicealb.DescribeLoadBalancerListeners(c_elbv2, lb.LoadBalancerArn).Listeners {
				for _, i := range servicealb.CheckContainIpAddress(c_elbv2, li.ListenerArn, ipaddresses[0]) {
					table.Append([]string{"ALB", "Listener, Port: " + strconv.Itoa(int(*li.Port)), *lb.LoadBalancerName, i})
				}
			}
		}

		// accountid
		sts := sts.NewFromConfig(cfg)
		fmt.Println("AccountId: " + *servicests.GetAccountId(sts))

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
}
