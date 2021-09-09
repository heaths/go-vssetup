//go:build windows
// +build windows

package interop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"golang.org/x/sys/windows"
)

func (v *ISetupInstance) GetInstanceId() (*uint16, error) {
	var bstrInstanceId *uint16
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstanceId,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstrInstanceId)),
		0,
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return bstrInstanceId, nil
}

func (v *ISetupInstance) GetInstallDate() (*filetime, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) GetInstallationName() (*uint16, error) {
	var bstrInstallationName *uint16
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstallationName,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstrInstallationName)),
		0,
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return bstrInstallationName, nil
}

func (v *ISetupInstance) GetInstallationPath() (*uint16, error) {
	var bstrInstallationPath *uint16
	hr, _, _ := syscall.Syscall(
		v.VTable().GetInstallationPath,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstrInstallationPath)),
		0,
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return bstrInstallationPath, nil
}

func (v *ISetupInstance) GetInstallationVersion() (*uint16, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) GetDisplayName(lcid uint32) (*uint16, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) GetDescription(lcid uint32) (*uint16, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) ResolvePath(pwszRelativePath *uint16) (*uint16, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}
