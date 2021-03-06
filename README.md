# Visual Studio Setup Configuration

[![releases](https://img.shields.io/github/v/release/heaths/go-vssetup.svg?logo=github)](https://github.com/heaths/go-vssetup/releases/latest)
[![reference](https://pkg.go.dev/badge/github.com/heaths/go-vssetup.svg)](https://pkg.go.dev/github.com/heaths/go-vssetup?GOOS=windows)
[![ci](https://github.com/heaths/go-vssetup/actions/workflows/ci.yml/badge.svg)](https://github.com/heaths/go-vssetup/actions/workflows/ci.yml)

[Microsoft Visual Studio](https://visualstudio.microsoft.com) 2017 and newer uses a new setup engine that allows multiple instances to be installed quickly and in different configuration. To enumerate these instances and find one that fulfills your requirements, the [Setup Configuration API](https://devblogs.microsoft.com/setup/documentation-available-for-the-setup-configuration-api) provides a set of interface. This crate provides a safe and idiomatic wrapper for [Go](https://golang.org).

## Example

First install the heaths/go-vssetup module:

```bash
go get github.com/heaths/go-vssetup
```

You can then use the `vssetup` module to enumerate instances:

```go
// Enumerate launchable instances.
instances, _ := vssetup.Instances(false)
for _, instance := range instances {
    if installationPath, err := instance.InstallationPath(); err == nil {
        fmt.Println("InstallationPath =", installationPath)
    }
}
```

### vswhere

While this project is not meant to replace the official [vswhere] that ships with Visual Studio and is available in some package managers, nor is it meant to be fully compatible with all official vswhere options, this project does have a `vswhere` command that allows you to quickly see this module in action:

```bash
go run ./cmd/vswhere
```

You can also download a binary from [releases](https://github.com/heaths/go-vssetup/releases). It is not call-compatible with the official [vswhere].

## FAQ

* **On what platforms does this work?**

  This module should compile on any platform but will only return available instances on Windows if Visual Studio 2017 or newer is installed.

* **Is this project supported by Microsoft?**

  Though I am the developer who wrote the Setup Configuration API while working for Microsoft, this module is unsupported by Microsoft.

[vswhere]: https://github.com/microsoft/vswhere
