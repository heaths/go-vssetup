package vssetup

import (
	"runtime"

	"github.com/heaths/go-vssetup/internal/interop"
)

// ErrorState contains information about failed and skipped packages,
// as well as log file paths.
type ErrorState struct {
	v  *interop.ISetupErrorState
	v2 *interop.ISetupErrorState2
	v3 *interop.ISetupErrorState3
}

func newErrorState(v *interop.ISetupErrorState) *ErrorState {
	e := &ErrorState{
		v: v,
	}

	runtime.SetFinalizer(e, (*ErrorState).Close)
	return e
}

// Close releases any resources used by this Instance immediately.
func (e *ErrorState) Close() error {
	if e.v != nil {
		// Call IUnknown.Release() but leave v assigned to avoid AV exceptions.
		e.v.Release()

		if e.v2 != nil {
			e.v2.Release()
		}

		if e.v3 != nil {
			e.v3.Release()
		}

		runtime.SetFinalizer(e, nil)
	}

	return nil
}

// FailedPackages gets an array of failed package references.
func (e *ErrorState) FailedPackages() ([]*FailedPackageReference, error) {
	if array, err := e.v.GetFailedPackages(); err != nil {
		return nil, err
	} else {
		packages := make([]*FailedPackageReference, len(array))
		for i, vf := range array {
			if v, err := vf.GetISetupPackageReference(); err == nil {
				packages[i] = newFailedPackageReference(v, vf)
			}
		}

		return packages, nil
	}
}

// SkippedPackages gets an array of package references that were not installed.
func (e *ErrorState) SkippedPackages() ([]*PackageReference, error) {
	if array, err := e.v.GetSkippedPackages(); err != nil {
		return nil, err
	} else {
		packages := make([]*PackageReference, len(array))
		for i, v := range array {
			packages[i] = newPackageReference(v)
		}

		return packages, nil
	}
}

// ErrorLogPath gets the path to the last errors log file if errors were logged.
func (e *ErrorState) ErrorLogPath() (string, error) {
	if err := e.v.ISetupErrorState2(&e.v2); err != nil {
		return "", err
	}

	return getStringFunc(e.v2.GetErrorLogFilePath)
}

// LogPath gets the path to the last log file.
func (e *ErrorState) LogPath() (string, error) {
	if err := e.v.ISetupErrorState2(&e.v2); err != nil {
		return "", err
	}

	return getStringFunc(e.v2.GetLogFilePath)
}
