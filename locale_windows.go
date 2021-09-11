//go:build windows
// +build windows

package vssetup

import "syscall"

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetUserDefaultLCID = modkernel32.NewProc("GetUserDefaultLCID")
)

// Locale on Windows returns the current user's locale e.g., 1033 for en-us.
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
