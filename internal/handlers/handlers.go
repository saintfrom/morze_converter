package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"morze_converter/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "failed to read index.html", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "failed to parse form", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "failed to get file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "failed to read file", http.StatusInternalServerError)
		return
	}

	converted, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "failed to convert", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	fileName := time.Now().UTC().String() + ext

	if err := os.WriteFile(fileName, []byte(converted), 0o644); err != nil {
		http.Error(w, "failed to write result file", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, converted)
}
