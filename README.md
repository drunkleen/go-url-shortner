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

3. **Run the Server**:
   By default, the server will run on port `8080`:

   ```bash
   make run
   ```
   or
   ```bash
   go run cmd/main.go
   ```

4. **Use a Custom Port**:
   To specify a custom port, pass the `-port` argument:

   ```bash
   go run cmd/main.go -port 2048
   ```

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

**Request Body**:

```json
{
    "url": "https://example.com"
}
```

**Sample Response**:

```json
{
    "shortUrl": "abc123"
}
```

### 3. **Redirect Short URL**

- **URL**: `/:shortUrl`
- **Method**: `GET`
- **Description**: Redirects to the original URL corresponding to the short URL.

**Behavior**: If the `shortUrl` exists, it redirects to the `longUrl`. Otherwise, it returns an error.

## Generating UUID from User IP

This project includes a demonstration feature where UUIDs are generated based on the user's IP address. This is shown as an educational example.

## Contributing

Feel free to contribute to this project by submitting issues or pull requests. This project is intended for educational purposes, so any improvements or suggestions are welcome.

## Author

**DrunkLeen**\
YouTube: [DrunkLeen](https://www.youtube.com/@drunkleen)

