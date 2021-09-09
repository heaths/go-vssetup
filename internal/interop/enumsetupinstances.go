//go:build !windows
// +build !windows

package interop

import "github.com/go-ole/go-ole"

func (v *IEnumSetupInstances) Next(celt uint32) ([]*ISetupInstance, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *IEnumSetupInstances) Skip(celt uint32) error {
	return ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *IEnumSetupInstances) Reset() error {
	return ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *IEnumSetupInstances) Clone() (*IEnumSetupInstances, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}
