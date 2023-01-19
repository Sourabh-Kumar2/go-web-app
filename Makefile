.PHONY: build
build:
	rm -rf build/
	go build -x -o build/ ./cmd/...

.PHONY: run
run:
	./build/web
