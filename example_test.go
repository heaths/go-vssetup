//go:build live
// +build live

package vssetup_test

import (
	"fmt"

	"github.com/heaths/go-vssetup"
	"golang.org/x/text/language"
)

func Example() {
	instances, _ := vssetup.Instances(false)
	for _, instance := range instances {
		if s, err := instance.DisplayName(language.AmericanEnglish); err == nil {
			fmt.Println(s)
		}
	}

	// Output: Visual Studio Enterprise 2019
}
