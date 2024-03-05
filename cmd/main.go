package main

import (
	"log"

	internal "github.com/ViPDanger/Golang/Internal"
)

func err_log(err error) bool {
	if err != nil {

		log.Println("Error: ", err, " - ", err.Error())
		return true
	}
	return false
}

func main() {

	config := internal.Read_Config()
	var NewServer internal.Server

	if err := NewServer.Run(config.Adress, config.Port); err != nil {
		log.Fatal(err)
	}

}
