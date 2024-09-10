# Variables
BINARY_NAME=goCreate
VERSION=v1.0.0

# Build for Linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64

# Build for macOS
build-mac:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin-amd64

# Build for Windows
build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)-windows-amd64.exe

# Build for all platforms
build-all: build-linux build-mac build-windows

# Clean up
clean:
	rm -f $(BINARY_NAME)-linux-amd64 $(BINARY_NAME)-darwin-amd64 $(BINARY_NAME)-windows-amd64.exe

# Release build
release: build-all
	@echo "Creating release $(VERSION)"
	# You can add commands here to push the binaries to GitHub Releases
	# or upload them to any distribution platform.

.PHONY: build-linux build-mac build-windows build-all clean release
