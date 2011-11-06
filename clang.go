package clang

// #cgo LDFLAGS: -L/usr/lib/llvm -lclang
// #include <clang-c/Index.h>
import "C"

type Index struct {
	ptr C.CXIndex
}

func CreateIndex(exclude_decl_from_pch, display_diags bool) Index {
	var arg1, arg2 C.int
	if exclude_decl_from_pch {
		arg1 = 1
	} else {
		arg1 = 0
	}
	if display_diags {
		arg2 = 1
	} else {
		arg1 = 0
	}
	p := C.clang_createIndex(arg1, arg2)
	return Index{p}
}

func (index *Index) Dispose() {
	if index != nil && index.ptr != nil {
		C.clang_disposeIndex(index.ptr)
		index.ptr = nil
	}
}

type CodeCompleteResults struct {
	Items []CompletionString
	ptr   *C.CXCodeCompleteResults
}

type TranslationUnit struct {
	ptr *C.CXTranslationUnit
}

func ParseTranslationUnit(idx Index,
	filename string,
	cmdline_args []string,
	unsaved_files []UnsavedFile,
	options TUOptions) {
}
