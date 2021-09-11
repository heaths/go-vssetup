//go:build !windows
// +build !windows

package interop

import "github.com/heaths/go-vssetup/internal/errors"

func (v *ISetupConfiguration) EnumInstances() (*IEnumSetupInstances, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupConfiguration) GetInstanceForCurrentProcess() (*ISetupInstance, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupConfiguration) GetInstanceForPath(path string) (*ISetupInstance, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupConfiguration2) EnumAllInstances() (*IEnumSetupInstances, error) {
	return nil, errors.NotImplemented(nil)
}
