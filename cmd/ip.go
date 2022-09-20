package cmd

import (
	"aws-ip-checker/pkg/aws/config"
	servicealb "aws-ip-checker/pkg/aws/service/alb"
	serviceec2 "aws-ip-checker/pkg/aws/service/ec2"
	servicests "aws-ip-checker/pkg/aws/service/sts"
	servicewaf "aws-ip-checker/pkg/aws/service/waf"
	"aws-ip-checker/pkg/table"
	"fmt"

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

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Args:  cobra.MinimumNArgs(1),
	Short: "Check if the specified IP is included",
	Long: `Specify the IP address you wish to check in the argument in CIDR format.
Multiple IP addresses can be specified.
e.g. 192.168.0.0/32 192.168.0.0/24`,
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
		)

		// SecurityGroup
		for _, s := range serviceec2.GetSecurityGroupIds(c_ec2) {
			for _, i := range serviceec2.CheckContainIpAddress(c_ec2, serviceec2.GetSecurityGroupRules(c_ec2, s[0]), args) {
				table.Append([]string{"SecurityGroup", i[0] + ", Port: " + i[2] + " - " + i[3], s[1], s[0], i[1]})
			}
		}

		// WAFv2
		// Regional
		for _, v := range servicewaf.V2GetIpSets(c_wafv2, "REGIONAL") {
			for _, w := range servicewaf.V2CheckContainIpAddress(c_wafv2, &v[0], &v[1], "REGIONAL", args) {
				table.Append([]string{"WAF v2", "IPSet, Regional", w[0], v[0], w[1]})
			}
		}
		// Cloudfront
		for _, v := range servicewaf.V2GetIpSets(u_c_wafv2, "CLOUDFRONT") {
			for _, w := range servicewaf.V2CheckContainIpAddress(u_c_wafv2, &v[0], &v[1], "CLOUDFRONT", args) {
				table.Append([]string{"WAF v2", "IPSet, CloudFront", w[0], v[0], w[1]})
			}
		}

		// WAFv1
		// Regional
		for _, v := range servicewaf.V1RegionalGetIpSets(c_wafv1_r) {
			for _, w := range servicewaf.V1RegionalCheckContainIpAddress(c_wafv1_r, &v[0], &v[1], args) {
				table.Append([]string{"WAF Classic", "IPSet, Regional", w[0], v[0], w[1]})
			}
		}
		// Cloudfront
		for _, v := range servicewaf.V1CloudFrontGetIpSets(c_wafv1_c) {
			for _, w := range servicewaf.V1CloudFrontCheckContainIpAddress(c_wafv1_c, &v[0], &v[1], args) {
				table.Append([]string{"WAF Classic", "IPSet, CloudFront", w[0], v[0], w[1]})
			}
		}

		// alb
		for _, lb := range servicealb.GetLoadBalancers(c_elbv2) {
			for _, li := range servicealb.DescribeLoadBalancerListeners(c_elbv2, &lb[1]) {
				for _, i := range servicealb.CheckContainIpAddress(c_elbv2, &li[0], args) {
					table.Append([]string{"ALB", "Listener, Port: " + li[1], lb[0], li[0], i})
				}
			}
		}

		// accountid
		sts := sts.NewFromConfig(cfg)
		fmt.Println("AccountId: " + *servicests.GetAccountId(sts))

		table.SetRowLine(true)
		table.SetAutoMergeCellsByColumnIndex([]int{2, 3})
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
}
