package main

import (
  "net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("ok"))
}

func main() {
  http.HandleFunc("/health", healthHandler)
  http.ListenAndServe(":8080", nil)
}
