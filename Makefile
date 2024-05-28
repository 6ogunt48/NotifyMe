# Set the Go binary name
BINARY := checkemailbot

# Set the Go build flags
LDFLAGS := -ldflags="-s -w"

# Linux targets
linux:
    GOOS=linux GOARCH=amd64 $(MAKE) build

linux-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)

linux-run: linux-build
	./$(BINARY)

linux-clean:
	rm -f $(BINARY)

# macOS targets
darwin:
    GOOS=darwin GOARCH=amd64 $(MAKE) build

darwin-build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)

darwin-run: darwin-build
	./$(BINARY)

darwin-clean:
	rm -f $(BINARY)

# Clean up binaries
clean: linux-clean darwin-clean

# Print the help message
help:
	@echo "Usage:"
	@echo "  make linux         Build the binary for amd64/linux"
	@echo "  make linux-run     Build and run the binary for amd64/linux"
	@echo "  make linux-clean   Clean up the amd64/linux binary"
	@echo "  make darwin        Build the binary for amd64/macOS"
	@echo "  make darwin-run    Build and run the binary for amd64/macOS"
	@echo "  make darwin-clean  Clean up the amd64/macOS binary"
	@echo "  make clean         Clean up all binaries"
	@echo "  make help          Print this help message"