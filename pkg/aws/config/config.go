package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func LoadConfig(region string) aws.Config {

	resp, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return resp
}

// func UsEast1LoadConfig() aws.Config {

// 	resp, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
// 	if err != nil {
// 		log.Fatalf("unable to load SDK config, %v", err)
// 	}

// 	return resp
// }
