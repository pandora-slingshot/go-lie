package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    username, hostname := readArgs();
    fmt.Printf("%v", username)
    fmt.Printf("%v", hostname)
    conn, err := net.Dial("tcp", hostname + ":8080")
    if err != nil {
        println("danger" + ": ", err.Error())
        os.Exit(1)
    }
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
