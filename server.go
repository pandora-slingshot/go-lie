package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
	"encoding/json"
    "net/http"
    "log"
    "./utils"
	"math/rand"
	"time"
)

type room struct {
	Key string
}

rooms := []room{}

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
func StartGame(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	hash := RandomString(5)
	var game_room := room{hash: hash}
	rooms = append(rooms, game_room)
	fmt.Fprintf(w, hash)
}
func RandomString(strlen int) string {
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}
func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.POST("/vote/", AcceptVote)
    router.POST("/start_game/", StartGame)
    log.Fatal(http.ListenAndServe(":8080", router))
}
