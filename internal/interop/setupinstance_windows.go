//go:build windows
// +build windows

package interop

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/types"
	"github.com/heaths/go-vssetup/internal/windows"
)

func (v *ISetupInstance) GetInstanceId() (*types.Bstr, error) { //nolint:stylecheck
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetInstanceId)
}

func (v *ISetupInstance) GetInstallDate() (*types.Filetime, error) {
	var ft types.Filetime
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstallDate,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&ft)),
		0,
	)

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return &ft, nil
}

func (v *ISetupInstance) GetInstallationName() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetInstallationName)
}

func (v *ISetupInstance) GetInstallationPath() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetInstallationPath)
}

func (v *ISetupInstance) GetInstallationVersion() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetInstallationVersion)
}

func (v *ISetupInstance) GetDisplayName(lcid uint32) (*types.Bstr, error) {
	return localizedBstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetDisplayName, lcid)
}

func (v *ISetupInstance) GetDescription(lcid uint32) (*types.Bstr, error) {
	return localizedBstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetDescription, lcid)
}

func (v *ISetupInstance) ResolvePath(relativePath string) (*types.Bstr, error) {
	wcs, err := syscall.UTF16PtrFromString(relativePath)
	if err != nil {
		return nil, err
	}

	var bstr types.Bstr
	hr, _, _ := syscall.Syscall(
		v.VTable().ResolvePath,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(wcs)),
		bstr.Addr(),
	)

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return &bstr, nil
}

func (v *ISetupInstance) ISetupInstance2(v2 **ISetupInstance2) error {
	if *v2 != nil {
		return nil
	}

	hr, _, _ := syscall.Syscall(
		v.IUnknown.VTable().QueryInterface,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(IID_ISetupInstance2)),
		uintptr(unsafe.Pointer(v2)),
	)

	if hr != ole.S_OK {
		return ole.NewError(hr)
	}

	return nil
}

func (v *ISetupInstance2) GetState() (uint32, error) {
	var state uint32
	hr, _, _ := syscall.Syscall(
		v.VTable().GetState,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&state)),
		0,
	)

	if hr != ole.S_OK {
		return 0, ole.NewError(hr)
	}

	return state, nil
}

func (v *ISetupInstance2) GetPackages() ([]*ISetupPackageReference, error) {
	sa, err := safeArrayFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetPackages)
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
		var v *ISetupPackageReference
		if err := windows.SafeArrayGetElement(sa, i, unsafe.Pointer(&v)); err != nil {
			return nil, err
		}

		packages[i] = v
	}

	return packages, nil
}

func (v *ISetupInstance2) GetProduct() (*ISetupPackageReference, error) {
	var product *ISetupPackageReference
	hr, _, _ := syscall.Syscall(
		v.VTable().GetProduct,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&product)),
		0,
	)

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return product, nil
}

func (v *ISetupInstance2) GetProductPath() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetProductPath)
}

func (v *ISetupInstance2) GetErrors() (*ISetupErrorState, error) {
	var errors *ISetupErrorState
	hr, _, _ := syscall.Syscall(
		v.VTable().GetErrors,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&errors)),
		0,
	)

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
	} else if errors == nil {
		return nil, nil
	}

	return errors, nil
}

func (v *ISetupInstance2) IsLaunchable() (bool, error) {
	return boolFunc(uintptr(unsafe.Pointer(v)), v.VTable().IsLaunchable)
}

func (v *ISetupInstance2) IsComplete() (bool, error) {
	return boolFunc(uintptr(unsafe.Pointer(v)), v.VTable().IsComplete)
}

func (v *ISetupInstance2) GetProperties() (*ISetupPropertyStore, error) {
	var properties *ISetupPropertyStore
	hr, _, _ := syscall.Syscall(
		v.VTable().GetProperties,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&properties)),
		0,
	)

	if hr != ole.S_OK {
		return nil, ole.NewError(hr)
	}

	return properties, nil
}

func (v *ISetupInstance2) GetEnginePath() (*types.Bstr, error) {
	return bstrFunc(uintptr(unsafe.Pointer(v)), v.VTable().GetEnginePath)
}
