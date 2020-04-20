package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Word from POST request
type Word struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// WordSha To be Json
type WordSha struct {
	Created string `json:"created"`
	Sha     string `json:"sha"`
}

// DbShaResult is dummy database record
type DbShaResult struct {
	ID      string `json:"id"`
	Created string `json:"created"`
	Sha     string `json:"sha"`
}

var dbDummy = []DbShaResult{
	{
		ID:      "1",
		Created: time.Now().String(),
		Sha:     "sh4r3sult-1",
	},
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/sha", shaHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func shaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case "GET":
		jsonData, err := json.Marshal(&dbDummy)
		if err != nil {
			log.Println(err)
		}
		// send status to front end
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		w.Write((jsonData))
		// w.WriteHeader(http.StatusOK)
		// w.Write([]byte(`{"message": "hello GET"}`))
	case "POST":
		// set Header for CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// UnMarshal Received JSON
		var wo Word
		err := json.NewDecoder(r.Body).Decode(&wo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// fmt.Printf("%+v", wo)

		// Pass to func to calculate the SHA Sum
		r := shaSum(&wo)

		//storeSha
		var wsha WordSha
		storeSha(r, &wsha)
		fmt.Printf("%+v\n", wsha)
		fmt.Println("=========================================================================")
		jsonData, err := json.Marshal(&wsha)
		if err != nil {
			log.Println(err)
		}

		// send status to front end
		w.WriteHeader(http.StatusOK)
		w.Write((jsonData))
		// w.Write([]byte(`{"status": "OK"}`))

	}
}

func shaSum(word *Word) string {
	// Calculate the sha of the word
	sum := sha256.Sum256([]byte(word.Text))
	// convert to []byte
	s := sum[:]

	// build the string from []byte
	var b strings.Builder
	b.Write(s)

	r := b.String()
	rf := fmt.Sprintf("%x", r)
	return rf
}

func storeSha(sha string, ws *WordSha) {
	// convert to struct
	ws.Created = time.Now().String()
	ws.Sha = sha
	// fmt.Printf("%+v\n", ws)

	// Store to DB
}
