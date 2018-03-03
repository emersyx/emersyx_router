emrtr.so: goget
	go build -buildmode=plugin -o emrtr.so emrtr/*

.PHONY: goget
goget:
	go get emersyx.net/emersyx_apis/emcomapi

.PHONY: test
test: emrtr.so
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l emrtr/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
