package secretmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

//CreateSecret create secret with default option and region
func CreateSecret(name, SecretString, Description string) {

	svc := secretsmanager.NewFromConfig(*Cfg)

	opts := func(o *secretsmanager.Options) {
		o.Region = Region
	}

	input := &secretsmanager.CreateSecretInput{
		Name:         aws.String(name),
		Description:  aws.String(Description),
		SecretString: aws.String(SecretString),
	}

	result, err := svc.CreateSecret(context.TODO(), input, opts)

	if err != nil {

		log.Fatal(err)
	}

	output, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))

}

//GetSecret get secret with the given name and default AWSCURRENT
func GetSecret(secret string) string {

	svc := secretsmanager.NewFromConfig(*Cfg)

	opts := func(o *secretsmanager.Options) {
		o.Region = Region
	}

	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secret),
	}

	result, err := svc.GetSecretValue(context.TODO(), input, opts)

	if err != nil {

		log.Fatal(err)
	}

	output, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))

	return *result.SecretString

}

//DeleteSecret delete secret with given name. (force delete without recovery)
func DeleteSecret(secret string) {

	svc := secretsmanager.NewFromConfig(*Cfg)

	opts := func(o *secretsmanager.Options) {
		o.Region = Region
	}

	input := &secretsmanager.DeleteSecretInput{
		SecretId:                   aws.String(secret),
		ForceDeleteWithoutRecovery: true,
	}

	result, err := svc.DeleteSecret(context.TODO(), input, opts)

	if err != nil {

		log.Fatal(err)
	}

	output, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))

}
