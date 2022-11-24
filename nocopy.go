package nocopy

import (
	"unsafe"
)

// noCopy may be added to structs which must not be copied
// after the first use.
//
// See https://golang.org/issues/8005#issuecomment-190753527
// for details.
//
// Note that it must not be embedded, due to the Lock and Unlock methods.
type NoCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*NoCopy) Lock()   {}
func (*NoCopy) Unlock() {}

// BytesToString convert []byte to string via unsafe.Pointer
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes convert string to []byte via unsafe.Pointer,
//
// WARNING: if the string s is allocated in constant pool, do not change the returned []byte value
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
