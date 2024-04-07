package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	config "github.com/ViPDanger/Golang/Internal/config"
	txt "github.com/ViPDanger/Golang/Internal/txt_file"
)

func TXT_AddHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	data := r.URL.Path
	data = data[strings.IndexRune(data[1:], '/')+2:]
	log.Println("GetData: ", data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New string has been added: " + data + "\n"))
	txt.TXT_Add_Data(data)

	// Показ нового списка
	TXT_ShowHandler(w, r)
}

func TXT_ShowHandler(w http.ResponseWriter, r *http.Request) {
	config := config.Read_Config()
	data := *txt.TXT_Read_Data(config.Data_File)
	log.Println("Data was readed")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Current List:\n"))
	for i := 0; i < len(data); i++ {
		w.Write([]byte(strconv.Itoa(i) + ": " + data[i] + "\n"))
	}
}

func TXT_DeleteHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	data := r.URL.Path
	data = data[strings.IndexRune(data[1:], '/')+2:]

	// Удаление
	data_int, _ := strconv.Atoi(data)
	txt.TXT_Delete_Data(data_int)
	log.Println("String deleted: ", data_int)
	w.WriteHeader(http.StatusOK)

	// Показ нового списка
	TXT_ShowHandler(w, r)
}
