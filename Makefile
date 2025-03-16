.PHONY: run build test clean

# Default target
all: build

# Run the application
run:
	go run cmd/api/main.go

# Build the application
build:
	go build -o bin/app cmd/api/main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Generate API documentation
# docs:
# 	swag

# Help target
help:
	@echo "Available targets:"
	@echo "  run      - Run the application"
	@echo "  build    - Build the application"
	@echo "  test     - Run tests"
	@echo "  clean    - Clean build artifacts"
	@echo "  docs     - Generate API documentation"
