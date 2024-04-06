package main

import (
	"context"
	"log"
	"time"

	c "github.com/ViPDanger/Golang/Internal/config"
	pg "github.com/ViPDanger/Golang/Internal/postgres"
	"github.com/ViPDanger/Golang/Internal/structures"
)

func main() {

	config := c.Read_Config()
	client, _ := pg.NewClient(context.Background(), config)
	rep := pg.NewRepository(client)

	author := structures.ContentDTO{
		Content:   "Content from Golang!",
		Author_id: 54,
		Date:      time.Now().Format("2006.02.03 15:04:05"),
	}
	log.Println(author)
	authors, _ := rep.All_Authors()
	log.Println(authors)
	/*


		var NewServer internal.Server
		if err := NewServer.Run(config.Adress, config.Port); err != nil {
			log.Fatal(err)
		}
	*/
}
