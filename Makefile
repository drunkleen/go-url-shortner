build:
	@go build -o bin/main cmd/main.go

test:
	go test -v ./...

run: build
	@./bin/main

clean:
	@rm -rf bin