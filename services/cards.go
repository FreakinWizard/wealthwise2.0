package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Card struct {
	Title         string `json:"title"`
	Content1      string `json:"content1"`
	ButtonContent string `json:"button_content"`
	ID            int    `json:"id"`
}

func HandleCardData(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var card Card
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE Cards SET title = $1, content1 = $2, button_content = $3 WHERE id = $4`

	err := db.QueryRow(query, card.Title, card.Content1, card.ButtonContent, card.ID)

	if err != nil {
		log.Fatal(err)
	}

}

func HandleCardRetrieve(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var cardData struct {
		ID      int
		Title   string
		Content string
	}

	queryParams := r.URL.Query()
	entityIDStr := queryParams.Get("id")

	if entityIDStr == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	entityID, err := strconv.Atoi(entityIDStr)
	if err != nil {
		http.Error(w, "Invalid 'id' query parameter", http.StatusBadRequest)
		return
	}

	selectQuery := `SELECT id, title, content FROM Cards WHERE id = $1`
	err = db.QueryRow(selectQuery, entityID).Scan(&cardData.ID, &cardData.Title, &cardData.Content)
	if err != nil {
		http.Error(w, "Failed to get data from the database", http.StatusInternalServerError)
		return
	}

	response := struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}{
		ID:      cardData.ID,
		Title:   cardData.Title,
		Content: cardData.Content,
	}

	// Set response header to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the response struct to JSON and send it
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
