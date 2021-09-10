//go:build windows
// +build windows

package vssetup

import "syscall"

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetUserDefaultLCID = modkernel32.NewProc("GetUserDefaultLCID")
)

func Locale() uint32 {
	lcid, _, _ := syscall.Syscall(
		procGetUserDefaultLCID.Addr(),
		0,
		0,
		0,
		0,
	)

	return uint32(lcid)
}
