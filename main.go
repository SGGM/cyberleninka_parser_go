package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Obj map[string]any

var API = "https://cyberleninka.ru/api/search"

func main() {

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

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Api returned a non-200 status code:", resp.Status)
		return
	}
	fmt.Println("response status:", resp.Status)

	var responseValidJSON map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&responseValidJSON); err != nil {
		fmt.Println("Error decoding JSON response", err)
		return
	}

	foundValue, found := responseValidJSON["found"]
	if found {
		fmt.Println("Number of papers found:", foundValue)
	} else {
		fmt.Println("Key 'found' not found in the response")
	}

	//Second request

	post_body2, err := json.Marshal(Obj{
		"mode": "articles",
		"size": foundValue,
		"q":    "golang",
		"from": 0,
	})
	if err != nil {
		panic(err)
	}

	resp2, err := http.Post(API, "application/json", bytes.NewReader(post_body2))
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	if resp2.StatusCode != http.StatusOK {
		fmt.Println("Api returned a non-200 status code:", resp2.Status)
		return
	}
	fmt.Println("response status:", resp2.Status)

	var responseValidJSON2 map[string]interface{}
	decoder2 := json.NewDecoder(resp2.Body)
	if err := decoder2.Decode(&responseValidJSON2); err != nil {
		fmt.Println("Error decoding JSON response", err)
		return
	}

	articles := []map[string]interface{}{}
	numberOfPapers := foundValue
	var numberOfPapers any

	fmt.Println(foundValue)

	for articleNum := 0; articleNum < number; articleNum++ {
		currArticle := map[string]interface{}{}

		currArticle["name"] = responseValidJSON["articles"].([]interface{})[articleNum].(map[string]interface{})["name"]
		currArticle["annotation"] = responseValidJSON["articles"].([]interface{})[articleNum].(map[string]interface{})["annotation"]
		currArticle["authors"] = responseValidJSON["articles"].([]interface{})[articleNum].(map[string]interface{})["authors"]
		currArticle["year"] = responseValidJSON["articles"].([]interface{})[articleNum].(map[string]interface{})["year"]

		articles = append(articles, currArticle)
	}

	fmt.Println(articles)
}
