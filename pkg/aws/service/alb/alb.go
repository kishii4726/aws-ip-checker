package alb

import (
	"aws-ip-checker/pkg/utils"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func GetLoadBalancers(client *elasticloadbalancingv2.Client) *elasticloadbalancingv2.DescribeLoadBalancersOutput {
	resp, err := client.DescribeLoadBalancers(context.TODO(), &elasticloadbalancingv2.DescribeLoadBalancersInput{
		PageSize: aws.Int32(100),
	})
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func DescribeLoadBalancerListeners(client *elasticloadbalancingv2.Client, load_balancer_arn *string) *elasticloadbalancingv2.DescribeListenersOutput {
	resp, err := client.DescribeListeners(context.TODO(), &elasticloadbalancingv2.DescribeListenersInput{
		LoadBalancerArn: load_balancer_arn,
	})
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func CheckContainIpAddress(client *elasticloadbalancingv2.Client, listen_arn *string, ipaddresses []string) []string {
	var l []string
	resp, err := client.DescribeRules(context.TODO(), &elasticloadbalancingv2.DescribeRulesInput{
		ListenerArn: listen_arn,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range resp.Rules {
		for _, c := range r.Conditions {
			if c.SourceIpConfig != nil {
				for _, v := range *&c.SourceIpConfig.Values {
					if utils.Contains(ipaddresses, v) == true {
						l = append(l, v)
					}
				}
			}
		}
	}
	return l
}
