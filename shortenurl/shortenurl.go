package shortenurl

import (
	"fmt"
	"math/rand"
)

// This function takes a url string and returns a shortened version of the url by appending a random number to the end of the fixed base url.
func ShortenURL(url string) string {
	// create a random string of size 6 from the alphabet
	s := ""
	for i := 0; i < 6; i++ {
		s += string(rand.Intn(26) + 97)
	}

	shortendURL := fmt.Sprintf("http://localhost:8080/%s", s)
	return shortendURL
}
