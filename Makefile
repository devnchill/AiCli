BINARY := build/aicli

MAIN := ./cmd/aicli

SRC := $(shell find . -type f -name '*.go')

all: $(BINARY)

$(BINARY): $(SRC)
	mkdir -p build
	go build -o $(BINARY) $(MAIN)

.PHONY: clean run

run: $(BINARY)
	./$(BINARY)

clean:
	rm -rf build

