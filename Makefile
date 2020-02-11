.PHONY: build

build:
	go build -v ./...

app:
	go build -v ./cmd/fractalexplorer

test:
	go test -v ./...

test-ci:
	go test -v -race -failfast ./...

upgrade-deps:
	go get -u -t -v ./...

ci-image:
	docker build -t registry.gitlab.com/jameshiew/fractal-explorer build/ci

lint-go: ci-image
	docker run --rm -v $(shell pwd):/go/src/github.com/jameshiew/fractal-explorer registry.gitlab.com/jameshiew/fractal-explorer golangci-lint -v run
