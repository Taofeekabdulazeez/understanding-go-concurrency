package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type APIResponse struct {
	Word     string   `json:"word"`
	Results  int      `json:"results"`
	Anagrams []string `json:"anagrams"`
}

func main() {
	ctx := context.Background()
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	api_url := "https://anagram-solver.onrender.com/master"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, api_url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var data APIResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	println("Word:", data.Word)
	println("Results:", data.Results)
	println("Anagrams:")

	for _, anagram := range data.Anagrams {
		println("-", anagram)
	}

}
