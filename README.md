# URL Shortener

A simple URL shortener written in Go.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:

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
Shortening a URL
To shorten a URL, make a POST request to http://localhost:8080/shorten with a JSON payload containing the original URL:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"url":"https://www.example.com"}' http://localhost:8080/shorten
```
The response will be a JSON object containing the short URL:

json
Copy code
{"short_url":"http://localhost:8080/abc123"}
Redirecting to the Original URL
To redirect to the original URL, make a GET request to the short URL:

bash
Copy code
curl -L http://localhost:8080/abc123
This will redirect you to the original URL.

### Built With
Go - The programming language used

### Authors
Pratim Bhosale - Initial work
See also the list of contributors who participated in this project.

### License
This project is licensed under the MIT License - see the LICENSE.md file for details.


