package main

import (
	"github.com/adlio/trello"
	"github.com/creatly/leads-api/internal/crm"
	"github.com/creatly/leads-api/internal/server"
	"log"
	"os"
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

	log.Println("Starting server on port 8000")

	srv := server.New("8000", trelloCRM)
	if err := srv.Init(); err != nil {
		log.Fatal(err)
	}
}
