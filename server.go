package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
    "net/http"
    "log"
    "./utils"
)

type resp_struct struct {
    Text string
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func AcceptVote(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    decoder := json.NewDecoder(r.Body)
    var resp resp_struct
    err := decoder.Decode(&resp)
    utils.CheckError(err, "connection error")
    defer r.Body.Close()
    log.Println(resp)
    fmt.Fprintf(w, resp.Text)
}

func JoinGame(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func EndGame(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func Submit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func DisplayAnswers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func ShowAnswers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.POST("/vote", AcceptVote)
    router.POST("/join", JoinGame)
    router.POST("/end", EndGame)
    router.POST("/submit", Submit)
    router.GET("/show_answers", ShowAnswers)
    router.POST("/create_user", CreateUser)

    log.Fatal(http.ListenAndServe(":8080", router))
}
