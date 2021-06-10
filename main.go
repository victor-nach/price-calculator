package main

import (
	"fmt"
	"github.com/victor-nach/price-calculator/config"
	"github.com/victor-nach/price-calculator/db/mongo"
	"github.com/victor-nach/price-calculator/lib/coindesk"
	"github.com/victor-nach/price-calculator/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.LoadSecrets()

	coindeskClient := coindesk.NewClient(cfg.CoindeskURL)

	mongoStore, _, err := mongo.New(cfg.DBURL, cfg.DBName)
	if err != nil {
		log.Fatalf("failed to open mongodb: %v", err)
	}
	srv := server.NewServer(coindeskClient, mongoStore)

	// create channel to listen to shutdown signals
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		addr := fmt.Sprintf(":%s", cfg.Port)
		err := srv.Run(addr)
		if err != nil {
			log.Fatal(fmt.Sprintf("failed to start service: %v", err))
		}
	}()

	<-shutdownChan
	log.Println("Closing application")
	// do cleanups before exit
}
