//go:build !windows
// +build !windows

package interop

import (
	"github.com/go-ole/go-ole"
)

//nolint:stylecheck
func (v *ISetupInstance) GetInstanceId() (*uint16, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) GetInstallDate() (*filetime, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) GetInstallationName() (*uint16, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) GetInstallationPath() (*uint16, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
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
