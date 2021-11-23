package vssetup

import (
	"github.com/heaths/go-vssetup/internal/interop"
)

type PropertyStore interface {
	queryISetupPropertyStore() (*interop.ISetupPropertyStore, error)
}

// GetProperties gets all intrinsic properties for an Instance or ErrorState.
func GetProperties(store PropertyStore) (map[string]interface{}, error) {
	if props, err := store.queryISetupPropertyStore(); err != nil {
		return nil, err
	} else {
		defer props.Release()

		var names []string
		if names, err = props.GetNames(); err != nil {
			return nil, err
		}

		properties := make(map[string]interface{}, len(names))
		for _, name := range names {
			if vt, err := props.GetValue(name); err != nil {
				return nil, err
			} else {
				properties[name] = vt.Value()
			}
		}

		return properties, nil
	}
}
