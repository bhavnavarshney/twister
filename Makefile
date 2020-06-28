lint:
	golangci-lint run ./...
test: # Runs unit tests
	go test -tags=unit ./...
test-integration:
	go test -tags=integration ./...
cover:
	go test -tags=unit -coverprofile cp.out ./...
	go tool cover -html=cp.out
build:
	rm -rf dist
	mkdir dist
	go build -o dist ./...
build-windows:
	wails build -x windows/amd64 -p
windows:
	GOOS=windows GOARCH=amd64 go build cmd/twister/main.go