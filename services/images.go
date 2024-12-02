package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const ImagePathBase string = "assets/images"

func UploadImage(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query().Get("name")

	if err := r.ParseMultipartForm(300 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rawData, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileData, err := io.ReadAll(rawData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filePath := filepath.Join(ImagePathBase, imageName)

	// Check if a file with the same name exists and delete it
	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil {
			http.Error(w, fmt.Sprintf("Failed to delete existing image: %v", err), http.StatusInternalServerError)
			return
		}
	}

	// Write the new image to the file system
	if err := os.WriteFile(filePath, fileData, 0644); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query().Get("name")

	if len(imageName) == 0 {
		http.Error(w, fmt.Errorf("no image 'name' is not given as a query param").Error(), http.StatusBadRequest)
		return
	}

	imagePath := filepath.Join(ImagePathBase, imageName)

	http.ServeFile(w, r, imagePath)
}
