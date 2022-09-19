package alb

import (
	"aws-ip-checker/pkg/utils"
	"context"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func GetLoadBalancers(client *elasticloadbalancingv2.Client) [][]string {
	var d [][]string
	req_params := &elasticloadbalancingv2.DescribeLoadBalancersInput{
		PageSize: aws.Int32(100),
	}

	for {
		resp, err := client.DescribeLoadBalancers(context.TODO(), req_params)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range resp.LoadBalancers {
			d = append(d, []string{*v.LoadBalancerName, *v.LoadBalancerArn})
		}
		if resp.NextMarker != nil {
			req_params.Marker = resp.NextMarker
		} else {
			break
		}
	}

	return d
}

func DescribeLoadBalancerListeners(client *elasticloadbalancingv2.Client, load_balancer_arn *string) [][]string {
	var d [][]string
	req_params := &elasticloadbalancingv2.DescribeListenersInput{
		LoadBalancerArn: load_balancer_arn,
	}

	for {
		resp, err := client.DescribeListeners(context.TODO(), req_params)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range resp.Listeners {
			d = append(d, []string{*v.ListenerArn, strconv.Itoa(int(*v.Port))})
		}
		if resp.NextMarker != nil {
			req_params.Marker = resp.NextMarker
		} else {
			break
		}
	}

	return d
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
