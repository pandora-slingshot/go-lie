package main

import (
    "fmt"
    "net"
    "os"
    "./utils"
)

func main() {
    username, hostname := readArgs();
    fmt.Printf("%v", username)
    fmt.Printf("%v", hostname)
    conn, err := net.Dial("tcp", hostname + ":8080")
    utils.CheckError(err, "connection error")

    defer conn.Close()

    for true {
        println("waiting")
    }
}

func readArgs() (string, string) {
    username := os.Args[1]
    hostname:= os.Args[2]
    return username, hostname
}
