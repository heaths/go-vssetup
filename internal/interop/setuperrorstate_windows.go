package interop

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/types"
	"github.com/heaths/go-vssetup/internal/windows"
)

func (v *ISetupErrorState) GetFailedPackages() ([]*ISetupFailedPackageReference, error) {
	sa, err := safeArrayFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetFailedPackages)
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
		if err := windows.SafeArrayGetElement(sa, i, unsafe.Pointer(&ref)); err != nil {
			return nil, err
		}

		packages[i] = ref
	}

	return packages, nil
}

func (v *ISetupErrorState) GetSkippedPackages() ([]*ISetupPackageReference, error) {
	sa, err := safeArrayFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetSkippedPackages)
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
		if err := windows.SafeArrayGetElement(sa, i, unsafe.Pointer(&ref)); err != nil {
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

	if hr != ole.S_OK {
		return ole.NewError(hr)
	}

	return nil
}

func (v *ISetupErrorState2) GetErrorLogFilePath() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetErrorLogFilePath)
}

func (v *ISetupErrorState2) GetLogFilePath() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetLogFilePath)
}

func (v *ISetupErrorState) ISetupErrorState3(v3 **ISetupErrorState3) error {
	if *v3 != nil {
		return nil
	}

	hr, _, _ := syscall.Syscall(
		v.IUnknown.VTable().QueryInterface,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(IID_ISetupErrorState3)),
		uintptr(unsafe.Pointer(v3)),
	)

	if hr != ole.S_OK {
		return ole.NewError(hr)
	}

	return nil
}

func (v *ISetupErrorState3) GetRuntimeError() (*ISetupErrorInfo, error) {
	var info *ISetupErrorInfo
	hr, _, _ := syscall.Syscall(
		v.VTable().GetRuntimeError,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&info)),
		0,
	)

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return info, nil
}

func (v *ISetupErrorInfo) GetErrorHResult() (uint32, error) {
	return uint32Func(uintptr(unsafe.Pointer(v)), v.VTable().GetErrorHResult)
}

func (v *ISetupErrorInfo) GetErrorClassName() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetErrorClassName)
}

func (v *ISetupErrorInfo) GetErrorMessage() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetErrorMessage)
}
