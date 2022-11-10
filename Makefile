# init command params
GO      := go
GOROOT  := $(shell $(GO) env GOROOT)
GOMOD   := $(GO) mod
GOBUILD := $(GO) build
GOTEST  := $(GO) test -gcflags="-N -l"
GOPKGS  := $$($(GO) list ./...| grep -vE "vendor")

# make all
all: test

# set go env
set-env:
	$(GO) env -w GO111MODULE=on
	$(GO) env -w GOPROXY=https://goproxy.io,direct
	$(GO) env -w GONOSUMDB=off

# make prepare
prepare: gomod
gomod: set-env
	$(GOMOD) tidy

# make test
test: prepare test-case
test-case:
	$(GOTEST) -cover $(GOPKGS)

# avoid filename conflict and speed up build
.PHONY: all prepare test 
