package clang

// #cgo LDFLAGS: -L/usr/lib/llvm -lclang
// #include <clang-c/Index.h>
import "C"

type CodeCompleteResults struct {
	Items []CompletionString
	ptr   *C.CXCodeCompleteResults
}

func (res *CodeCompleteResults) Dispose() {
	if res == nil {
		return
	}
	res.Items = nil
	if res.ptr != nil {
		C.clang_disposeCodeCompleteResults(res.ptr)
		res.ptr = nil
	}
}
