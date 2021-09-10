package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/heaths/go-vssetup"
)

func main() {
	var (
		all  bool
		path string
	)

	flag.BoolVar(&all, "all", false, "Finds all instances even if they are incomplete and may not launch.")
	flag.StringVar(&path, "path", "", "Gets an instance for the given path, if any defined for that path.")

	var instances []vssetup.Instance
	if path != "" {
		instance, err := vssetup.InstanceForPath(path)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		instances = []vssetup.Instance{*instance}
	} else {
		var err error
		instances, err = vssetup.Instances(all)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

	for i, instance := range instances {
		if i > 0 {
			fmt.Println()
		}

		if instanceID, err := instance.InstanceID(); err == nil {
			fmt.Println("InstanceID =", instanceID)
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
	}
}
