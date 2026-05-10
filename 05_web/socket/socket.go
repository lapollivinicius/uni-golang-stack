package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

  listen, err := net.Listen("tcp", ":3000")

  if err != nil {

    panic(err)

  }

  defer listen.Close()

  for {

    connection, err := listen.Accept()
    if err != nil {
      panic(err)
    }

    go func(connection net.Conn) {
      for {

        data, err := bufio.NewReader(connection).ReadString('\n')

        if err != nil {
          panic(err)
        }

        if strings.Contains(data, "quit") {
          break
        }

        fmt.Println(data)

      }
      
      connection.Close()
      
      fmt.Println("connection closed")
    }(connection)

  }

}