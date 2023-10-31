package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Obj map[string]any

func main() {

	API := "https://cyberleninka.ru/api/search"

	post_body, err := json.Marshal(Obj{
		"mode": "articles",
		"size": 10,
		"q":    "golang",
		"from": 0,
	})
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(API, "application/json", bytes.NewReader(post_body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response status:", resp.Status)
}
