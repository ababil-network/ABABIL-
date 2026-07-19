
BINARY=ababild

build:
	go build -o build/$(BINARY) ./cmd/ababild

install:
	go install ./cmd/ababild

run:
	$(BINARY) start

test:
	go test ./...
