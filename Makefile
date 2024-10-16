VERSION=$(shell git describe --tags --always --dirty)

# Build target for darwin and linux for both amd64 and arm64
build:
	@echo "Building for darwin and linux..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o credential-darwin-amd64 ./...
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o credential-darwin-arm64 ./...
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o credential-linux-amd64 ./...
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o credential-linux-arm64 ./...

# Build target for windows for both amd64 and arm64
build-windows:
	@echo "Building for windows..."
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o credential-windows-amd64.exe ./...
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags "-s -w" -o credential-windows-arm64.exe ./...

# Clean up binaries
clean:
	@echo "Cleaning up..."
	rm -f credential-darwin-amd64
	rm -f credential-darwin-arm64
	rm -f credential-linux-amd64
	rm -f credential-linux-arm64
	rm -f credential-windows-amd64.exe
	rm -f credential-windows-arm64.exe

.PHONY: build build-windows clean
