package waf

import (
	"aws-ip-checker/pkg/utils"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
)

func V2GetIpSets(client *wafv2.Client, scope types.Scope) *wafv2.ListIPSetsOutput {
	resp, err := client.ListIPSets(context.TODO(), &wafv2.ListIPSetsInput{
		Scope: scope,
		Limit: aws.Int32(100),
	})
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func V2CheckContainIpAddress(client *wafv2.Client, id *string, name *string, scope types.Scope, ipaddresses []string) [][]string {
	var d [][]string
	resp, err := client.GetIPSet(context.TODO(), &wafv2.GetIPSetInput{
		Id:    id,
		Name:  name,
		Scope: scope,
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range resp.IPSet.Addresses {
		if utils.Contains(ipaddresses, ip) == true {
			d = append(d, []string{*name, ip})
		}
	}

	return d
}
