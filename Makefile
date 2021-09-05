# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOMOD=$(GOCMD) mod
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
    
all: build
build:
		rm -f xplorer && $(GOBUILD) -v -ldflags="-extldflags=-static" -o "xplorer" main.go
test: 
		$(GOTEST) -v ./...
tidy:
		$(GOMOD) tidy