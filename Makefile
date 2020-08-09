all: lint test generate build

.PHONY: lint test test-integration generate cover build install-frontend build-frontend

lint:
	golangci-lint run ./...

test: # Runs unit tests
	go test -tags=unit ./...

test-integration:
	go test -tags=integration ./...

install-frontend:
	cd frontend && yarn install

test-frontend: install-frontend
	cd frontend && yarn test

build-frontend: install-frontend
	cd frontend && yarn build

generate: build-frontend
	go generate && echo '//+build ignorelint\n' | cat - assets.go > temp && mv temp assets.go
	
cover:
	go test -tags=unit -coverprofile cp.out ./...
	go tool cover -html=cp.out
	
build: generate
	rm -rf dist
	mkdir dist
	go build -tags=ignorelint -o dist