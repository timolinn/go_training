package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
	"time"
)

var books []Book

// Book model
type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}

func loadBooks(w http.ResponseWriter, r *http.Request) {
	bk, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bk)
}

func create(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./assets/form.html")
	tmpl.Execute(w, nil)
}

func postBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "must be a post request", http.StatusInternalServerError)
		return
	}
	book := Book{
		Name:        r.FormValue("name"),
		Author:      r.FormValue("author"),
		PublishedAt: time.Now().Local().String(),
	}
	books = append(books, book)

	fmt.Fprintf(w, "New Book added: %s\n", book.Name)
}

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request received %v\n", r.URL.Path)
		next(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", logger(loadBooks))

	fs := http.FileServer(http.Dir("assets/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	// Forms
	mux.HandleFunc("/new", logger(create))
	mux.HandleFunc("/create", logger(postBook))
	// Middlewares
	http.ListenAndServe(":8000", mux)
	// Database
}

/*package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"
)

// Book model
type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}

func loadBooks(w http.ResponseWriter, r *http.Request) {
	book := []Book{
		Book{"The Silmarilion", "JRR Tolkein", time.Now().Local().String()},
		Book{"The Hobbit", "JRR Tolkein", time.Now().Local().String()},
	}

	bk, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bk)
}

func show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	tmpl := template.Must(template.ParseFiles("./assets/index.html"))
	book := Book{"The Silmarilion", "JRR Tolkein", time.Now().Local().String()}
	tmpl.Execute(w, book)
}

func main() {
	// Serving JSON
	mux := http.NewServeMux()
	mux.HandleFunc("/books", loadBooks)

	fs := http.FileServer(http.Dir("assets/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	// Templating
	mux.HandleFunc("/book", show)
	http.ListenAndServe(":8000", mux)
	// Forms
	// Middlewares
	// Database
}*/

/*package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Book model
type Book struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}

func loadBooks(w http.ResponseWriter, r *http.Request) {
	book := []Book{
		Book{"The Silmarilion", "JRR Tolkein", time.Now().Local().String()},
		Book{"The Hobbit", "JRR Tolkein", time.Now().Local().String()},
	}

	bk, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bk)
}

func main() {
	// Serving JSON
	mux := http.NewServeMux()
	mux.HandleFunc("/books", loadBooks)

	fs := http.FileServer(http.Dir("assets/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe(":8000", mux)
	// Templating
	// Forms
	// Middlewares
	// Database
}
*/
/*
package main

import (
	"fmt"
	"net/http"
)

type books map[string]string

func (b books) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/home":
		fmt.Fprintf(w, "Welcome to books server. You came via %s\n", r.URL.String())
	case "/about":
		fmt.Fprintf(w, "All about our library of books. You came via %s\n", r.URL.String())
	case "/book":
		item := r.URL.Query().Get("item")
		book, ok := b[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "book not found for: %q\n", item)
			return
		}
		fmt.Fprintf(w, "The book is %s\n", book)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "url not found! %s\n", r.URL.String())
	}
}

func main() {
	// Http Server
	// http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello, Go HTTP\n")
	// })

	// http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "About, HTTP in Go\n")
	// })

	// Routing
	bk := books{}
	bk["jrrtolkein"] = "The Silmarilion"
	fmt.Println("server listening on port 8000")
	mux := http.NewServeMux()
	mux.Handle("/book", bk)

	fs := http.FileServer(http.Dir("assets/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8000", mux)
}

*/
