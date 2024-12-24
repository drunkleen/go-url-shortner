# URL Shortener API

This is a simple URL Shortener API built with Golang and Gin framework. This project demonstrates how to create a basic web service that can generate short URLs and redirect users to the original URLs. It's designed as an educational project for my [YouTube channel](https://www.youtube.com/@drunkleen).

## Features

- Generate a short URL for any given long URL.
- Redirect to the original URL using the short URL.
- Support for custom port configuration via command-line arguments.
- Generate UUIDs based on user IP addresses (demonstration purpose).

## Installation and Setup

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/drunkleen/go-url-shortner.git
   cd go-url-shortner
   ```

2. **Install Dependencies**:
   Make sure you have Go installed. Then run:

   ```bash
   go mod tidy
   ```


3. Create a `.env` file (optional) for environment variables, or use command-line flags.

   Example `.env` file:

   ```
   PORT=8080
   HOST=http://127.0.0.1/
   REDIS_URL=localhost
   REDIS_PORT=6379
   REDIS_PASSWORD=your_redis_password
   CACHE_DURATION=60
   DEBUG_MODE=true
   ```

4. **Run the Server**:

   ```bash
   make run
   ```
   or
   ```bash
   go run cmd/main.go
   ```

   The server will start on the configured port (default: `8080`).

### Configuration

You can configure the application in the following ways:

- **Environment variables**: Define configuration values in the `.env` file (or environment variables).
- **Command-line flags**: Override values when running the application (e.g., `-port 9090`).
- **Default values**: If not set, the application will use default values.

The following environment variables/flags are supported:

- `PORT` - The port to run the API on (default: `8080`).
- `HOST` - The base URL for the API (default: `http://127.0.0.1/`).
- `REDIS_URL` - The Redis server URL (default: `localhost`).
- `REDIS_PORT` - The Redis port (default: `6379`).
- `REDIS_PASSWORD` - The Redis password (default: empty).
- `CACHE_DURATION` - The duration in minutes for URL cache expiry (default: `60`).
- `DEBUG_MODE` - Set to `true` to enable debug mode (default: `false`).



## API Endpoints

### 1. **Welcome Endpoint**

- **URL**: `/`
- **Method**: `GET`
- **Description**: Returns a welcome message.

**Sample Response**:

```json
{
    "message": "Welcome to the URL Shortener API"
}
```

### 2. **Create Short URL**

- **URL**: `/create-short-url`
- **Method**: `POST`
- **Description**: Creates a short URL for a given long URL.

- **Request body**:
  ```json
  {
    "url": "https://www.example.com"
  }
  ```

- **Response**:
  ```json
  {
    "message": "short url created successfully",
    "short_url": "http://localhost:8080/abc12345"
  }
  ```

### 3. **Redirect Short URL**

- **URL**: `/:shortUrl`
- **Method**: `GET`
- **Description**: Redirects to the original URL corresponding to the short URL.

**Behavior**: If the `shortUrl` exists, it redirects to the `longUrl`. Otherwise, it returns an error.


### Example Usage

1. **Create a short URL**:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"url": "https://www.example.com"}' http://localhost:8080/create-short-url
   ```

2. **Redirect to the original URL**:
   After creating a short URL, visit it in a browser or use curl to test the redirection:
   ```bash
   curl -L http://localhost:8080/abc12345
   ```

### Error Handling

- If the input URL is missing or invalid, the server will return a `400 Bad Request` with an error message.
- If the short URL does not exist, the server will return a `404 Not Found`.

### Testing

You can run tests for the project using Go's built-in testing tools.

```bash
go test ./...
```

or 

```bash
make test
```

## Generating UUID from User IP

This project includes a demonstration feature where UUIDs are generated based on the user's IP address. This is shown as an educational example.

## Contributing

Feel free to contribute to this project by submitting issues or pull requests. This project is intended for educational purposes, so any improvements or suggestions are welcome.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Author

**DrunkLeen**\
YouTube: [DrunkLeen](https://www.youtube.com/@drunkleen)

