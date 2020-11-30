# Copyright (c) 2019 Clearmatics Technologies Ltd

.DEFAULT_GOAL=build

TARGET_BINARY=valuationprovider.so

GOFILES_ONLYROOT=$(shell find . -maxdepth 1 -type f -name '*.go')

PACKAGES=../binding

TMP_DIRECTORY=${GOPATH}/src/github.com/clearmatics/valuation-provider-bindings/tmp_contracts
ABIGEN_CMD="${GOPATH}/src/github.com/clearmatics/autonity/build/bin/abigen"
VALUATION_PROVIDER_CONTRACT=ValuationProvider.sol
OUTPUT_JSON=${GOPATH}/src/github.com/clearmatics/valuation-provider-bindings/binding/valuationprovider.go

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

autogen:
	# cd $(AUTONITY) && go install ./cmd/abigen
	mkdir -p $(TMP_DIRECTORY)
	git clone git@github.com:clearmatics/valuation-provider-contract.git $(TMP_DIRECTORY)
	cd $(TMP_DIRECTORY); $(ABIGEN_CMD) --sol=$(TMP_DIRECTORY)/contracts/$(VALUATION_PROVIDER_CONTRACT) --pkg=valuationprovider --out=${OUTPUT_JSON}
lint:
	golangci-lint run -E gosec -E gofmt --deadline 4m0s
