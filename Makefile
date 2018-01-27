.PHONY: emrtr-dep emrtr-goget test

emrtr-dep:
	dep ensure
	go build -buildmode=plugin -o emrtr.so emrtr/*

emrtr-goget:
	go get -t -buildmode=plugin ./emrtr

test:
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l emrtr/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
