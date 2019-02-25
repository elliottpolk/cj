BIN=cj
PKG=github.com/elliottpolk/cj
VERSION=`cat .version`
GOOS?=linux
BUILD_DIR=./build/bin

M = $(shell printf "\033[34;1m◉\033[0m")

default: clean build ;                                              @ ## defaulting to clean and build

.PHONY: all
all: clean test build install

.PHONEY: build-dir
build-dir: ;
	@[ ! -d "${BUILD_DIR}" ] && mkdir -vp "${BUILD_DIR}/public" || true

.PHONY: build
build: build-dir; $(info $(M) building ...)                        	@ ## build the binary
	@GOOS=$(GOOS) go build -ldflags "-X main.version=$(VERSION) -X main.compiled=$(date +%s)" -o ./build/bin/$(BIN) cmd/main.go
	@chmod +x ./build/bin/$(BIN)

.PHONY: install
install: ; $(info $(M) installing locally...)                       @ ## install the binary locally
	@GOOS=$(GOOS) go build -ldflags "-X main.version=$(VERSION) -X main.compiled=$(date +%s)" -o $(GOPATH)/bin/$(BIN) cmd/main.go
	@chmod +x $(GOPATH)/bin/$(BIN)

.PHONY: test
test: ; $(info $(M) running unit tests ...)                         @ ## run the unit tests
	@go test -v -cover ./...

.PHONY: clean
clean: ; $(info $(M) running clean ...)                             @ ## clean up the old build dir
	@rm -vrf build || true
	@rm -v $(GOPATH)/bin/$(BIN) || true
	@go clean

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

