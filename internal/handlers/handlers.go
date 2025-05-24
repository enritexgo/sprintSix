package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "mistake - incorrect method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "mistake in ParseMultipartForm", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "mistake in FormFile", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileMsg, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "mistake in ReadAll", http.StatusInternalServerError)
		return
	}

	convertedMsg, err := service.Conv(string(fileMsg))
	msg := []byte(convertedMsg)
	if err != nil {
		http.Error(w, "mistake in Conversion", http.StatusInternalServerError)
		return
	}

	fileExt := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("%s%s", time.Now().UTC().String(), fileExt)

	if err := os.WriteFile(fileName, msg, 0755); err != nil {
		http.Error(w, "mistake in WriteFile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(msg); err != nil {
		http.Error(w, "mistake in Write", http.StatusInternalServerError)
		return
	}
}
