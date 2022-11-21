.PHONY: all
all: build

PKG = main

ZIPNAME = build.zip

.PHONY: lambda
lambda:
	GOOS=linux GOARCH=amd64 go build -o main main.go
