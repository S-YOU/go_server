package main

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/markdown", GenerateMarkdown)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":" + port)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
