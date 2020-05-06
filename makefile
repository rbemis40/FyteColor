UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	DEF_OS := linux
	DEF_NAME := linux
endif
ifeq ($(UNAME_S),Darwin)
	DEF_OS := darwin
	DEF_NAME := macos
endif

default:
	env GOOS=$(DEF_OS) GOARCH=amd64 go build -o bin/$(DEF_NAME)/64/FyteColor cmd/main.go

all: macos linux windows

macos:
	env GOOS=darwin GOARCH=amd64 go build -o bin/macos/64/FyteColor cmd/main.go
	env GOOS=darwin GOARCH=386 go build -o bin/macos/32/FyteColor cmd/main.go
	
windows:
	env GOOS=windows GOARCH=amd64 go build -o bin/win/64/FyteColor.exe cmd/main.go
	env GOOS=windows GOARCH=386 go build -o bin/win/32/FyteColor.exe cmd/main.go
	
linux:	
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/64/FyteColor cmd/main.go
	env GOOS=linux GOARCH=386 go build -o bin/linux/32/FyteColor cmd/main.go
