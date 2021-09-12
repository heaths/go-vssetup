//go:build windows
// +build windows

package vssetup

import (
	"github.com/heaths/go-vssetup/internal/windows"
	win32 "golang.org/x/sys/windows"
	"golang.org/x/text/language"
)

// UserPreferredLanguage returns the user-preferred language name.
// The fallback if no language is preferred is "en".
func UserPreferredLanguage() language.Tag {
	if names, err := win32.GetUserPreferredUILanguages(win32.MUI_LANGUAGE_NAME); err == nil {
		for _, name := range names {
			if tag, err := language.Parse(name); err == nil {
				return tag
			}
		}
	}

	return language.English
}

func lcid(locale language.Tag) uint32 {
	lcid := windows.LocaleNameToLCID(locale.String())
	if lcid == 0 && !locale.Parent().IsRoot() {
		t := locale.Parent()
		lcid = windows.LocaleNameToLCID(t.String())
	}

	return lcid
}
