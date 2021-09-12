# Changelog

All notable changes to this project will be documented here.
See [releases](https://github.com/heaths/go-vssetup/releases) for artifacts.

## v0.2.0

* Added `UserPreferredLanguage()` to instead return a `language.Tag`.
* Enabled color output if supported by the terminal.
* Removed `Locale()` that returned an LCID.

## v0.1.1 (2021-09-11)

* Fixed error handling so nil array returned if not supported or no instances found.

## v0.1.0 (2021-09-11)

* Initial release
