package main

import (
	"encoding/json"
	"fmt"
	"github.com/polina/grammar/pkg/suggest"
	"github.com/polina/grammar/pkg/suggest/radixTree"
	"io"
	"log"
	"net/http"
)

const portNum string = ":8080"

var radix *radixTree.Node

type SuggestRequest struct {
	Value string `json:"value"`
}

// {
//  "queries": [{}],
//
//}

type SuggestQuery struct {
	Text string `json:"text"`
}

type SuggestResponse struct {
	Queries []SuggestQuery `json:"queries"`
}

func main() {
	radix = suggest.BuildRadixTree()
	fmt.Println("radix built")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/suggest", Suggest)
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/html")
	http.ServeFile(w, r, "/home/polina/src/grammar/static/site")
}

func Suggest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Not POST Method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Cannot read request body", http.StatusBadRequest)
	}

	res := SuggestRequest{}

	err = json.Unmarshal(body, &res)
	if err != nil {
		http.Error(w, "bad JSON", http.StatusBadRequest)
		return
	}

	prefix := res.Value
	fmt.Println(prefix)

	payload := radix.Search(prefix)
	resp := SuggestResponse{
		Queries: make([]SuggestQuery, 0, 8),
	}
	for _, query := range payload {
		fmt.Println(query.Name())
		query := SuggestQuery{
			Text: query.Name(),
		}
		resp.Queries = append(resp.Queries, query)
	}
	payloadJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "bad JSON", http.StatusBadRequest)
		return
	}
	w.Write(payloadJSON)

}
