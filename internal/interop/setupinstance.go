//go:build !windows
// +build !windows

package interop

import (
	"github.com/go-ole/go-ole"
	"golang.org/x/sys/windows"
)

func (v *ISetupInstance) GetInstanceId() (*uint16, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func (v *ISetupInstance) GetInstallDate() (*windows.Filetime, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func (v *ISetupInstance) GetInstallationName() (*uint16, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func (v *ISetupInstance) GetInstallationPath() (*uint16, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func (v *ISetupInstance) GetInstallationVersion() (*uint16, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func (v *ISetupInstance) GetDisplayName(lcid uint32) (*uint16, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func (v *ISetupInstance) GetDescription(lcid uint32) (*uint16, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}

func (v *ISetupInstance) ResolvePath(pwszRelativePath *uint16) (*uint16, error) {
	return nil, ole.NewError(ole.E_NOTIMPL)
}
