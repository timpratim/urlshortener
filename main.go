// Description: This program takes a url as input and returns a shortened version of the url. The shortened version of the url is a random number appended to the base url. The shortened url is stored in a database and the original url is returned when the shortened url is requested.
package main

import (
	"fmt"
	"github.com/timpratim/urlshortener/shortenurl"
	"net/http"
)

func main() {
	//db is a pointer to a database connection.
	db, err := shortenurl.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	//AutoMigrate takes a pointer to a struct and creates a table in the database.
	db.AutoMigrate(&shortenurl.URL{})
	//HandleFunc takes a path and a function as arguments. The function is called when the path is requested.
	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		original := r.FormValue("url")
		shortened := shortenurl.ShortenURL(original)
		fmt.Println(shortened)
		//Create takes a pointer to a struct and creates a new record in the database.
		db.Create(&shortenurl.URL{Original: original, Shortened: shortened})

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortenurl.RedirectURL(db, w, r)
	})
	http.ListenAndServe(":8080", nil)
}

//https://www.youtube.com/watch?v=4KfuQwB5rIs&t=1s
//curl -X POST -d "url=https://www.youtube.com/watch?v=4KfuQwB5rIs&t=1s" http://localhost:8080/shorten
