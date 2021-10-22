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
