package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"wealthwise2.0/db"
	"wealthwise2.0/defaults"
	"wealthwise2.0/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	cfg := &db.DBConfig{
		Address:  os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	db, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	defer db.Close()

	defaults.CreateTable(db)
	defaults.CreateCardTable(db)
	defaults.Defaults(db)
	//blocks
	http.HandleFunc("/block", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Handle retrieving block data
			services.HandleBlockRetrieve(w, r, db)
		} else if r.Method == http.MethodPost {
			// Handle updating block data
			services.HandleBLockData(w, r, db)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/card", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			services.HandleCardRetrieve(w, r, db)
		} else if r.Method == http.MethodPost {
			services.HandleCardData(w, r, db)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
