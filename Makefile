# Variables
BIN_DIR = bin
CMD_DIR = cmd/textextract

# Default target
all: build

# Rule for building everything
build:
	go build -o $(BIN_DIR)/textextract ./cmd/textextract
	go build -o $(BIN_DIR)/win32resolve ./cmd/win32resolve

# Clean up build files
clean:
	rm -rf $(BIN_DIR)

# Ensure the bin directory exists
$(BIN_DIR):
	mkdir -p $(BIN_DIR)
