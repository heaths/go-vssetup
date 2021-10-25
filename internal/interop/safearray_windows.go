//go:build windows
// +build windows

package interop

import (
	"unsafe"

	"github.com/go-ole/go-ole"
	"golang.org/x/sys/windows"
)

var (
	modoleaut32 = windows.NewLazySystemDLL("oleaut32.dll")

	procSafeArrayGetElement = modoleaut32.NewProc("SafeArrayGetElement")
)

func safeArrayGetElement(sa *ole.SafeArray, index int32, pv unsafe.Pointer) (err error) {
	hr, _, _ := procSafeArrayGetElement.Call(
		uintptr(unsafe.Pointer(sa)),
		uintptr(unsafe.Pointer(&index)),
		uintptr(pv),
	)

	if hr != 0 {
		err = ole.NewError(hr)
	}

	return
}
