build:
	@go build -o bin/main cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/main -port=${PORT} -host=${HOST} -redis-url=${REDIS_URL} -redis-port=${REDIS_PORT} -redis-password=${REDIS_PASSWORD} -cache-duration=${CACHE_DURATION}

clean:
	@rm -rf bin
