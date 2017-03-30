package main

import (
    "net"
    "./utils"
)

func main() {
    socket, err := net.Listen("tcp", ":8080")
    utils.CheckError(err, "connection error")
    for true {
        conn, err := socket.Accept()
        utils.CheckError(err, "connection error")
        println("connected: ", conn)
    }
}
