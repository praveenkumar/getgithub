# Go and compilation related variables
BUILD_DIR ?= out

ORG := github.com/praveenkumar
REPOPATH ?= $(ORG)/getgithub
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

ifeq ($(GOOS),windows)
	IS_EXE := .exe
endif
getgithub_BINARY ?= $(GOPATH)/bin/getgithub$(IS_EXE)

$(BUILD_DIR)/$(GOOS)-$(GOARCH):
	mkdir -p $(BUILD_DIR)/$(GOOS)-$(GOARCH)

$(BUILD_DIR)/darwin-amd64/getgithub: vendor $(BUILD_DIR)/$(GOOS)-$(GOARCH) ## Cross compiles the darwin executable and places it in $(BUILD_DIR)/darwin-amd64/getgithub
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build --installsuffix "static" -o $(BUILD_DIR)/darwin-amd64/getgithub

$(BUILD_DIR)/linux-amd64/getgithub: vendor $(BUILD_DIR)/$(GOOS)-$(GOARCH) ## Cross compiles the linux executable and places it in $(BUILD_DIR)/linux-amd64/getgithub
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build --installsuffix "static" -o $(BUILD_DIR)/linux-amd64/getgithub

$(BUILD_DIR)/windows-amd64/getgithub.exe: vendor $(BUILD_DIR)/$(GOOS)-$(GOARCH) ## Cross compiles the windows executable and places it in $(BUILD_DIR)/windows-amd64/getgithub
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build --installsuffix "static" -o $(BUILD_DIR)/windows-amd64/getgithub.exe

vendor:
	dep ensure -v

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	rm -rf vendor

.PHONY: cross ## Cross compiles all binaries
cross: $(BUILD_DIR)/darwin-amd64/getgithub $(BUILD_DIR)/linux-amd64/getgithub $(BUILD_DIR)/windows-amd64/getgithub.exe

.PHONY: build
build: $(BUILD_DIR) vendor
	go build -installsuffix "static" -o $(BUILD_DIR)/getgithub
	chmod +x $(BUILD_DIR)/getgithub
