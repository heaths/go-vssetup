package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/heaths/go-vssetup"
)

func main() {
	var all bool
	flag.BoolVar(&all, "all", false, "Finds all instances even if they are incomplete and may not launch.")

	instances, err := vssetup.Instances(all)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	for _, instance := range instances {
		if instanceId, err := instance.InstanceId(); err == nil {
			fmt.Println("InstanceId =", instanceId)
		}

		if installDate, err := instance.InstallDate(); err == nil {
			fmt.Println("InstallDate =", installDate)
		}

		if installationName, err := instance.InstallationName(); err == nil {
			fmt.Println("InstallationName =", installationName)
		}

		if installationPath, err := instance.InstallationPath(); err == nil {
			fmt.Println("InstallationPath =", installationPath)
		}

		fmt.Println()
	}
}
