package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"back-end/internal/repository"
)

func HandleGetMenu(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		menus, err := repository.FetchMenus(db)
		if err != nil {
			http.Error(w, "Gagal mengambil data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(menus)
	}
}