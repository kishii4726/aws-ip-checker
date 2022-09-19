package waf

import (
	"aws-ip-checker/pkg/utils"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func V1RegionalGetIpSets(client *wafregional.Client) [][]string {
	var d [][]string
	req_params := &wafregional.ListIPSetsInput{
		Limit: 100,
	}

	for {
		resp, err := client.ListIPSets(context.TODO(), req_params)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range resp.IPSets {
			d = append(d, []string{*v.IPSetId, *v.Name})
		}
		if resp.NextMarker != nil {
			req_params.NextMarker = resp.NextMarker
		} else {
			break
		}
	}

	return d
}

func V1CloudFrontGetIpSets(client *waf.Client) [][]string {
	var d [][]string
	req_params := &waf.ListIPSetsInput{
		Limit: 100,
	}

	for {
		resp, err := client.ListIPSets(context.TODO(), req_params)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range resp.IPSets {
			d = append(d, []string{*v.IPSetId, *v.Name})
		}
		if resp.NextMarker != nil {
			req_params.NextMarker = resp.NextMarker
		} else {
			break
		}
	}

	return d
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
