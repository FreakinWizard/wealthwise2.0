package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Block struct {
	Title         string `json:"title"`
	Content1      string `json:"content1"`
	Content2      string `json:"content2"`
	ButtonContent string `json:"button_content"`
	ID            int    `json:"id"`
}

func HandleBLockData(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var block Block
	if err := json.NewDecoder(r.Body).Decode(&block); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE Blocks SET title = $1, content1 = $2, content2 = $3, button_content = $4 WHERE id = $5`

	err := db.QueryRow(query, block.Title, block.Content1, block.Content2, block.ButtonContent, block.ID)

	if err != nil {
		log.Fatal(err)
	}
}

func HandleBlockRetrieve(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var blockData struct {
		ID            int
		Title         string
		Content1      string
		Content2      string
		ButtonContent string
	}

	// Get the "id" query parameter from the URL
	queryParams := r.URL.Query()
	entityIDStr := queryParams.Get("id") // Retrieve "id" from the query string, e.g., ?id=1

	// Ensure the ID parameter is provided and valid
	if entityIDStr == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	// Convert the string ID to an integer
	entityID, err := strconv.Atoi(entityIDStr)
	if err != nil {
		http.Error(w, "Invalid 'id' query parameter", http.StatusBadRequest)
		return
	}

	// Query the database for the block data
	selectQuery := `SELECT id, title, content1, content2, button_content FROM Blocks WHERE id = $1`
	err = db.QueryRow(selectQuery, entityID).Scan(&blockData.ID, &blockData.Title, &blockData.Content1, &blockData.Content2, &blockData.ButtonContent)
	if err != nil {
		http.Error(w, "Failed to get data from the database", http.StatusInternalServerError)
		return
	}

	// Prepare the response as JSON
	response := struct {
		ID            int    `json:"id"`
		Title         string `json:"title"`
		Content1      string `json:"content1"`
		Content2      string `json:"content2"`
		ButtonContent string `json:"button_content"`
	}{
		ID:            blockData.ID,
		Title:         blockData.Title,
		Content1:      blockData.Content1,
		Content2:      blockData.Content2,
		ButtonContent: blockData.ButtonContent,
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
