.PHONY: build
build:
	go build -o app.exe ./cmd/main.go

.PHONY: test
test: 
	go test -v ./...
