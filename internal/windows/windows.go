//go:build windows
// +build windows

package windows

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	modoleaut32 = syscall.NewLazyDLL("oleaut32.dll")

	procGetUserDefaultLCID = modkernel32.NewProc("GetUserDefaultLCID")
	procSysAllocString     = modoleaut32.NewProc("SysAllocString")
	procSysFreeString      = modoleaut32.NewProc("SysFreeString")
)

func GetUserDefaultLCID() uint32 {
	lcid, _, _ := syscall.Syscall(
		procGetUserDefaultLCID.Addr(),
		0,
		0,
		0,
		0,
	)

	return uint32(lcid)
}

// Work around https://github.com/go-ole/go-ole/issues/221
func SysAllocString(s string) *uint16 {
	if wcs, err := syscall.UTF16PtrFromString(s); err == nil {
		bstr, _, _ := procSysAllocString.Call(uintptr(unsafe.Pointer(wcs)))
		return (*uint16)(unsafe.Pointer(bstr))
	}

	return nil
}

// Work around https://github.com/go-ole/go-ole/issues/221
func SysFreeString(v *uint16) error {
	hr, _, _ := procSysFreeString.Call(uintptr(unsafe.Pointer(v)))
	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
