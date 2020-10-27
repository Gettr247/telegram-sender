BIN_PATH=bin
BIN_NAME=telegram-sender
BIN_ARM=$(BIN_NAME)_linux_arm
BIN_AMD64=$(BIN_NAME)_linux_amd64
PACKAGES ?= $(shell go list ./...)

# Go parameters
GO=go
GOBUILD=$(GO) build
GOCLEAN=$(GO) clean

all: lint build

build:
	$(GOBUILD) -o $(BIN_PATH)/$(BIN_NAME) -v

lint:
	@which golangci-lint > /dev/null; if [ $$? -ne 0 ]; then \
		GO111MODULE=off $(GO) get -u github.com/golangci/golangci-lint/cmd/golangci-lint@v1.31.0; \
	fi
	golangci-lint run -v

fmt:
	$(GO) fmt $(PACKAGES)

clean:
	$(GOCLEAN) -i ./...
	rm -f $(BIN_PATH)/$(BIN_NAME)
	rm -f $(BIN_PATH)/$(BIN_ARM)
	rm -f $(BIN_PATH)/$(BIN_AMD64)

run:
	$(GOBUILD) -o $(BIN_PATH)/$(BIN_NAME) -v
	./$(BIN_PATH)/$(BIN_NAME)

# Cross compilation
build-linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BIN_PATH)/$(BIN_AMD64) -v

# Cross compilation
build-linux-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -o $(BIN_PATH)/$(BIN_ARM) -v
