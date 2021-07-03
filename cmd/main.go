package main

import (
	"context"
	"github.com/adlio/trello"
	"github.com/creatly/leads-api/internal/config"
	"github.com/creatly/leads-api/internal/crm"
	"github.com/creatly/leads-api/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatalf("Received an error from config package: %v", err)
	}

	trelloClient := trello.NewClient(cfg.Trello.ApiKey, cfg.Trello.ApiToken)
	trelloCRM, err := crm.NewTrelloClient(trelloClient, cfg.Trello.BoardId, cfg.Trello.ListName)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(8000, trelloCRM)
	go func() {
		if err := srv.Init(); err != nil {
			log.Printf("server error: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
