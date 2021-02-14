package main

import (
	"fmt"

	cf "github.com/blinchik/aws/config"
	de "github.com/blinchik/aws/services/ec2"
	sc "github.com/blinchik/aws/services/secretmanager"
)

func main() {

	cfg, region := cf.CreateConfigFromEC2Role()

	// create secret
	sc.CreateSecret("brain-test3", "secretValue", "Description", cfg, region)

	// read secret
	v := sc.GetSecret("brain-test3", cfg, region)
	fmt.Println(v)

	// delete secret
	sc.DeleteSecret("brain-test3", cfg, region)

	desOutput := de.DescribeInstanceByTag("Function", "brain", "Environment", "dev", cfg, region)

	tagValue := de.GetTagValue("Name", desOutput)

	fmt.Println(tagValue)

}
