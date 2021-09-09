package interop

import "time"

type filetime struct {
	lowDateTime  uint32
	highDateTime uint32
}

// Copied from https://cs.opensource.google/go/x/sys/+/master:windows/types_windows.go;l=766;drc=0f9fa26af87c481a6877a4ca1330699ba9a30673.
func (ft *filetime) nanoseconds() int64 {
	// 100-nanosecond intervals since January 1, 1601
	nsec := int64(ft.highDateTime)<<32 + int64(ft.lowDateTime)
	// change starting time to the Epoch (00:00:00 UTC, January 1, 1970)
	nsec -= 116444736000000000
	// convert into nanoseconds
	nsec *= 100
	return nsec
}

func (ft *filetime) time() time.Time {
	return time.Unix(0, ft.nanoseconds())
}
