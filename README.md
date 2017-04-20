Setup
===
Install dependencies
---
To set your go path: `export GOPATH=<the working directory you have the code in>`
To install depencies: `go get`
If that doesnt work : 'mkdir bin
                       export GOBIN=$GOPATH/bin`
Dependencies should show up in the `src` folder.

Usage
===

Starting the Client
---
`go run client.go <username> <host>`
Host will probably be localhost for now

Starting the Server
---
`go run server.go <host>`
Host will probably be localhost for now

API Endpoints
===
Method: AcceptVote
---
Example usage:
 `curl -H "Content-Type: application/json" -X POST -d '{"Text": "Some answer text submitted by user"}' http://localhost:8080/vote/`

using this for inspiration: https://github.com/jhudson8/golang-chat-example
