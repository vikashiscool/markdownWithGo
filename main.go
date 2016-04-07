package main

import (
  "net/http"
  "os"
  "github.com/russross/blackfriday"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
    port = "8080"
    }


    //http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))

     // calling http.Handle on the "/" pattern will act as a catch-all route, so we define that route last. http.FileServer returns an http.Handler so we use http.Handle to map a pattern string to a handler. The alternative method, http.HandleFunc, uses an http.HandlerFunc instead of an http.Handler. This may be more convenient, to think of handling routes via a function instead of an object.
    http.HandleFunc("/markdown", generateMarkdown)
    http.Handle("/", http.FileServer(http.Dir("public")))
    // http.ListenAndServe(":8080", nil)
    http.ListenAndServe(":"+port, nil)

}

func generateMarkdown(rw http.ResponseWriter, r *http.Request) {
  markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
  rw.Write(markdown)
}