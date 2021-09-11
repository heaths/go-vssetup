package errors

import (
	"testing"

	"github.com/go-ole/go-ole"
)

func Test_NotImplemented(t *testing.T) {
	tests := []struct {
		name  string
		err   error
		wantE bool
	}{
		{
			name:  "with error",
			err:   ole.NewError(ole.E_NOTIMPL),
			wantE: true,
		},
		{
			name: "without error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NotImplemented(tt.err)
			if got.Code() != ole.E_NOTIMPL {
				t.Errorf("NotImplemented() code = %#x, expected %#x", got.Code(), ole.E_NOTIMPL)
			} else if got.Error() != "Not implemented" {
				t.Errorf(`NotImplemented() = "%s", expected "%s"`, got.Error(), "Not implemented")
			} else if (got.SubError() != nil) != tt.wantE {
				t.Errorf("NotImplemented() error = %v, expected? %t", got.SubError(), tt.wantE)
			}
		})
	}
}

func Test_ComError_Compatibility(t *testing.T) {
	tests := []struct {
		name  string
		err   ComError
		wantE bool
	}{
		{
			name: "NotImplemented",
			err:  NotImplemented(nil),
		},
		{
			name:  "NotImplemented with SubError",
			err:   NotImplemented(ole.NewError(ole.E_NOTIMPL)),
			wantE: true,
		},
		{
			name: "ole.OleError",
			err:  ole.NewErrorWithDescription(ole.E_NOTIMPL, "Not implemented"),
		},
		{
			name:  "ole.OleError with SubError",
			err:   ole.NewErrorWithSubError(ole.E_NOTIMPL, "Not implemented", ole.NewError(ole.E_NOTIMPL)),
			wantE: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Code() != ole.E_NOTIMPL {
				t.Errorf("Code() = %#x, expected %#x", tt.err.Code(), ole.E_NOTIMPL)
			} else if tt.err.Error() == "" {
				t.Errorf(`Error() = "", expected "%s"`, tt.err.Error())
			} else if (tt.err.SubError() != nil) != tt.wantE {
				t.Errorf("SubError() = %v, expected? %t", tt.err.SubError(), tt.wantE)
			}
		})
	}
}
