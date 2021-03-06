package interop

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/types"
	"github.com/heaths/go-vssetup/internal/windows"
)

func (v *ISetupPackageReference) GetId() (*types.Bstr, error) { //nolint:stylecheck
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetId)
}

func (v *ISetupPackageReference) GetVersion() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetVersion)
}

func (v *ISetupPackageReference) GetChip() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetChip)
}

func (v *ISetupPackageReference) GetLanguage() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetLanguage)
}

func (v *ISetupPackageReference) GetBranch() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetBranch)
}

func (v *ISetupPackageReference) GetType() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetType)
}

func (v *ISetupPackageReference) GetUniqueId() (*types.Bstr, error) { //nolint:stylecheck
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetUniqueId)
}

func (v *ISetupPackageReference) GetIsExtension() (bool, error) {
	return boolFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetIsExtension)
}

func (v *ISetupFailedPackageReference) GetISetupPackageReference() (*ISetupPackageReference, error) {
	var ref *ISetupPackageReference
	hr, _, _ := syscall.Syscall(
		v.IUnknown.VTable().QueryInterface,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(IID_ISetupPackageReference)),
		uintptr(unsafe.Pointer(&ref)),
	)

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return ref, nil
}

func (v *ISetupFailedPackageReference) ISetupFailedPackageReference2(v2 **ISetupFailedPackageReference2) error {
	if *v2 != nil {
		return nil
	}

	hr, _, _ := syscall.Syscall(
		v.IUnknown.VTable().QueryInterface,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(IID_ISetupFailedPackageReference2)),
		uintptr(unsafe.Pointer(v2)),
	)

	if hr != ole.S_OK {
		return ole.NewError(hr)
	}

	return nil
}

func (v *ISetupFailedPackageReference2) GetLogFilePath() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetLogFilePath)
}

func (v *ISetupFailedPackageReference2) GetDescription() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetDescription)
}

func (v *ISetupFailedPackageReference2) GetSignature() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetSignature)
}

func (v *ISetupFailedPackageReference2) GetDetails() ([]string, error) {
	return stringArrayFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetDetails)
}

func (v *ISetupFailedPackageReference2) GetAffectedPackages() ([]*ISetupPackageReference, error) {
	var sa *ole.SafeArray
	hr, _, _ := syscall.Syscall(
		v.VTable().GetAffectedPackages,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&sa)),
		0,
	)

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
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
		var v *ISetupPackageReference
		if err := windows.SafeArrayGetElement(sa, i, unsafe.Pointer(&v)); err != nil {
			return nil, err
		}

		packages[i] = v
	}

	return packages, nil
}

func (v *ISetupFailedPackageReference) ISetupFailedPackageReference3(v3 **ISetupFailedPackageReference3) error {
	if *v3 != nil {
		return nil
	}

	hr, _, _ := syscall.Syscall(
		v.IUnknown.VTable().QueryInterface,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(IID_ISetupFailedPackageReference3)),
		uintptr(unsafe.Pointer(v3)),
	)

	if hr != ole.S_OK {
		return ole.NewError(hr)
	}

	return nil
}

func (v *ISetupFailedPackageReference3) GetAction() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetAction)
}

func (v *ISetupFailedPackageReference3) GetReturnCode() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetReturnCode)
}

func (v *ISetupPackageReference) ISetupProductReference(pr **ISetupProductReference) error {
	if *pr != nil {
		return nil
	}

	hr, _, _ := syscall.Syscall(
		v.IUnknown.VTable().QueryInterface,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(IID_ISetupProductReference)),
		uintptr(unsafe.Pointer(pr)),
	)

	if hr != ole.S_OK {
		return ole.NewError(hr)
	}

	return nil
}

func (v *ISetupProductReference) GetIsInstalled() (bool, error) {
	return boolFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetIsInstalled)
}

func (v *ISetupPackageReference) ISetupProductReference2(pr2 **ISetupProductReference2) error {
	if *pr2 != nil {
		return nil
	}

	hr, _, _ := syscall.Syscall(
		v.IUnknown.VTable().QueryInterface,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(IID_ISetupProductReference2)),
		uintptr(unsafe.Pointer(pr2)),
	)

	if hr != ole.S_OK {
		return ole.NewError(hr)
	}

	return nil
}

func (v *ISetupProductReference2) GetSupportsExtensions() (bool, error) {
	return boolFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetSupportsExtensions)
}
