package shortenurl

import (
	"fmt"
	"math/rand"
)

// This function takes a url string and returns a shortened version of the url.
func ShortenURL(url string) string {
	// create a random string of size 6 from the alphabet
	s := ""
	//rand.Intn(26) returns a random number between 0 and 25. 97 is the ascii value of 'a'. So rand.Intn(26) + 97 returns a random lowercase letter.
	for i := 0; i < 6; i++ {
		s += string(rand.Intn(26) + 97)
	}

	shortendURL := fmt.Sprintf("http://localhost:8080/%s", s)
	return shortendURL
}
