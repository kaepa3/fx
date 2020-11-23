GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=RadioHub
BINARY_PI=$(BINARY_NAME)_pi
BINARY_LINUX=$(BINARY_NAME)_lnx

all: test build pi
test:
	$(GOTEST) -v ./...
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
