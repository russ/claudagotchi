BINARY  := claudagotchi
VERSION ?= $(shell git describe --tags --dirty --always 2>/dev/null || echo dev)
LDFLAGS := -s -w -X main.version=$(VERSION)
GO      ?= go
DIST    := dist

.PHONY: all build run install tidy test fmt vet lint clean release \
        linux-arm64 linux-amd64 darwin-arm64 darwin-amd64

all: build

build:
	$(GO) build -trimpath -ldflags='$(LDFLAGS)' -o $(BINARY) .

run: build
	./$(BINARY)

install:
	$(GO) install -trimpath -ldflags='$(LDFLAGS)' .

tidy:
	$(GO) mod tidy

test:
	$(GO) test ./...

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

lint: fmt vet

clean:
	rm -rf $(BINARY) $(DIST)

$(DIST):
	@mkdir -p $(DIST)

linux-arm64: $(DIST)
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 $(GO) build -trimpath -ldflags='$(LDFLAGS)' -o $(DIST)/$(BINARY)-linux-arm64 .

linux-amd64: $(DIST)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GO) build -trimpath -ldflags='$(LDFLAGS)' -o $(DIST)/$(BINARY)-linux-amd64 .

darwin-arm64: $(DIST)
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 $(GO) build -trimpath -ldflags='$(LDFLAGS)' -o $(DIST)/$(BINARY)-darwin-arm64 .

darwin-amd64: $(DIST)
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 $(GO) build -trimpath -ldflags='$(LDFLAGS)' -o $(DIST)/$(BINARY)-darwin-amd64 .

release: linux-arm64 linux-amd64 darwin-arm64 darwin-amd64
