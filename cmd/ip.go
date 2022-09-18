/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"aws-ip-checker/pkg/config"
	"aws-ip-checker/pkg/table"
	"aws-ip-checker/pkg/utils"
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Args:  cobra.MinimumNArgs(1),
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		table := table.SetTable()

		// SecurityGroup
		c_ec2 := ec2.NewFromConfig(cfg)
		resp, err := c_ec2.DescribeSecurityGroups(context.TODO(), &ec2.DescribeSecurityGroupsInput{
			MaxResults: aws.Int32(100),
		})
		if err != nil {
			log.Fatal(err)
		}
		for _, s := range resp.SecurityGroups {
			resp, err := c_ec2.DescribeSecurityGroupRules(context.TODO(), &ec2.DescribeSecurityGroupRulesInput{
				Filters: []types.Filter{
					{
						Name:   aws.String("group-id"),
						Values: []string{*s.GroupId},
					},
				},
			})
			if err != nil {
				log.Fatal(err)
			}

			var direction string = "Ingress"
			for _, v := range resp.SecurityGroupRules {
				if v.CidrIpv4 != nil {
					if *v.IsEgress {
						direction = "Egress"
					}
					if utils.Contains(args, *v.CidrIpv4) == true {
						table.Append([]string{"SecurityGroup", direction, *s.GroupName, *v.CidrIpv4})
					}
				}
			}
		}

		//WAFv2
		c_wafv2 := wafv2.NewFromConfig(cfg)
		//reginal
		resp2, err := c_wafv2.ListIPSets(context.TODO(), &wafv2.ListIPSetsInput{
			Scope: "REGIONAL",
			Limit: aws.Int32(100),
		})
		if err != nil {
			log.Fatal(err)
		}
		// regional
		for _, v := range resp2.IPSets {
			resp, err := c_wafv2.GetIPSet(context.TODO(), &wafv2.GetIPSetInput{
				Id:    *&v.Id,
				Name:  *&v.Name,
				Scope: "REGIONAL",
			})
			if err != nil {
				log.Fatal(err)
			}
			for _, ip := range resp.IPSet.Addresses {
				if utils.Contains(args, ip) == true {
					table.Append([]string{"WAFv2", "IPSet, Regional", *v.Name, ip})
				}
			}
		}

		// cloudfront
		us_east_1_cfg := config.UsEast1LoadConfig()
		u_c_wafv2 := wafv2.NewFromConfig(us_east_1_cfg)

		resp3, err := u_c_wafv2.ListIPSets(context.TODO(), &wafv2.ListIPSetsInput{
			Scope: "CLOUDFRONT",
			Limit: aws.Int32(100),
		})
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range resp3.IPSets {
			resp, err := u_c_wafv2.GetIPSet(context.TODO(), &wafv2.GetIPSetInput{
				Id:    *&v.Id,
				Name:  *&v.Name,
				Scope: "CLOUDFRONT",
			})
			if err != nil {
				log.Fatal(err)
			}
			for _, ip := range resp.IPSet.Addresses {
				if utils.Contains(args, ip) == true {
					table.Append([]string{"WAFv2", "IPSet, CloudFront", *v.Name, ip})
				}
			}
		}

		// alb
		c_elbv2 := elasticloadbalancingv2.NewFromConfig(cfg)
		resp4, err := c_elbv2.DescribeLoadBalancers(context.TODO(), &elasticloadbalancingv2.DescribeLoadBalancersInput{
			PageSize: aws.Int32(100),
		})
		if err != nil {
			log.Fatal(err)
		}
		for _, lb := range resp4.LoadBalancers {
			resp, err := c_elbv2.DescribeListeners(context.TODO(), &elasticloadbalancingv2.DescribeListenersInput{
				LoadBalancerArn: &*lb.LoadBalancerArn,
			})
			if err != nil {
				log.Fatal(err)
			}
			for _, li := range resp.Listeners {
				respx, err := c_elbv2.DescribeRules(context.TODO(), &elasticloadbalancingv2.DescribeRulesInput{
					ListenerArn: li.ListenerArn,
				})
				if err != nil {
					log.Fatal(err)
				}
				for _, v := range respx.Rules {
					for _, v := range v.Conditions {
						if v.SourceIpConfig != nil {
							for _, v := range *&v.SourceIpConfig.Values {
								if utils.Contains(args, v) == true {
									table.Append([]string{"ALB", "Listener, " + "Port: " + strconv.Itoa(int(*li.Port)), *lb.LoadBalancerName, v})
								}
							}
						}
					}
				}
			}
		}
		// accountid
		sts := sts.NewFromConfig(cfg)
		resp5, err := sts.GetCallerIdentity(context.TODO(), nil)
		fmt.Println("AccountId: " + *resp5.Account)
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
