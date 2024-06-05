package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"url-shortener/pkg/models"
	"url-shortener/pkg/utils"
)

type ShortenRequest struct {
	OriginalURL string `json:"original_url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shortCode, err := utils.GenerateShortURL()
	if err != nil {
		http.Error(w, "Failed to generate short URL", http.StatusInternalServerError)
		return
	}

	url := models.URL{OriginalURL: req.OriginalURL, ShortURL: shortCode}
	models.DB.Create(&url)

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	response := ShortenResponse{ShortURL: baseURL + "/" + shortCode}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
