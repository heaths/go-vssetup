//go:build !windows
// +build !windows

package vssetup

// Locale on non-Windows platforms always returns 0.
func Locale() uint32 {
	return 0
}
