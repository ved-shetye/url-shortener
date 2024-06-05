package handlers

import (
	"net/http"
	"url-shortener/pkg/models"

	"github.com/gorilla/mux"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	var url models.URL
	if models.DB.Where("short_url = ?", shortURL).First(&url).RecordNotFound() {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusMovedPermanently)
}
