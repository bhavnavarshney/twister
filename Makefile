.PHONY: lint test generate cover build
lint:
	golangci-lint run ./...
test: # Runs unit tests
	go test -tags=unit ./...
test-integration:
	go test -tags=integration ./...
generate:
	go generate
cover:
	go test -tags=unit -coverprofile cp.out ./...
	go tool cover -html=cp.out
build: generate
	rm -rf dist
	mkdir dist
	go build -o dist