package sts

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func GetAccountId(client *sts.Client) *string {
	resp, err := client.GetCallerIdentity(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Account
}
