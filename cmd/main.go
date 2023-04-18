package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer stop()

	if err := NewServer.Run(config.Adress, config.Port, ctx); err != nil {
		log.Fatal(err)
	}

}
