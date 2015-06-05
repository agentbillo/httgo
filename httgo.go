package main

import (
    "fmt"
//    "io/ioutil"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Yo %s!", r.URL.Path[1:])
}


func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
