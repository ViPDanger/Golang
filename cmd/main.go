package main

import (
	internal "github.com/ViPDanger/Golang/Internal"
)

func main() {
	/*
		config := internal.Read_Config()

			var NewServer internal.Server
			if err := NewServer.Run(config.Adress, config.Port); err != nil {
				log.Fatal(err)
			}
	*/
	internal.PG_connect()

}
