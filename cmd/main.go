package main

import (
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
	//var Buffer bytes.Buffer
	//Buffer.Write([]byte("SUSUAT"))
	var NewServer internal.Server
	var handler http.Handler
	addres, port := internal.Read_Config()
	go NewServer.Run(addres, port, handler)
	//response, err := http.Post(addres+":"+port+"/get", "Data", )
	//err_log(err)

	//log.Println(string(content))
}
