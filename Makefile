# Copyright (c) 2019 Clearmatics Technologies Ltd

.DEFAULT_GOAL := build

TARGET_BINARY = valuationprovider.so

GOFILES_ONLYROOT = $(shell find . -maxdepth 1 -type f -name '*.go')

PACKAGES = . ./binding

clean:
	@rm -f $(TARGET_BINARY)

build: modules
	@go build -buildmode=plugin -o $(TARGET_BINARY) ${GOFILES_ONLYROOT}

format:
	@gofmt -s -w .

test: modules
	go test $(PACKAGES)

coverage: modules
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out

modules: 
	@go mod download

lint:
	golangci-lint run -E gosec -E gofmt --deadline 4m0s
