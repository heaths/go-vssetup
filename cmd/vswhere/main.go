package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/heaths/go-vssetup"
	"github.com/heaths/go-vssetup/internal/formatting"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

type options struct {
	all     bool
	include formatting.Includes
	locale  *language.Tag
	path    string
	raw     bool
}

func main() {
	opts := options{}
	var include []string
	var locale string

	const validIncludes = "errors, packages"

	root := cobra.Command{
		Short: "Locates instances of Visual Studio",
		Use:   filepath.Base(os.Args[0]),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			for _, arg := range include {
				switch arg {
				case "errors":
					opts.include |= formatting.Errors
				case "packages":
					opts.include |= formatting.Packages
				default:
					return fmt.Errorf("invalid argument: %s, accepts: %s", arg, validIncludes)
				}
			}
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
	root.Flags().StringArrayVar(&include, "include", nil, fmt.Sprintf("Other information to include: %s", validIncludes))
	root.Flags().StringVar(&locale, "locale", "en", "The locale to use for localized values. The default is your preferred system locale.")
	root.Flags().StringVar(&opts.path, "path", "", "Gets an instance for the given path, if any defined for that path.")
	root.Flags().BoolVar(&opts.raw, "raw", false, "Format all intrinsic properties as defined.")

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

	options := formatting.Options{
		Include: opts.include,
		Locale:  *opts.locale,
		Raw:     opts.raw,
	}

	for i, instance := range instances {
		if i > 0 {
			fmt.Println()
		}
		formatting.PrintInstance(os.Stdout, instance, options)
	}

	return nil
}
