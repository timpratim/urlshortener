# URL Shortener

A simple URL shortener written in Go.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
Things you need on your machine.

- Go
- Any text editor or IDE

### Installing

Clone the repository to your local machine.

```
git clone https://github.com/timpratim/url-shortener.git
```

### Change into the project directory:

```bash
cd url-shortener
```

### Build the project:

```go
go build
```

### Run the project:

```go
go run main.go
```

### Usage
- Shortening a URL
To shorten a URL, make a POST request to http://localhost:8080/shorten with a JSON payload containing the original URL:

```bash
curl -X POST -d "url=https://twitter.com/BhosalePratim" http://localhost:8080/shorten
```
- The response will be a JSON object containing the short URL:

```json
{"short_url":"http://localhost:8080/abc123"}
```

### Redirecting to the Original URL
To redirect to the original URL, make a GET request to the short URL:

This will redirect you to the original URL.

### Built With
Go - The programming language.

### Authors
Pratim Bhosale



