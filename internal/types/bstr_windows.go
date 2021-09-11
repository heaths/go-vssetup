//go:build windows
// +build windows

package types

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

var (
	modoleaut32 = syscall.NewLazyDLL("oleaut32.dll")

	procSysAllocString = modoleaut32.NewProc("SysAllocString")
	procSysFreeString  = modoleaut32.NewProc("SysFreeString")
)

func NewBstr(s string) *Bstr {
	return &Bstr{
		val: sysAllocString(s),
	}
}

func (b *Bstr) Close() {
	if err := sysFreeString(b.val); err != nil {
		panic(err)
	}
}

// Work around https://github.com/go-ole/go-ole/issues/221
func sysAllocString(s string) *uint16 {
	if wcs, err := syscall.UTF16PtrFromString(s); err == nil {
		bstr, _, _ := procSysAllocString.Call(uintptr(unsafe.Pointer(wcs)))
		return (*uint16)(unsafe.Pointer(bstr))
	}

	return nil
}

// Work around https://github.com/go-ole/go-ole/issues/221
func sysFreeString(v *uint16) error {
	hr, _, _ := procSysFreeString.Call(uintptr(unsafe.Pointer(v)))
	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
