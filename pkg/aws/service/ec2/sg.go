package ec2

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func GetSecurityGroupIds(client *ec2.Client) []string {
	var l []string
	resp, err := client.DescribeSecurityGroups(context.TODO(), &ec2.DescribeSecurityGroupsInput{
		MaxResults: aws.Int32(100),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range resp.SecurityGroups {
		l = append(l, *s.GroupId)
	}

	return l
}

// 		var direction string = "Ingress"
// 		for _, v := range resp.SecurityGroupRules {
// 			if v.CidrIpv4 != nil {
// 				if *v.IsEgress {
// 					direction = "Egress"
// 				}
// 				if utils.Contains(args, *v.CidrIpv4) == true {
// 					table.Append([]string{"SecurityGroup", direction, *s.GroupName, *v.CidrIpv4})
// 				}
// 			}
// 		}
// 	}
// }
