package main

import (
	"fmt"
	"os"

	"github.com/heaths/go-vssetup"
	"github.com/heaths/go-vssetup/internal/formatting"
	"github.com/spf13/cobra"
)

type options struct {
	all  bool
	path string
}

func main() {
	opts := options{}
	root := cobra.Command{
		Use: "Locates instances of Visual Studio",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			return run(&opts)
		},
	}

	root.Flags().BoolVar(&opts.all, "all", false, "Finds all instances even if they are incomplete and may not launch.")
	root.Flags().StringVar(&opts.path, "path", "", "Gets an instance for the given path, if any defined for that path.")

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(opts *options) error {
	var instances []*vssetup.Instance
	if opts.path != "" {
		if instance, err := vssetup.InstanceForPath(opts.path); err != nil {
			return err
		} else if instance == nil {
			return nil
		} else {
			instances = []*vssetup.Instance{instance}
		}
	} else {
		var err error
		instances, err = vssetup.Instances(opts.all)
		if err != nil {
			return err
		}
	}

	for i, instance := range instances {
		if i > 0 {
			fmt.Println()
		}

		formatting.PrintInstance(os.Stdout, instance)
	}

	return nil
}
