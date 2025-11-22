DEFAULT_GOAL: run

.PHONY: run tidy

tidy:
	go mod tidy

run: tidy
	go run -tags x11 ./...

build: tidy
	go build -tags x11 ./...
