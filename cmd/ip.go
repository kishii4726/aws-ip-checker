/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"aws-ip-checker/pkg/config"
	"aws-ip-checker/pkg/table"
	"aws-ip-checker/pkg/utils"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
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

		//SecurityGroup
		// c_ec2 := ec2.NewFromConfig(cfg)
		// resp, err := c_ec2.DescribeSecurityGroups(context.TODO(), &ec2.DescribeSecurityGroupsInput{
		// 	MaxResults: aws.Int32(100),
		// })
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// for _, v := range resp.SecurityGroups {
		// 	fmt.Println(*v.GroupName + "(" + *v.GroupId + ")")
		// 	resp, err := c_ec2.DescribeSecurityGroupRules(context.TODO(), &ec2.DescribeSecurityGroupRulesInput{
		// 		Filters: []types.Filter{
		// 			{
		// 				Name:   aws.String("group-id"),
		// 				Values: []string{*v.GroupId},
		// 			},
		// 		},
		// 	})
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	for _, v := range resp.SecurityGroupRules {
		// 		fmt.Println(*v.SecurityGroupRuleId)
		// 		// fmt.Printf("%T\n", *v.GroupId)
		// 		// fmt.Println(*v.CidrIpv4)
		// 	}
		// }

		//WAFv2
		c_wafv2 := wafv2.NewFromConfig(cfg)
		//reginal
		resp, err := c_wafv2.ListIPSets(context.TODO(), &wafv2.ListIPSetsInput{
			// "CLOUDFRONT"
			Scope: "REGIONAL",
			Limit: aws.Int32(100),
		})
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range resp.IPSets {
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
					table.Append([]string{"WAFv2(Reginal)", *v.Name, ip})
				}
			}
		}
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
