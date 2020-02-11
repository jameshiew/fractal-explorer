.PHONY: build

build:
	go build -v ./...

app:
	go build -v ./cmd/fractalexplorer

test:
	go test -v ./...

test-ci:
	go test -v -race -failfast ./...

ci-image:
	docker build -t registry.gitlab.com/jameshiew/fractal-explorer build/ci
