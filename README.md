# Visual Studio Setup Configuration

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
    instanceId, _ := instance.InstanceId()
    fmt.Println("InstanceId = ", instanceId)
}
```

## FAQ

* **On what platforms does this work?**

  This crate will only compile and work on Windows.

* **Is this project supported by Microsoft?**

  Though I am the developer who wrote the Setup Configuration API while working for Microsoft, this crate is unsupported by Microsoft.
