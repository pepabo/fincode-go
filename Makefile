depsdev:
	go install github.com/ogen-go/ogen/cmd/ogen@latest

build:
	go generate ./...
	go mod tidy

test: build
	go test ./...
