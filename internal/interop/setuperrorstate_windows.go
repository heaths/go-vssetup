//go:build windows
// +build windows

package interop

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/types"
)

func (v *ISetupErrorState) GetFailedPackages() ([]*ISetupFailedPackageReference, error) {
	sa, err := v.safeArrayFunc(v.VTable().GetFailedPackages)
	if err != nil {
		return nil, err
	}

	array := ole.SafeArrayConversion{
		Array: sa,
	}
	defer array.Release()

	if vt, err := array.GetType(); err != nil {
		return nil, err
	} else if vt != uint16(ole.VT_UNKNOWN) {
		return nil, fmt.Errorf("unknown packages array type: %d", vt)
	}

	count, err := array.TotalElements(0)
	if err != nil {
		return nil, err
	}

	packages := make([]*ISetupFailedPackageReference, count)
	for i := int32(0); i < count; i++ {
		var ref *ISetupFailedPackageReference
		if err := safeArrayGetElement(sa, i, unsafe.Pointer(&v)); err != nil {
			return nil, err
		}

		packages[i] = ref
	}

	return packages, nil
}

func (v *ISetupErrorState) GetSkippedPackages() ([]*ISetupPackageReference, error) {
	sa, err := v.safeArrayFunc(v.VTable().GetSkippedPackages)
	if err != nil {
		return nil, err
	}

	array := ole.SafeArrayConversion{
		Array: sa,
	}
	defer array.Release()

	if vt, err := array.GetType(); err != nil {
		return nil, err
	} else if vt != uint16(ole.VT_UNKNOWN) {
		return nil, fmt.Errorf("unknown packages array type: %d", vt)
	}

	count, err := array.TotalElements(0)
	if err != nil {
		return nil, err
	}

	packages := make([]*ISetupPackageReference, count)
	for i := int32(0); i < count; i++ {
		var ref *ISetupPackageReference
		if err := safeArrayGetElement(sa, i, unsafe.Pointer(&v)); err != nil {
			return nil, err
		}

		packages[i] = ref
	}

	return packages, nil
}

func (v *ISetupErrorState) ISetupErrorState2(v2 **ISetupErrorState2) error {
	if *v2 != nil {
		return nil
	}

	hr, _, _ := syscall.Syscall(
		v.IUnknown.VTable().QueryInterface,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(IID_ISetupErrorState2)),
		uintptr(unsafe.Pointer(v2)),
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *ISetupErrorState2) GetErrorLogFilePath() (*types.Bstr, error) {
	return v.bstrFunc(v.VTable().GetErrorLogFilePath)
}

func (v *ISetupErrorState2) GetLogFilePath() (*types.Bstr, error) {
	return v.bstrFunc(v.VTable().GetLogFilePath)
}

func (v *ISetupErrorState) safeArrayFunc(fn uintptr) (array *ole.SafeArray, err error) {
	hr, _, _ := syscall.Syscall(
		fn,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&array)),
		0,
	)

	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}

func (v *ISetupErrorState2) bstrFunc(fn uintptr) (*types.Bstr, error) {
	var bstr types.Bstr
	hr, _, _ := syscall.Syscall(
		fn,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(bstr.Addr()),
		0,
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return &bstr, nil
}
