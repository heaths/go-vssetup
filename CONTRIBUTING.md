# Contributing guid

Contributions are welcome! The goal is to provide bindings for the [Visual Studio Setup Configuration API][vssetup] and make it easy for developers on Windows to find and use Visual Studio.

## Getting started

Because the [Visual Studio Setup Configuration API][vssetup] only runs on Windows, building and running tests on Windows is highly encouraged. All the non-Windows functions should return `E_NOTIMPL` (0x80004001), which is conveniently available from the `internal/errors` package.

### File naming convention

To support generating examples for `godoc`, any Windows-only file names should have the suffix `_windows`, while any Windows-only tests should have the suffix `_windows_test`. No `//go:build windows` comment is necessary.

Conversely, non-Windows files should have the suffix `_other` but must have a comment at the very top before the `package` declaration, with a blank line preceding the `package` declaration:

```go
//go:build !windows
// +build !windows

```

## Testing

To run tests, use `go test`:

```bash
go test ./...
go test -cover ./...
```

On Windows, this will also exercise examples that may fail depending on what version(s) of Visual Studio  you have installed. You can either ignore these failures, or temporarily change the `// Output:` comment to match what you have installed. Do not commit or attempt to merge these changes, however. The expected output should match [what version(s) of Visual Studio is installed][hosted-agents] on `windows-latest` agents for GitHub Actions.

### vswhere

For integration testing on Windows, a `cmd/vswhere` command is included. This is not meant as a replacement for the [official vswhere][vswhere], but supports similar switches to test different features of these Go bindings. `cmd/vswhere`, thus, is only meant for development purposes with this repository and should not be used in place of the official `vswhere`.

[hosted-agents]: https://docs.github.com/actions/using-github-hosted-runners/about-github-hosted-runners
[vssetup]: https://devblogs.microsoft.com/setup/documentation-available-for-the-setup-configuration-api
[vswhere]: https://github.com/microsoft/vswhere
