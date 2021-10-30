package vssetup

import (
	"fmt"

	"golang.org/x/text/language"
)

func ExampleInstances() {
	instances, _ := Instances(false)
	for _, instance := range instances {
		if s, err := instance.DisplayName(language.AmericanEnglish); err == nil {
			fmt.Println(s)
		}
	}

	// Output: Visual Studio Enterprise 2019
}

func ExampleParseVersion() {
	if version, err := ParseVersion("1.2.3.4"); err == nil {
		fmt.Println(version)
	}

	// Output: 281483566841860
}

func ExampleParseVersionRange() {
	if min, max, err := ParseVersionRange("(1.0,2.0]"); err == nil {
		fmt.Println(min, max)
	}

	// Output: 281474976710657 562949953421312
}
