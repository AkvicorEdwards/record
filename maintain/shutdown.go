package maintenance

import (
	"log"
	"os"
	"os/signal"
	"record/dam"
	"time"
)

func ShutDownListener() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("ShutDownListener recover", err)
		}
	}()
	down := make(chan os.Signal, 1)
	signal.Notify(down, os.Interrupt, os.Kill)
	<-down
	go func() {
		select {
		case <-time.After(30*time.Second):
			os.Exit(0)
		}
	}()

	log.Println("Lock Database")
	dam.LockAll()

	os.Exit(0)
}
