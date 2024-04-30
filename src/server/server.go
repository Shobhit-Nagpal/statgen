package server

import (
  "os"
  "fmt"
  "net/http"
)

func Start(dir, port string) {

  _, err := os.Stat(dir)
  if os.IsNotExist(err) {
    fmt.Printf("Directory %s not found\n", dir)
    os.Exit(1)
  }
  
  fs := http.FileServer(http.Dir(dir))

  handler := handleIndex(fs)
  http.HandleFunc("/", handler)

  fmt.Printf("Serving HTTP on http://localhost:%s from directory %s", port, dir)

  http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func handleIndex(h http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "*")

    if req.Method == "OPTIONS" {
      w.WriteHeader(http.StatusOK)
      return
    }

    h.ServeHTTP(w, req)
  }
}
