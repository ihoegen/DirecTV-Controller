project_dir = $(shell pwd)

.PHONY: all test fmt vet dep build build-docker

all: test build

# Run tests
test: dep vet fmt
	go test ./pkg/... ./cmd/... -coverprofile cover.out

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...

# Get dependencies
dep: 
	go get -d ./...

# Build all binaries
build: dep
	@mkdir -p bin/
	@GOBIN=$(project_dir)/bin/ go install ./cmd/...
	@echo "Binaries built to $(project_dir)/bin"
