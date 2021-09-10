//go:build !windows
// +build !windows

package vssetup

func Locale() uint32 {
	return 0
}
