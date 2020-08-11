package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stopSignal := make(chan os.Signal)
	signal.Notify(stopSignal, os.Interrupt)

	Server := ServerInit()
	go func() {
		log.Println("Server listening at ", Server.Addr)
		if err := Server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	select {
	case <-stopSignal:
		log.Println("shut down the server")
	}
	Server.Shutdown(ctx)
	os.Exit(0)
}
