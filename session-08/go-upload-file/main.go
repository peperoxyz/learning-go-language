package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/upload", uploadHandler)
	fmt.Println("Server started at :1010")
	http.ListenAndServe(":1010", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Could not get file from form data", http.StatusBadRequest)
		}
		defer file.Close()

		// folder untuk taro file
		uploadDir := "./uploads"
		os.MkdirAll(uploadDir, os.ModePerm)
		
		// define file path
		filePath := filepath.Join(uploadDir, handler.Filename)

		// create dulu, blm proses upload
		newFile, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Could not create file", http.StatusInternalServerError)
			return
		} 
		defer newFile.Close()

		// store ke dalam folder dengan i/o
		_, err = io.Copy(newFile, file)
		if err != nil {
			http.Error(w, "Could not copy file data", http.StatusInternalServerError)
			return
		}

		responseWithJSON(w, http.StatusOK, map[string]string{"message": "File uploaded successfully"})
	} else {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// helper function to send an error message as JSON
func respondWithError(w http.ResponseWriter, message string, statusCode int) {
	responseWithJSON(w, statusCode, map[string]string{"error":message})
}

func responseWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the response data}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

