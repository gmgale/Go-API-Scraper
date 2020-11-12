# Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
    BINARY_NAME=goapiscraper.exe

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/gorilla/mux

# Create a fresh docker build and run
docker-new:
	docker-compose up --build --force-recreate
# Build and run the API on the localhost server (localhost flag)
local-run:
	api/goapiscraper.exe -host localhost

	
	
