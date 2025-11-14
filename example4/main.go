package main

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Word     string   `json:"word"`
	Results  int      `json:"results"`
	Anagrams []string `json:"anagrams"`
}

func main() {
	client := &http.Client{}

	resp, err := client.Get("https://anagram-solver.onrender.com/master")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var data APIResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	println("Data:", data.Results)
}
