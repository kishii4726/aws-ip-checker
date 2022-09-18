package waf

import (
	"aws-ip-checker/pkg/utils"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func V1RegionalGetIpSets(client *wafregional.Client) *wafregional.ListIPSetsOutput {
	resp, err := client.ListIPSets(context.TODO(), &wafregional.ListIPSetsInput{
		Limit: 100,
	})
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func V1CloudFrontGetIpSets(client *waf.Client) *waf.ListIPSetsOutput {
	resp, err := client.ListIPSets(context.TODO(), &waf.ListIPSetsInput{
		Limit: 100,
	})
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func V1RegionalCheckContainIpAddress(client *wafregional.Client, id *string, name *string, ipaddresses []string) [][]string {
	var d [][]string
	resp, err := client.GetIPSet(context.TODO(), &wafregional.GetIPSetInput{
		IPSetId: id,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range resp.IPSet.IPSetDescriptors {
		if utils.Contains(ipaddresses, *ip.Value) == true {
			d = append(d, []string{*name, *ip.Value})
		}
	}

	return d
}

func V1CloudFrontCheckContainIpAddress(client *waf.Client, id *string, name *string, ipaddresses []string) [][]string {
	var d [][]string
	resp, err := client.GetIPSet(context.TODO(), &waf.GetIPSetInput{
		IPSetId: id,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range resp.IPSet.IPSetDescriptors {
		if utils.Contains(ipaddresses, *ip.Value) == true {
			d = append(d, []string{*name, *ip.Value})
		}
	}

	return d
}
