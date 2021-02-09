package main

import (
	"fmt"

	sc "github.com/blinchik/aws/secretmanager"
)

func main() {

	// create secret
	sc.CreateSecret("brain-test3", "secretValue", "Description")

	// read secret
	v := sc.GetSecret("brain-test3")
	fmt.Println(v)

	// delete secret
	sc.DeleteSecret("brain-test3")

}
