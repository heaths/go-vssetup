//go:build windows && live
// +build windows,live

package vssetup_test

import (
	"fmt"

	"github.com/heaths/go-vssetup"
)

func Example() {
	instances, _ := vssetup.Instances(false)
	for _, instance := range instances {
		// Get the display name for en-us (LCID: 1033).
		if s, err := instance.DisplayName(1033); err == nil {
			fmt.Println(s)
		}
	}

	// Output: Visual Studio Enterprise 2019
}
