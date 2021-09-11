//go:build windows
// +build windows

package interop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/types"
)

func (v *ISetupInstance) GetInstanceId() (*uint16, error) { //nolint:stylecheck
	var bstr *uint16
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstanceId,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0,
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return bstr, nil
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

func (v *ISetupInstance) GetInstallationName() (*uint16, error) {
	var bstr *uint16
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstallationName,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0,
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return bstr, nil
}

func (v *ISetupInstance) GetInstallationPath() (*uint16, error) {
	var bstr *uint16
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstallationPath,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0,
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return bstr, nil
}

func (v *ISetupInstance) GetInstallationVersion() (*uint16, error) {
	var bstr *uint16
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstallationVersion,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0,
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return bstr, nil
}

func (v *ISetupInstance) GetDisplayName(lcid uint32) (*uint16, error) {
	var bstr *uint16
	hr, _, _ := syscall.Syscall(
		v.VTable().GetDisplayName,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(&bstr)),
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return bstr, nil
}

func (v *ISetupInstance) GetDescription(lcid uint32) (*uint16, error) {
	var bstr *uint16
	hr, _, _ := syscall.Syscall(
		v.VTable().GetDescription,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(lcid),
		uintptr(unsafe.Pointer(&bstr)),
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return bstr, nil
}

func (v *ISetupInstance) ResolvePath(pwszRelativePath *uint16) (*uint16, error) {
	var bstr *uint16
	hr, _, _ := syscall.Syscall(
		v.VTable().ResolvePath,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pwszRelativePath)),
		uintptr(unsafe.Pointer(&bstr)),
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return bstr, nil
}
