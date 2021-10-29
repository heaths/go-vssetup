package interop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func (v *ISetupPropertyStore) GetNames() ([]string, error) {
	return stringArrayFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetNames)
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

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return &value, nil
}
