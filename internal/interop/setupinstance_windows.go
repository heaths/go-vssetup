//go:build windows
// +build windows

package interop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/types"
)

func (v *ISetupInstance) GetInstanceId() (*types.Bstr, error) { //nolint:stylecheck
	var bstr types.Bstr
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstanceId,
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

func (v *ISetupInstance) GetInstallDate() (*types.Filetime, error) {
	var ft types.Filetime
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstallDate,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&ft)),
		0,
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return &ft, nil
}

func (v *ISetupInstance) GetInstallationName() (*types.Bstr, error) {
	var bstr types.Bstr
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstallationName,
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

func (v *ISetupInstance) GetInstallationPath() (*types.Bstr, error) {
	var bstr types.Bstr
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstallationPath,
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

func (v *ISetupInstance) GetInstallationVersion() (*types.Bstr, error) {
	var bstr types.Bstr
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstallationVersion,
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

func (v *ISetupInstance) GetDisplayName(lcid uint32) (*types.Bstr, error) {
	var bstr types.Bstr
	hr, _, _ := syscall.Syscall(
		v.VTable().GetDisplayName,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(lcid),
		uintptr(bstr.Addr()),
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return &bstr, nil
}

func (v *ISetupInstance) GetDescription(lcid uint32) (*types.Bstr, error) {
	var bstr types.Bstr
	hr, _, _ := syscall.Syscall(
		v.VTable().GetDescription,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(lcid),
		uintptr(bstr.Addr()),
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return &bstr, nil
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

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return &bstr, nil
}
