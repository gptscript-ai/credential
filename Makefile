build:
	CGO_ENABLED=0 go build -o bin/gptscript-go-tool -ldflags "-s -w"