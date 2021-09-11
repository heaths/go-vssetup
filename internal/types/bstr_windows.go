//go:build windows
// +build windows

package types

import (
	"runtime"
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
	bstr := &Bstr{
		val: sysAllocString(s),
	}

	runtime.SetFinalizer(bstr, (*Bstr).Close)
	return bstr
}

func (b *Bstr) Close() error {
	if b.val != nil {
		if err := sysFreeString(b.val); err != nil {
			return err
		}

		b.val = nil
		runtime.SetFinalizer(b, nil)
	}

	return nil
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
