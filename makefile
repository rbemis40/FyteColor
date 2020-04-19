default:
	go build -o bin/macos/64/FyteColor cmd/main.go

all:
	env GOOS=darwin GOARCH=amd64 go build -o bin/macos/64/FyteColor cmd/main.go
	env GOOS=darwin GOARCH=386 go build -o bin/macos/32/FyteColor cmd/main.go
	env GOOS=windows GOARCH=amd64 go build -o bin/win/64/FyteColor.exe cmd/main.go
	env GOOS=windows GOARCH=386 go build -o bin/win/32/FyteColor.exe cmd/main.go
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/64/FyteColor cmd/main.go
	env GOOS=linux GOARCH=386 go build -o bin/linux/32/FyteColor cmd/main.go