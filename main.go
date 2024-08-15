package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const (
	translatorAPIURL = "https://google-translator9.p.rapidapi.com/v2"
	apiHostHeader    = "google-translator9.p.rapidapi.com"
	apiKeyHeader     = "06de4ebb25mshad1d579533fcb83p1d7bebjsne0c7e73548b9" // Replace with your actual API key
)

type TranslationRequest struct {
	Q      string `json:"q"`
	Source string `json:"source"`
	Target string `json:"target"`
	Format string `json:"format"`
}

type TranslationResponse struct {
	Data struct {
		Translations []struct {
			TranslatedText string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

func translateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req TranslationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Prepare the request to the translation API
	apiRequestBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to prepare request", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	reqAPI, err := http.NewRequest(http.MethodPost, translatorAPIURL, bytes.NewBuffer(apiRequestBody))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	reqAPI.Header.Add("Content-Type", "application/json")
	reqAPI.Header.Add("x-rapidapi-host", apiHostHeader)
	reqAPI.Header.Add("x-rapidapi-key", apiKeyHeader)

	respAPI, err := client.Do(reqAPI)
	if err != nil {
		http.Error(w, "Failed to call translation API", http.StatusInternalServerError)
		return
	}
	defer respAPI.Body.Close()

	if respAPI.StatusCode != http.StatusOK {
		http.Error(w, "Translation API error", http.StatusInternalServerError)
		return
	}

	var translationResp TranslationResponse
	if err := json.NewDecoder(respAPI.Body).Decode(&translationResp); err != nil {
		http.Error(w, "Failed to parse translation response", http.StatusInternalServerError)
		return
	}

	// Send the response back to the client
	response := map[string]interface{}{
		"data": translationResp.Data,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func main() {
	http.HandleFunc("/translate", translateHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}