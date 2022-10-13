package main

import (
    "fmt"
    "log"
    "net/http"
)

func helloHandler( w http.ResponseWriter, r *http.Request ) {
    if r.URLPath != "hello" {
        http.Error( w, "404 not found.", http.StatusNotFound )
        return
    }

    if r.Method != "GET" {
        http.Error( w, "Method is not supported.", http.StatusNotFound )
        return
    }
    fmt.Fprintf( w, "Hello!" )
}

func main() {
    http.HandleFunc( "/hello", helloHandler )

    http.HandleFunc( "/", func( w http.ResponseWriter, r *http.Request ) {
        fmt.Fprintf( w, "HomePage!" )
    } )

    fmt.Println( "Starting web server at port 8080" )

    err := http.ListenAndServe( ":8080", nil )
    if err != nil {
        log.Fatal( err )
    }
}
