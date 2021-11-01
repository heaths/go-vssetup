package vssetup

import (
	"runtime"

	"github.com/heaths/go-vssetup/internal/interop"
)

// PackageReference describes unique attributes of a package.
type PackageReference struct {
	v *interop.ISetupPackageReference
}

func newPackageReference(v *interop.ISetupPackageReference) *PackageReference {
	p := &PackageReference{
		v: v,
	}

	runtime.SetFinalizer(p, (*PackageReference).Close)
	return p
}

// Close releases any resources used by this PackageReference immediately.
func (p *PackageReference) Close() error {
	if p.v != nil {
		// Call IUnknown.Release() but leave v assigned to avoid AV exceptions.
		p.v.Release()

		runtime.SetFinalizer(p, nil)
	}

	return nil
}

// ID gets the package reference ID.
func (p *PackageReference) ID() (string, error) {
	return getStringFunc(p.v.GetId)
}

// Version gets the package reference version.
func (p *PackageReference) Version() (string, error) {
	return getStringFunc(p.v.GetVersion)
}

// Chip gets the package reference chip.
func (p *PackageReference) Chip() (string, error) {
	return getStringFunc(p.v.GetChip)
}

// Language gets the package reference language.
func (p *PackageReference) Language() (string, error) {
	return getStringFunc(p.v.GetLanguage)
}

// Branch gets the package reference branch.
func (p *PackageReference) Branch() (string, error) {
	return getStringFunc(p.v.GetBranch)
}

// Type gets the package reference type.
func (p *PackageReference) Type() (string, error) {
	return getStringFunc(p.v.GetType)
}

// UniqueID gets a unique, formatted ID for the package reference.
func (p *PackageReference) UniqueID() (string, error) {
	return getStringFunc(p.v.GetUniqueId)
}

// IsExtension gets whether the package reference refers to an extension package.
func (p *PackageReference) IsExtension() (bool, error) {
	return p.v.GetIsExtension()
}

// FailedPackageReference describes unique attributes of a failed package.
type FailedPackageReference struct {
	PackageReference
	vf  *interop.ISetupFailedPackageReference
	vf2 *interop.ISetupFailedPackageReference2
	vf3 *interop.ISetupFailedPackageReference3
}

func newFailedPackageReference(v *interop.ISetupPackageReference, vf *interop.ISetupFailedPackageReference) *FailedPackageReference {
	p := &FailedPackageReference{
		PackageReference: PackageReference{
			v: v,
		},
		vf: vf,
	}

	runtime.SetFinalizer(p, (*FailedPackageReference).Close)
	return p
}

// Close releases any resources used by this FailedPackageReference immediately.
func (p *FailedPackageReference) Close() error {
	if p.v != nil {
		if p.vf != nil {
			p.vf.Release()
		}

		if p.vf2 != nil {
			p.vf2.Release()
		}

		if p.vf3 != nil {
			p.vf3.Release()
		}

		return p.PackageReference.Close()
	}

	return nil
}

// LogFilePath gets the path to the failed package log file.
func (p *FailedPackageReference) LogFilePath() (string, error) {
	if err := p.vf.ISetupFailedPackageReference2(&p.vf2); err != nil {
		return "", err
	}
	return getStringFunc(p.vf2.GetLogFilePath)
}

// Description gets a description of the package failure.
func (p *FailedPackageReference) Description() (string, error) {
	if err := p.vf.ISetupFailedPackageReference2(&p.vf2); err != nil {
		return "", err
	}
	return getStringFunc(p.vf2.GetDescription)
}

// Signature gets a unique signature of the package failure for error reporting.
func (p *FailedPackageReference) Signature() (string, error) {
	if err := p.vf.ISetupFailedPackageReference2(&p.vf2); err != nil {
		return "", err
	}
	return getStringFunc(p.vf2.GetSignature)
}

// Details gets the details of the package failure.
func (p *FailedPackageReference) Details() ([]string, error) {
	if err := p.vf.ISetupFailedPackageReference2(&p.vf2); err != nil {
		return nil, err
	}
	return p.vf2.GetDetails()
}

// AffectedPackages gets the list of packages that were not installed because of this package failure.
func (p *FailedPackageReference) AffectedPackages() ([]*PackageReference, error) {
	if err := p.vf.ISetupFailedPackageReference2(&p.vf2); err != nil {
		return nil, err
	}

	if affectedPackages, err := p.vf2.GetAffectedPackages(); err != nil {
		return nil, err
	} else {
		packages := make([]*PackageReference, len(affectedPackages))
		for i, ref := range affectedPackages {
			defer ref.Release()
			packages[i] = newPackageReference(ref)
		}
		return packages, nil
	}
}

// Action gets the attempted install action for the failed package.
func (p *FailedPackageReference) Action() (string, error) {
	if err := p.vf.ISetupFailedPackageReference3(&p.vf3); err != nil {
		return "", err
	}
	return getStringFunc(p.vf3.GetAction)
}

// ReturnCode gets the return code of the failed package.
func (p *FailedPackageReference) ReturnCode() (string, error) {
	if err := p.vf.ISetupFailedPackageReference3(&p.vf3); err != nil {
		return "", err
	}
	return getStringFunc(p.vf3.GetReturnCode)
}

// ProductReference describes unique attributes of a product package.
type ProductReference struct {
	PackageReference
	vp  *interop.ISetupProductReference
	vp2 *interop.ISetupProductReference2
}

func newProductReference(v *interop.ISetupPackageReference) *ProductReference {
	p := &ProductReference{
		PackageReference: PackageReference{
			v: v,
		},
	}

	runtime.SetFinalizer(p, (*ProductReference).Close)
	return p
}

// Close releases any resources used by this ProductReference immediately.
func (p *ProductReference) Close() error {
	if p.v != nil {
		if p.vp != nil {
			p.vp.Release()
		}

		if p.vp2 != nil {
			p.vp2.Release()
		}

		return p.PackageReference.Close()
	}

	return nil
}

// IsInstalled gets whether the product is completely installed.
func (p *ProductReference) IsInstalled() (bool, error) {
	if err := p.v.ISetupProductReference(&p.vp); err != nil {
		return false, err
	}

	return p.vp.GetIsInstalled()
}

// SupportsExtensions gets whether the product supports custom extensions e.g., VSIX packages.
func (p *ProductReference) SupportsExtensions() (bool, error) {
	if err := p.v.ISetupProductReference2(&p.vp2); err != nil {
		return false, err
	}

	return p.vp2.GetSupportsExtensions()
}
