package main

import (
	config "github.com/ViPDanger/Golang/Internal/Config"
	pg "github.com/ViPDanger/Golang/Internal/Postgress"
)

func main() {
	config := config.Read_Config()
	/*


		var NewServer internal.Server
		if err := NewServer.Run(config.Adress, config.Port); err != nil {
			log.Fatal(err)
		}
	*/
	pg.PG_connect(config)

}
