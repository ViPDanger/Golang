package main

import (
	"context"

	c "github.com/ViPDanger/Golang/Internal/config"
	pg "github.com/ViPDanger/Golang/Internal/postgres"
)

func main() {

	config := c.Read_Config()
	pg.NewClient(context.Background(),config)
	/*


		var NewServer internal.Server
		if err := NewServer.Run(config.Adress, config.Port); err != nil {
			log.Fatal(err)
		}
	*/
}
