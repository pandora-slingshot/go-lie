Start the server like this: `go run server.go localhost`
Start the client like this: `go run client.go tom localhost`
using this for inspiration: https://github.com/jhudson8/golang-chat-example

Methods
===
AcceptVote
=====
 `curl -H "Content-Type: application/json" -X POST -d '{"Text": "Some answer text submitted by user"}' http://localhost:8080/vote/`
