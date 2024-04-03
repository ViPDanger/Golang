package main

import (
	"github.com/ViPDanger/Golang/Internal/config"
	pg "github.com/ViPDanger/Golang/Internal/postgres"
)

func main() {

	config := config.Read_Config()
	pg.NewClient(config)
	/*


		var NewServer internal.Server
		if err := NewServer.Run(config.Adress, config.Port); err != nil {
			log.Fatal(err)
		}
	*/
}
