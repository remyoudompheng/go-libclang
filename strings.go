package clang

// #include <clang-c/Index.h>
import "C"

func from_CXString(s C.CXString) string {
	p := C.clang_getCString(s)
	return C.GoString(p)
}

type CompletionString struct {
  ptr C.CXCompletionString
}

func (s CompletionString) NumChunks() uint {
	n := C.clang_getNumCompletionChunks(s.ptr)
	return uint(n)
}

func (s CompletionString) CompletionPriority() uint {
	n := C.clang_getCompletionPriority(s.ptr)
	return uint(n)
}
