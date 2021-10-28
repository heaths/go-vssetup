//go:build windows
// +build windows

package interop

import (
	"syscall"

	"github.com/go-ole/go-ole"
	"github.com/heaths/go-vssetup/internal/types"
)

func bstrFunc(p uintptr, fn uintptr) (*types.Bstr, error) {
	var bstr types.Bstr
	if hr, _, _ := syscall.Syscall(fn, 2, p, uintptr(bstr.Addr()), 0); hr != 0 {
		return nil, ole.NewError(hr)
	}

	return &bstr, nil
}
