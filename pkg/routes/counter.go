package routes

import (
	"net/http"
	"visit-counter/pkg/db/models"

	"gorm.io/gorm"
)

func GetCount(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		url := query.Get("url")

		if url == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		var page models.Page
		result := db.Where(&models.Page{Url: url}).First(&page)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				http.Error(w, "Page not found", http.StatusNotFound)
			} else {
				http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write([]byte(page.CreateSVG()))
		go func() {
			db.Model(&page).Update("visits", page.Visits+1)
		}()
	}
}
