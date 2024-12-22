package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"visit-counter/pkg/db/models"

	"gorm.io/gorm"
)

func Create(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		token := query.Get("token")
		if token != os.Getenv("TOKEN") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		url := query.Get("url")

		if url == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		page := models.Page{
			Url: url,
		}

		result := db.Create(&page)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		db.First(&page, page.ID)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(page)
	}
}