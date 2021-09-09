//go:build !windows
// +build !windows

package interop

import "github.com/go-ole/go-ole"

func (v *ISetupConfiguration) EnumInstances() (*IEnumSetupInstances, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func (v *ISetupConfiguration) GetInstanceForCurrentProcess() (*ISetupInstance, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func (v *ISetupConfiguration) GetInstanceForPath() (*ISetupInstance, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func (v *ISetupConfiguration2) EnumAllInstances() (*IEnumSetupInstances, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}
