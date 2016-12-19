package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/pressly/chi"
    chimiddle "github.com/pressly/chi/middleware"
)

func main() {
    r := chi.NewRouter()
    
    //Setup middleware
    r.Use(chimiddle.RequestID)
    r.Use(chimiddle.RealIP)
    r.Use(chimiddle.Logger)
    r.Use(chimiddle.Recoverer)
    r.Use(chimiddle.CloseNotify)
    
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, world")
    })
    
    log.Fatal(http.ListenAndServe(":3000", r))
}