package main

import (
	"context"
	"github.com/adlio/trello"
	"github.com/creatly/leads-api/internal/crm"
	"github.com/creatly/leads-api/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	apiKey, apiToken := os.Getenv("TRELLO_API_KEY"), os.Getenv("TRELLO_API_TOKEN")
	if apiKey == "" || apiToken == "" {
		log.Fatal("empty trello credentials")
	}

	boardID, targetListName := os.Getenv("TRELLO_BOARD_ID"), os.Getenv("TRELLO_LIST_NAME")
	trelloClient := trello.NewClient(apiKey, apiToken)
	trelloCRM, err := crm.NewTrelloClient(trelloClient, boardID, targetListName)
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
