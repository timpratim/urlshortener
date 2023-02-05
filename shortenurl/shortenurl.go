package shortenurl

import (
	"fmt"
	"math/rand"
)

// This function takes a url string and returns a shortened version of the url by appending a random number to the end of the fixed base url.
func ShortenURL(url string) string {
	id := rand.Intn(10000)
	shortendURL := fmt.Sprintf("http://localhost:8080/%d", id)
	return shortendURL
}
