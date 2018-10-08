package main

import (
	"net"
  "fmt"
  "bufio"
  "log"
  // "strings"
)

func startserver(){
  fmt.Println("Launching server...")

  // listen on all interfaces
  listener, err := net.Listen("tcp", "127.0.0.1:8081")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Listening on port [8081]\n")

  for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
         reader, _ := bufio.NewReader(conn).ReadString('\n')
         fmt.Println(reader)
    }
}

func main() {
  go startserver()
}
