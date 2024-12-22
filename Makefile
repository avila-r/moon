BIN_DIR = ./bin
BIN_NAME = moon

build:
	@echo "Building the application..."
	go build -o ./bin/moon

clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

test:
	go test ./...
