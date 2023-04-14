package main

import (
	"io/ioutil"
	"log"
	"net/http"

	internal "github.com/ViPDanger/Golang/internal"
)

func err_log(err error) bool {
	if err != nil {

		log.Println("Error: ", err, " - ", err.Error())
		return true
	}
	return false
}

func main() {
	addres, port := internal.Read_Config()
	response, err := http.Get(addres + ":" + port + "/get")
	err_log(err)
	log.Println("Status: ", response.StatusCode)
	content, _ := ioutil.ReadAll(response.Body)
	log.Println(string(content))
	return
}
