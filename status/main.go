package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("starting")
	defer log.Println("stopped")
	ctx, cancel := context.WithCancel(context.Background())

	rand.Seed(time.Now().Unix())

	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		<-interrupt
		cancel()
	}()
	<-ctx.Done()
}
