//go:build windows
// +build windows

package interop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func (v *IEnumSetupInstances) Next(celt uint32) ([]*ISetupInstance, error) {
	rgelt := make([]*ISetupInstance, celt)
	var celtFetched uint32
	hr, _, _ := syscall.Syscall6(
		v.VTable().Next,
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(celt),
		uintptr(unsafe.Pointer(&rgelt[0])),
		uintptr(unsafe.Pointer(&celtFetched)),
		0,
		0,
	)

	if hr != ole.S_OK && hr != S_FALSE {
		return nil, ole.NewError(hr)
	}

	if celtFetched == 0 {
		return []*ISetupInstance{}, nil
	}

	instances := unsafe.Slice(&rgelt[0], celtFetched)
	return instances, nil
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
