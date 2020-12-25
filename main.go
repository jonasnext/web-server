// https://golang.org/doc/articles/wiki/
// https://golang.org/pkg/net/http

package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "👋 🙌 %s", r.URL.Path[1:])
}

func main() {
    log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
