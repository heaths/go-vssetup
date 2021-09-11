//go:build windows
// +build windows

package interop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/types"
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
	var i *ISetupInstance
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstanceForCurrentProcess,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&i)),
		0,
	)

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return i, nil
}

func (v *ISetupConfiguration) GetInstanceForPath(path *types.Bstr) (*ISetupInstance, error) {
	var i *ISetupInstance
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstanceForPath,
		3,
		uintptr(unsafe.Pointer(v)),
		path.Pointer(),
		uintptr(unsafe.Pointer(&i)),
	)

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return i, nil
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
