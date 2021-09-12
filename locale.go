//go:build !windows
// +build !windows

package vssetup

import "golang.org/x/text/language"

// UserPreferredLanguage returns the user-preferred language name.
// On non-Windows system this function always returns "en".
func UserPreferredLanguage() language.Tag {
	return language.English
}

func lcid(t language.Tag) uint32 {
	return 0
}
