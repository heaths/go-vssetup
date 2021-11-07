//go:build !windows
// +build !windows

package interop

import (
	"github.com/heaths/go-vssetup/internal/errors"
	"github.com/heaths/go-vssetup/internal/types"
)

func (v *ISetupErrorState) GetFailedPackages() ([]*ISetupFailedPackageReference, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupErrorState) GetSkippedPackages() ([]*ISetupPackageReference, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupErrorState) ISetupErrorState2(v2 **ISetupErrorState2) error {
	return errors.NotImplemented(nil)
}

func (v *ISetupErrorState2) GetErrorLogFilePath() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupErrorState2) GetLogFilePath() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupErrorState) ISetupErrorState3(v3 **ISetupErrorState3) error {
	return errors.NotImplemented(nil)
}

func (v *ISetupErrorState3) GetRuntimeError() (*ISetupErrorInfo, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupErrorInfo) GetErrorHResult() (uint32, error) {
	return 0, errors.NotImplemented(nil)
}

func (v *ISetupErrorInfo) GetErrorClassName() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}

func (v *ISetupErrorInfo) GetErrorMessage() (*types.Bstr, error) {
	return nil, errors.NotImplemented(nil)
}
