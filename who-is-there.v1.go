package main

import (
  "fmt"
  "net"
  "net/http"
  "os"
  "time"
)


func main() {
  name, err := os.Hostname()
  if err != nil {
    panic(err)
  }

  now := time.Now()
  fmt.Println(now.Format("2006-01-02 15:04:05"), "- Server (version 1) started and listening on port 8080.")

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    now := time.Now()
    ip, _, _ := net.SplitHostPort(r.RemoteAddr)
    fmt.Println(now.Format("2006-01-02 15:04:05"), "-", name, "- Received a request from:", ip)
    fmt.Fprintf(w, "Hello, this is host: %s\n", name)
    })

  http.ListenAndServe(":8080", nil)
}
