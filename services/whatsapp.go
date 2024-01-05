package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func newRequest(requestMethod, path string, body map[string]interface{}) (*http.Request, error) {
	bodyJson, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	return http.NewRequest(requestMethod, path, bytes.NewReader(bodyJson))
}

func SendNewMessage(content string) error {
	path := os.Getenv("WHATSAPP_SERVICE_URL") + os.Getenv("WHATSAPP_ID") + "/messages"

	bodyRequest := map[string]interface{}{
		"messaging_product": "whatsapp",
		"recipient_type":    "individual",
		"to":                os.Getenv("PHONE_NUMBER"),
		"type":              "text",
		"text": map[string]interface{}{
			"preview_url": false,
			"body":        content,
		},
	}

	req, err := newRequest("POST", path, bodyRequest)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("WHATSAPP_TOKEN")))
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}

	response, err := client.Do(req)

	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}
