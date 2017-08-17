package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func main() {
  port := "8080"
  if fromEnv := os.Getenv("PORT"); fromEnv != "" {
    port = fromEnv
  }

  server := http.NewServeMux()
  server.HandleFunc("/", salutations)
  log.Printf("Server listening on port %s", port)
  log.Fatal(http.ListenAndServe(":"+port, server))
}

func salutations(w http.ResponseWriter, r *http.Request) {
  log.Printf("Serving request: %s", r.URL.Path)
  host, _ := os.Hostname()
  fmt.Fprintf(w, "Salutations, globe!\n")
  fmt.Fprintf(w, "Hostname: %s\n", host)
}
