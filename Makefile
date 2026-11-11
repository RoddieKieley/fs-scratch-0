BIN := fs-scratch-0

.PHONY: build
build:
	go build -o $(BIN) .
