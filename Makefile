default: test

ci: depsdev test

depsdev:
	go install github.com/Songmu/gocredits/cmd/gocredits@latest
	go install github.com/ogen-go/ogen/cmd/ogen@latest

fetch-spec:
	curl https://docs.fincode.jp/assets/files/api/fincode-openapi.yml?`date +%s` -o spec/fincode-openapi.yml

fetch-and-patch-spec: fetch-spec
	env GOWORK=off go run -modfile=misc/openapi-yaml-patch/go.mod misc/openapi-yaml-patch/main.go

build:
	go generate ./...
	go mod tidy

test: build
	go test ./... -coverprofile=coverage.out -covermode=count

test-scenario: build
	go test ./... -v -run TestScenario

lint:
	golangci-lint run ./...

prerelease_for_tagpr: depsdev
	gocredits . -w
	git add CHANGELOG.md CREDITS go.mod go.sum

.PHONY: depsdev build test lint prerelease_for_tagpr
