//go:build !windows
// +build !windows

package interop

import (
	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/errors"
)

func (v *ISetupPropertyStore) GetNames() (*ole.SafeArray, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupPropertyStore) GetValue(name string) (*ole.VARIANT, error) {
	return nil, errors.NotImplemented(nil)
}
