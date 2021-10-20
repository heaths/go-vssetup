//go:build windows
// +build windows

package interop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func (v *ISetupPropertyStore) GetNames() (*ole.SafeArray, error) {
	var names *ole.SafeArray
	hr, _, _ := syscall.Syscall(
		v.VTable().GetNames,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&names)),
		0,
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return names, nil
}

func (v *ISetupPropertyStore) GetValue(name string) (*ole.VARIANT, error) {
	wcs, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return nil, err
	}

	value := ole.NewVariant(ole.VT_EMPTY, 0)
	hr, _, _ := syscall.Syscall(
		v.VTable().GetValue,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(wcs)),
		uintptr(unsafe.Pointer(&value)),
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return &value, nil
}
