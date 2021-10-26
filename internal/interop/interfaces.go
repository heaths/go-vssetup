package interop

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

//nolint:stylecheck
var (
	CLSID_SetupConfiguration = ole.NewGUID("177F0C4A-1CD3-4DE7-A32C-71DBBB9FA36D")

	IID_ISetupConfiguration          = ole.NewGUID("42843719-DB4C-46C2-8E7C-64F1816EFD5B")
	IID_ISetupConfiguration2         = ole.NewGUID("26AAB78C-4A60-49D6-AF3B-3C35BC93365D")
	IID_IEnumSetupConfiguration      = ole.NewGUID("6380BCFF-41D3-4B2E-8B2E-BF8A6810C848")
	IID_ISetupInstance               = ole.NewGUID("B41463C3-8866-43B5-BC33-2B0676F7F42E")
	IID_ISetupInstance2              = ole.NewGUID("89143C9A-05AF-49B0-B717-72E218A2185C")
	IID_ISetupPropertyStore          = ole.NewGUID("C601C175-A3BE-44BC-91F6-4568D230FC83")
	IID_ISetupPackageReference       = ole.NewGUID("da8d8a16-b2b6-4487-a2f1-594ccccd6bf5")
	IID_ISetupErrorState             = ole.NewGUID("46DCCD94-A287-476A-851E-DFBC2FFDBC20")
	IID_ISetupErrorState2            = ole.NewGUID("9871385B-CA69-48F2-BC1F-7A37CBF0B1EF")
	IID_ISetupFailedPackageReference = ole.NewGUID("E73559CD-7003-4022-B134-27DC650B280F")
)

type ISetupConfiguration struct {
	ole.IUnknown
}

type ISetupConfigurationVtbl struct {
	ole.IUnknownVtbl
	EnumInstances                uintptr
	GetInstanceForCurrentProcess uintptr
	GetInstanceForPath           uintptr
}

func (v *ISetupConfiguration) VTable() *ISetupConfigurationVtbl {
	return (*ISetupConfigurationVtbl)(unsafe.Pointer(v.RawVTable))
}

type ISetupConfiguration2 struct {
	ISetupConfiguration
}

type ISetupConfiguration2Vtbl struct {
	ISetupConfigurationVtbl
	EnumAllInstances uintptr
}

func (v *ISetupConfiguration2) VTable() *ISetupConfiguration2Vtbl {
	return (*ISetupConfiguration2Vtbl)(unsafe.Pointer(v.RawVTable))
}

type IEnumSetupInstances struct {
	ole.IUnknown
}

type IEnumSetupInstancesVtbl struct {
	ole.IUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

func (v *IEnumSetupInstances) VTable() *IEnumSetupInstancesVtbl {
	return (*IEnumSetupInstancesVtbl)(unsafe.Pointer(v.RawVTable))
}

type ISetupInstance struct {
	ole.IUnknown
}

type ISetupInstanceVtbl struct {
	ole.IUnknownVtbl
	GetInstanceId          uintptr //nolint:stylecheck
	GetInstallDate         uintptr
	GetInstallationName    uintptr
	GetInstallationPath    uintptr
	GetInstallationVersion uintptr
	GetDisplayName         uintptr
	GetDescription         uintptr
	ResolvePath            uintptr
}

func (v *ISetupInstance) VTable() *ISetupInstanceVtbl {
	return (*ISetupInstanceVtbl)(unsafe.Pointer(v.RawVTable))
}

type ISetupInstance2 struct {
	ISetupInstance
}

type ISetupInstance2Vtbl struct {
	ISetupInstanceVtbl
	GetState       uintptr
	GetPackages    uintptr
	GetProduct     uintptr
	GetProductPath uintptr
	GetErrors      uintptr
	IsLaunchable   uintptr
	IsComplete     uintptr
	GetProperties  uintptr
	GetEnginePath  uintptr
}

func (v *ISetupInstance2) VTable() *ISetupInstance2Vtbl {
	return (*ISetupInstance2Vtbl)(unsafe.Pointer(v.RawVTable))
}

type ISetupPropertyStore struct {
	ole.IUnknown
}

type ISetupPropertyStoreVtbl struct {
	ole.IUnknownVtbl
	GetNames uintptr
	GetValue uintptr
}

func (v *ISetupPropertyStore) VTable() *ISetupPropertyStoreVtbl {
	return (*ISetupPropertyStoreVtbl)(unsafe.Pointer(v.RawVTable))
}

type ISetupPackageReference struct {
	ole.IUnknown
}

type ISetupPackageReferenceVtbl struct {
	ole.IUnknownVtbl
	GetId          uintptr //nolint:stylecheck
	GetVersion     uintptr
	GetChip        uintptr
	GetLanguage    uintptr
	GetBranch      uintptr
	GetType        uintptr
	GetUniqueId    uintptr //nolint:stylecheck
	GetIsExtension uintptr
}

func (v *ISetupPackageReference) VTable() *ISetupPackageReferenceVtbl {
	return (*ISetupPackageReferenceVtbl)(unsafe.Pointer(v.RawVTable))
}

type ISetupErrorState struct {
	ole.IUnknown
}

type ISetupErrorStateVtbl struct {
	ole.IUnknownVtbl
	GetFailedPackages  uintptr
	GetSkippedPackages uintptr
}

func (v *ISetupErrorState) VTable() *ISetupErrorStateVtbl {
	return (*ISetupErrorStateVtbl)(unsafe.Pointer(v.RawVTable))
}

type ISetupErrorState2 struct {
	ISetupErrorState
}

type ISetupErrorState2Vtbl struct {
	ISetupErrorStateVtbl
	GetErrorLogFilePath uintptr
	GetLogFilePath      uintptr
}

func (v *ISetupErrorState2) VTable() *ISetupErrorState2Vtbl {
	return (*ISetupErrorState2Vtbl)(unsafe.Pointer(v.RawVTable))
}

type ISetupFailedPackageReference struct {
	ISetupPackageReference
}

type ISetupFailedPackageReferenceVtbl struct {
	ISetupPackageReferenceVtbl
}

func (v *ISetupFailedPackageReference) VTable() *ISetupFailedPackageReferenceVtbl {
	return (*ISetupFailedPackageReferenceVtbl)(unsafe.Pointer(v.RawVTable))
}
