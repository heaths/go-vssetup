package main

import (
	"fmt"
	"os"

	"github.com/heaths/go-vssetup"
	"github.com/heaths/go-vssetup/internal/formatting"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

type options struct {
	all    bool
	locale *language.Tag
	path   string
}

func main() {
	opts := options{}
	var locale string

	root := cobra.Command{
		Use: "Locates instances of Visual Studio",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if locale, err := language.Parse(locale); err != nil {
				return fmt.Errorf("invalid locale: %w", err)
			} else {
				opts.locale = &locale
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			return run(&opts)
		},
	}

	root.Flags().BoolVar(&opts.all, "all", false, "Finds all instances even if they are incomplete and may not launch.")
	root.Flags().StringVar(&locale, "locale", "en", "The locale to use for localized values. The default is your preferred system locale.")
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

	if opts.locale == nil {
		opts.locale = &language.English
	}

	for i, instance := range instances {
		if i > 0 {
			fmt.Println()
		}

		formatting.PrintInstance(os.Stdout, instance, *opts.locale)
	}

	return nil
}
