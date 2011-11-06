include $(GOROOT)/src/Make.inc

TARG=github.com/remyoudompheng/go-libclang

CGOFILES=clang.go\
	 completion.go\
	 strings.go\

GOFILES=enums.go\
	files.go

include $(GOROOT)/src/Make.pkg

.PHONY: gofmt
gofmt:
	gofmt -l -s -w *.go
