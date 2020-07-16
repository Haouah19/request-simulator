package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"fmt"
)

type RequestBody struct {
	Typing         string
	ConversationID int
	Timestamp      int
}

type Hint struct {
	Id int `json:"id"`
	Text string `json:"text"`
	Score float32 `json:"score"`
	Rep string `json:"rep"`
}

type Response struct {
    ConversationID int
    Hints []Hint
    Timestamp int
}

func main() {

	req := RequestBody{
		Typing : os.Args[1],
		ConversationID : 0,
		Timestamp : 0,
	}
	encodedJson,_ := json.Marshal(req)

	resp, err := http.Post("http://0.0.0.0:5600/post-hints", "application/json", bytes.NewBuffer(encodedJson))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var Req_rep Response
	err = json.Unmarshal(body, &Req_rep)
	if err != nil {
		panic(err)
	}

	fmt.Println(Req_rep)
}

