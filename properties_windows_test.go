package vssetup

import "fmt"

func ExampleGetProperties() {
	instances, _ := Instances(false)
	for _, instance := range instances {
		if properties, err := GetProperties(instance); err == nil {
			for name, value := range properties {
				fmt.Println(name, "=", value)
			}
		}
	}
}
