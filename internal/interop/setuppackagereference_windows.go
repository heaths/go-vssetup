//go:build windows
// +build windows

package interop

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/types"
)

func (v *ISetupPackageReference) GetId() (*types.Bstr, error) { //nolint:stylecheck
	return v.bstrFunc(v.VTable().GetId)
}

func (v *ISetupPackageReference) GetVersion() (*types.Bstr, error) {
	return v.bstrFunc(v.VTable().GetVersion)
}

func (v *ISetupPackageReference) GetChip() (*types.Bstr, error) {
	return v.bstrFunc(v.VTable().GetChip)
}

func (v *ISetupPackageReference) GetLanguage() (*types.Bstr, error) {
	return v.bstrFunc(v.VTable().GetLanguage)
}

func (v *ISetupPackageReference) GetBranch() (*types.Bstr, error) {
	return v.bstrFunc(v.VTable().GetBranch)
}

func (v *ISetupPackageReference) GetType() (*types.Bstr, error) {
	return v.bstrFunc(v.VTable().GetType)
}

func (v *ISetupPackageReference) GetUniqueId() (*types.Bstr, error) { //nolint:stylecheck
	return v.bstrFunc(v.VTable().GetUniqueId)
}

func (v *ISetupPackageReference) GetIsExtension() (bool, error) {
	var b uint32
	hr, _, _ := syscall.Syscall(
		v.VTable().GetIsExtension,
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&b)),
		0,
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return b != 0, nil
}

func (v *ISetupPackageReference) bstrFunc(fn uintptr) (*types.Bstr, error) {
	var bstr types.Bstr
	hr, _, _ := syscall.Syscall(
		fn,
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
