
BINARY_NAME=gendiff
BINARY_PATH=bin/$(BINARY_NAME)
SRC_PATH=./cmd/gendiff

.PHONY: all build run lint fmt test clean

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(dir $(BINARY_PATH))
	go build -o $(BINARY_PATH) $(GOFLAGS) $(SRC_PATH)
	@echo "$(BINARY_NAME) built successfully to $(BINARY_PATH)"

run: build
	@echo "Running $(BINARY_NAME) with default path $(DEFAULT_PATH) and ARGS..."
	$(BINARY_PATH) $(DEFAULT_PATH) $(ARGS)

lint:
	@echo "Running linters..."
	golangci-lint run ./...

fmt:
	@echo "Formatting code..."
	goimports -w .

test:
	@go mod tidy
	@echo "Cleaning testdata..."
	@rm -rf ./testdata
	@echo "Running tests..."
	go test -v ./...

clean:
	rm -f $(BINARY_PATH)