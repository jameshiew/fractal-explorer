.PHONY: build

build:
	go build -v ./cmd/fractalexplorer

ci-image:
	docker build -t registry.gitlab.com/jameshiew/fractal-explorer build/ci
