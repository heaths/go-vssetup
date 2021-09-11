package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/heaths/go-vssetup"
	"github.com/heaths/go-vssetup/internal/formatting"
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

		formatting.PrintInstance(os.Stdout, &instance)
	}
}
