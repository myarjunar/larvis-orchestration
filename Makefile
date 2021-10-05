build:
	go build -o bin/larviscli main.go
setup:
	go mod download
test:
	go test -v ./...
