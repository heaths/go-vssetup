//go:build windows
// +build windows

package interop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func (v *ISetupConfiguration) EnumInstances() (*IEnumSetupInstances, error) {
	var e *IEnumSetupInstances
	hr, _, _ := syscall.Syscall(
		v.VTable().EnumInstances,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&e)),
		0,
	)

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return e, nil
}

func (v *ISetupConfiguration) GetInstanceForCurrentProcess() (*ISetupInstance, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupConfiguration) GetInstanceForPath() (*ISetupInstance, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupConfiguration2) EnumAllInstances() (*IEnumSetupInstances, error) {
	var e *IEnumSetupInstances
	hr, _, _ := syscall.Syscall(
		v.VTable().EnumAllInstances,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&e)),
		0,
	)

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return e, nil
}
