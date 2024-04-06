default: test

ci: depsdev test

depsdev:
	go install github.com/Songmu/gocredits/cmd/gocredits@latest
	go install github.com/ogen-go/ogen/cmd/ogen@latest

build:
	go generate ./...
	go mod tidy

test: build
	go test ./...

lint:
	golangci-lint run ./...

prerelease_for_tagpr: depsdev
	gocredits . -w
	git add CHANGELOG.md CREDITS go.mod go.sum

.PHONY: depsdev build test lint prerelease_for_tagpr
