package ec2

import (
	"aws-ip-checker/pkg/utils"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func GetSecurityGroupIds(client *ec2.Client) [][]string {
	var d [][]string

	req_params := &ec2.DescribeSecurityGroupsInput{
		MaxResults: aws.Int32(100),
		NextToken:  aws.String(""),
	}

	for {
		resp, err := client.DescribeSecurityGroups(context.TODO(), req_params)
		if err != nil {
			log.Fatal(err)
		}
		if len(resp.SecurityGroups) == 0 {
			log.Println("SecurityGroup does not exist.")
			break
		}
		for _, v := range resp.SecurityGroups {
			d = append(d, []string{*v.GroupId, *v.GroupName})
		}
		if resp.NextToken != nil {
			req_params.NextToken = resp.NextToken
		} else {
			break
		}
	}

	return d
}

func GetSecurityGroupRules(client *ec2.Client, security_group_id string) *ec2.DescribeSecurityGroupRulesOutput {
	req_params := &ec2.DescribeSecurityGroupRulesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("group-id"),
				Values: []string{security_group_id},
			},
		},
		MaxResults: aws.Int32(100),
	}

	resp, err := client.DescribeSecurityGroupRules(context.TODO(), req_params)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func CheckContainIpAddress(client *ec2.Client, security_group_rules *ec2.DescribeSecurityGroupRulesOutput, ipaddresses []string) [][]string {
	var d [][]string
	var direction string
	for _, v := range security_group_rules.SecurityGroupRules {
		if v.CidrIpv4 != nil {
			if *v.IsEgress {
				direction = "Egress"
			} else {
				direction = "Ingress"
			}
			if utils.Contains(ipaddresses, *v.CidrIpv4) == true {
				d = append(d, []string{direction, *v.CidrIpv4})
			}
		}
	}
	return d
}
