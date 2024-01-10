package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Article struct {
	Source struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"source"`
	Author      string `json:"autor"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type NewsPayload struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func GetNews() (articles []Article, err error) {
	params := url.Values{}
	params.Add("country", "br")
	params.Add("category", "technology")
	params.Add("apiKey", os.Getenv("NEWS_API_KEY"))

	path := os.Getenv("NEWS_API_URL") + "?" + params.Encode()

	fmt.Println(path)

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		log.Fatalf("Erro ao criar request: %s", err.Error())
		return nil, err
	}

	client := http.Client{}

	response, err := client.Do(req)

	if err != nil {
		log.Fatalf("Erro ao realizar requisição: %s", err.Error())
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Falha ao realizar requisição")
		return nil, errors.New("Error to realize request")
	}

	var bodyParsed NewsPayload

	err = json.NewDecoder(response.Body).Decode(&bodyParsed)

	if err != nil {
		log.Fatalf("Erro ao formatar json: %s", err.Error())
	}

	return bodyParsed.Articles[0:10], nil
}
