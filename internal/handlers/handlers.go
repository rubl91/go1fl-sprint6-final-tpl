package handlers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleMain(w http.ResponseWriter, r *http.Request) {
	page, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	w.Write(page)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "неудалось прочитать файл", http.StatusBadRequest)
		return
	}
	convertResult := service.MorseConverter(data)

	if err := os.MkdirAll("uploads", 0755); err != nil {
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	root, err := os.OpenRoot("uploads")
	if err != nil {
		http.Error(w, "внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	defer root.Close()

	dst, err := root.Create(header.Filename)
	if err != nil {
		http.Error(w, "ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = dst.Write([]byte(convertResult))
	if err != nil {
		http.Error(w, "ошибка при записи файла", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(convertResult))
}
