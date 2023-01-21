.PHONY: build
build:
	rm -rf build/
	go build -o build/ ./cmd/...

.PHONY: run
run:
	./build/web
