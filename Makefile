SHORT_NAME ?= stk
VERSION ?= v1.0.0
BUILD_DATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GOVERSION = $(shell go env GOVERSION)
GOMOD := $(shell sed -n '/^module.*/s/module //p' go.mod)
GOARCH = $(shell go env GOARCH)
GOOS = $(shell go env GOOS)
BUILD_PATH = ./cmd/stk/main.go

LDFLAGS := \
    -X ${GOMOD}/pkg/version.BuildDateUTC="${BUILD_DATE}" \
	-X ${GOMOD}/pkg/version.Version="${VERSION}" \
	-X ${GOMOD}/pkg/version.GoVersion="${GOVERSION}" \
	-X ${GOMOD}/pkg/version.GitRepo="${GOMOD}" \
	-X ${GOMOD}/pkg/version.Platform="${GOOS}/${GOARCH}" \
	-s -w

BUILD_CMD := go build -a -ldflags "${LDFLAGS}"  -o ./build/${GOOS}/$(SHORT_NAME)${EXT} ${BUILD_PATH} || exit 1
RUN_CMD := go run ${BUILD_PATH} || exit 1

run:
	${RUN_CMD}

tidy:
	go mod tidy

vendor: tidy
	go mod vendor

bin:
	rm -fr ./build/${GOOS}
	mkdir -p ./build/${GOOS}/
	CGO_ENABLED=0 ${BUILD_CMD}

linux-bin:
	GOOS=linux make bin

windows-bin:
	GOOS=windows EXT=".exe" make bin

build:
	GOOS=darwin make bin
	GOOS=linux make bin

clean:
	rm -fr build