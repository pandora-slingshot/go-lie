package main

import (
    "net"
    "os"
)

func main() {
    socket, err := net.Listen("tcp", ":8080")
    if err != nil {
        println("danger" + ": ", err.Error())
        os.Exit(1)
    }
    for true {
        conn, err := socket.Accept()
        if err != nil {
            println("danger" + ": ", err.Error())
            os.Exit(1)
        }
        println("connected: ", conn)
    }
}
