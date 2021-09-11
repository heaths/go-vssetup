//go:build !windows
// +build !windows

package interop

import (
	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/types"
)

//nolint:stylecheck
func (v *ISetupInstance) GetInstanceId() (*types.Bstr, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) GetInstallDate() (*types.Filetime, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) GetInstallationName() (*types.Bstr, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) GetInstallationPath() (*types.Bstr, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) GetInstallationVersion() (*types.Bstr, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) GetDisplayName(lcid uint32) (*types.Bstr, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) GetDescription(lcid uint32) (*types.Bstr, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}

func (v *ISetupInstance) ResolvePath(relativePath string) (*types.Bstr, error) {
	return nil, ole.NewErrorWithDescription(ole.E_NOTIMPL, "not implemented")
}
