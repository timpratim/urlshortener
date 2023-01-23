package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
)

type URL struct {
	ID        uint   `gorm:"primary_key"`
	Original  string `gorm:"not null"`
	Shortened string `gorm:"not null"`
}

func shortenURL(url string) string {
	id := rand.Intn(10000)
	shortendURL := fmt.Sprintf("http://localhost:8080/%d", id)
	return shortendURL
}

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func redirectURL(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[1:]
	var url URL
	db.First(&url, id)
	http.Redirect(w, r, url.Original, http.StatusFound)

}

func main() {
	db, err := Connect()
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&URL{})
	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		original := r.FormValue("url")
		shortened := shortenURL(original)
		fmt.Printf(shortened)
		db.Create(&URL{Original: original, Shortened: shortened})

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redirectURL(db, w, r)
	})
	http.ListenAndServe(":8080", nil)
}

//https://www.youtube.com/watch?v=4KfuQwB5rIs&t=1s
//curl -X POST -d "url=https://www.youtube.com/watch?v=4KfuQwB5rIs&t=1s" http://localhost:8080/shorten
