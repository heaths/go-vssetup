//go:build !windows
// +build !windows

package interop

import "github.com/go-ole/go-ole"

func (v *IEnumSetupInstances) Next(celt uint32) ([]*ISetupInstance, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func (v *IEnumSetupInstances) Skip(celt uint32) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IEnumSetupInstances) Reset() error {
	return ole.NewError(ole.E_NOTIMPL)
}

func (v *IEnumSetupInstances) Clone() (*IEnumSetupInstances, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}
