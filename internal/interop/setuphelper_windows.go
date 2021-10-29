package interop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func (v *ISetupHelper) ParseVersion(s string) (version uint64, err error) {
	wcs, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		return
	}

	hr, _, _ := syscall.Syscall(
		v.VTable().ParseVersion,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(wcs)),
		uintptr(unsafe.Pointer(&version)),
	)

	if hr != ole.S_OK {
		err = ole.NewError(hr)
	}

	return
}

func (v *ISetupHelper) ParseVersionRange(s string) (min, max uint64, err error) {
	wcs, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		return
	}

	hr, _, _ := syscall.Syscall6(
		v.VTable().ParseVersionRange,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(wcs)),
		uintptr(unsafe.Pointer(&min)),
		uintptr(unsafe.Pointer(&max)),
		0,
		0,
	)

	if hr != ole.S_OK {
		err = ole.NewError(hr)
	}

	return
}
