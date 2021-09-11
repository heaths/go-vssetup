/*
The vssetup package locates installations of Microsoft Visual Studio 2017 and
newer, as well as related products installed by the Visual Studio Setup engine
on Microsoft Windows. You should be able to compile this package on other
platforms, but using any APIs will return a "Not implemented" error.

You can enumerate launchable instances using vssetup.Instances(bool):

	instances, _ := vssetup.Instances(false)
	for _, instance := range instances {
		if installationPath, err := instance.InstallationPath(); err == nil {
			fmt.Println("InstallationPath =", installationPath)
		}
	}

This API wraps the Setup Configuration API, which you can read more about at
https://devblogs.microsoft.com/setup/documentation-available-for-the-setup-configuration-api.

These bindings for Go are *unofficial*, despite these APIs being written and
maintained by the architect and one of the original developers on the Visual
Studio Setup engine and Configuration APIs.
*/
package vssetup
