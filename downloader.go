package main

import (
  "net/http"
  "io"
  "os"
  "fmt"
)

func main() {
  if (len(os.Args[1:]) != 1) {
    fmt.Printf("Usage of %s:\n", os.Args[0])
    fmt.Printf("    downloader https://www.example.com > test.html\n")
    os.Exit(1)
  }

  resp, _ := http.Get(os.Args[1])
  defer resp.Body.Close()

  _, _ = io.Copy(os.Stdout, resp.Body)
}
