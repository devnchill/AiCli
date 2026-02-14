BINARY := build/aicli

MAIN := ./cmd/aicli

GOFLAGS ?= 

all: build

build:
	mkdir build
	go build $(GOFLAGS) -o $(BINARY) $(MAIN)

run: build
	./$(BINARY)

clean:
	rm -f $(BINARY)

install:
	go install $(GOFLAGS) $(MAIN)

