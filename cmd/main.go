package main

import (
	"context"
	"log"

	sv "github.com/ViPDanger/Golang/Internal/Handler"
	c "github.com/ViPDanger/Golang/Internal/config"
	pg "github.com/ViPDanger/Golang/Internal/postgres"
)

func main() {

	config := c.Read_Config()
	client, _ := pg.NewClient(context.Background(), config)
	rep := pg.NewRepository(client)
	log.Println(rep)
	var NewServer sv.Server
	if err := NewServer.Run(config.Adress, config.Port, rep); err != nil {
		log.Fatal(err)
	}
}
