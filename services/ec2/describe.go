package ec2

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

//DescribeInstanceByTag describe ec2 instance by tag
func DescribeInstanceByTag(tagKey, tagValue, tagKeySecond, tagValueSecond string, cfg *aws.Config, region string) *ec2.DescribeInstancesOutput {

	client := ec2.NewFromConfig(*cfg)

	input := &ec2.DescribeInstancesInput{
		Filters: []types.Filter{

			{
				Name: aws.String(fmt.Sprintf("tag:%s", tagKeySecond)),
				Values: []string{
					tagValueSecond,
				},
			},

			{
				Name: aws.String(fmt.Sprintf("tag:%s", tagKey)),
				Values: []string{
					tagValue,
				},
			},

			{
				Name: aws.String("instance-state-name"),
				Values: []string{
					"running",
				},
			},
		},
	}

	opts := func(o *ec2.Options) {
		o.Region = region
	}

	client.DescribeInstances(context.TODO(), input, opts)

	result, err := client.DescribeInstances(context.TODO(), input, opts)

	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		log.Fatal(err)
		return result
	}

	return result

}

//GetTagValue get the value of a tag from *ec2.DescribeInstancesOutput
func GetTagValue(tagKey string, ec2DecribeInstancesOutput *ec2.DescribeInstancesOutput) string {

	var tagValue string

	tagList := ec2DecribeInstancesOutput.Reservations[0].Instances[0].Tags

	for _, tag := range tagList {
		if *tag.Key == tagKey {
			tagValue = *tag.Value
		}

	}

	return tagValue

}
