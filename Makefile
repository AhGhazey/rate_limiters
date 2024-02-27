PROJECT=github.com/ahghazey/rate_limiter
export CURRENT_DIR=$(shell pwd)

all: mod fmt build

mod:
	go mod tidy

fmt:
	@echo "Running go fmt.."
	go fmt $(PROJECT)/...

build:
	@echo "Building:"
	go build -o rate_limiters cmd/api/main.go

run:
	@echo "Running:"
	go run cmd/api/main.go

install: build
	@echo "Installing:"
	go install $(PROJECT)/...

clean:
	@echo "Cleaning:"
	go clean

.PHONY: all mod fmt build run install clean