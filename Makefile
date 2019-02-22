OS = $(shell go env GOOS)
ARCH = $(shell go env GOARCH)

GOPATH := $(PWD):$(GOPATH)
BUILDMODE = c-shared

PKGDIR = ./pkg/$(OS)_$(ARCH)_shared

CFLAGS = -fPIC -Wall -I$(PKGDIR) -L$(PKGDIR)
LDFLAGS = -L. -lgoc -Wl,-rpath=.


all:libgoc main

main:
	@gcc $(CFLAGS) -o bin/main csrc/main.c $(LDFLAGS)

libgoc:
	@go install -buildmode=$(BUILDMODE) libgoc
