package secretmanager

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
)

//CreateConfigFromEC2Role loads default config from EC2 and also extracts the aws region.
func CreateConfigFromEC2Role() (*aws.Config, string) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := imds.NewFromConfig(cfg)

	region, err := client.GetRegion(context.TODO(), &imds.GetRegionInput{})
	if err != nil {
		log.Printf("Unable to retrieve the region from the EC2 instance %v\n", err)
	}

	return &cfg, region.Region

}

//Cfg default config
var Cfg *aws.Config

//Region defualt aws region
var Region string

func init() {
	Cfg, Region = CreateConfigFromEC2Role()
}
