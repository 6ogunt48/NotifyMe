BUILD_DIR := ./build
DIST_DIR := ./dist

TARGETS := linux-amd64 windows-amd64 darwin-amd64 darwin-arm64
CURRENT_OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
CURRENT_ARCH := $(shell uname -m)

ifeq ($(CURRENT_OS),darwin)
    ifeq ($(CURRENT_ARCH),arm64)
        CURRENT_TARGET := darwin-arm64
    else
        CURRENT_TARGET := darwin-amd64
    endif
else ifeq ($(CURRENT_OS),linux)
    CURRENT_TARGET := linux-amd64
else ifeq ($(CURRENT_OS),windows)
    CURRENT_TARGET := windows-amd64
else
    $(error Unsupported operating system)
endif

.PHONY: build build-all build-current zip clean release

build: build-current

build-all: $(TARGETS)

build-current:
	go build -o $(BUILD_DIR)/checkemailbot-$(CURRENT_TARGET)

linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/checkemailbot-$@

windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/checkemailbot-$@.exe

darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/checkemailbot-$@

darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/checkemailbot-$@

zip: build-all
	mkdir -p $(DIST_DIR)
	zip -j $(DIST_DIR)/checkemailbot-linux-amd64.zip $(BUILD_DIR)/checkemailbot-linux-amd64
	zip -j $(DIST_DIR)/checkemailbot-windows-amd64.zip $(BUILD_DIR)/checkemailbot-windows-amd64.exe
	zip -j $(DIST_DIR)/checkemailbot-darwin-amd64.zip $(BUILD_DIR)/checkemailbot-darwin-amd64
	zip -j $(DIST_DIR)/checkemailbot-darwin-arm64.zip $(BUILD_DIR)/checkemailbot-darwin-arm64

clean:
	rm -rf $(BUILD_DIR) $(DIST_DIR)

release:
	./scripts/release.sh
