package main

import (
	"fmt"
	"statgen/src/server"
  "flag"
)

func main() {
  dirPtr := flag.String("dir", ".", "Directory to serve files from")
  portPtr := flag.String("port", "8000", "Port to serve HTTP on")
  flag.Parse()


  fmt.Println("Starting up server...")

  server.Start(*dirPtr, *portPtr)
}
