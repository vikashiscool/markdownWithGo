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
    http.HandleFunc("/markdown", generateMarkdown)
    http.Handle("/", http.FileServer(http.Dir("public")))
    // http.ListenAndServe(":8080", nil)
    http.ListenAndServe(":"+port, nil)

}

func generateMarkdown(rw http.ResponseWriter, r *http.Request) {
  markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
  rw.Write(markdown)
}