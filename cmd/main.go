package main

import (
	"context"
	"log"

	c "github.com/ViPDanger/Golang/Internal/config"
	sv "github.com/ViPDanger/Golang/Internal/handler"
	pg "github.com/ViPDanger/Golang/Internal/postgres"
)

func main() {

	config := c.Read_Config()
	client, _ := pg.NewClient(context.Background(), config)
	rep := pg.NewRepository(client)
	var NewServer sv.Server
	if err := NewServer.Run(config.Adress, config.Port, rep); err != nil {
		log.Fatal(err)
	}
}
