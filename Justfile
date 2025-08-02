build:
	go build -v ./...

run:
	go run -v ./cmd/fractalexplorer

app:
	go build -v ./cmd/fractalexplorer

test:
	go test -v ./...

test-ci:
	go test -v -race -failfast ./...

upgrade-deps:
	go get -u -t -v ./...
