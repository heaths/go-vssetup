//go:build windows
// +build windows

package vssetup

import (
	"github.com/heaths/go-vssetup/internal/windows"
)

// Locale on Windows returns the current user's locale e.g., 1033 for en-us.
func Locale() uint32 {
	return windows.GetUserDefaultLCID()
}
