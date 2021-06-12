package crm

import (
	"errors"
	"fmt"
	"github.com/adlio/trello"
	"github.com/creatly/leads-api/internal/models"
)

type TrelloClient struct {
	client *trello.Client
	listID string
}

func NewTrelloClient(client *trello.Client, boardID, targetListName string) (*TrelloClient, error) {
	service := &TrelloClient{client: client}
	err := service.setListID(boardID, targetListName)

	return service, err
}

func (s *TrelloClient) SaveLead(lead models.Lead) error {
	list, err := s.client.GetList(s.listID, trello.Defaults())
	if err != nil {
		return err
	}

	card := &trello.Card{
		Name: lead.CardTitle(),
		Desc: lead.Info(),
	}

	return list.AddCard(card, trello.Defaults())
}

func (s *TrelloClient) setListID(boardID, targetListName string) error {
	board, err := s.client.GetBoard(boardID, trello.Defaults())
	if err != nil {
		return err
	}

	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return err
	}

	for _, list := range lists {
		if list.Name == targetListName {
			s.listID = list.ID
			return nil
		}
	}

	errorMessage := fmt.Sprintf("target list '%s' not found", targetListName)
	return errors.New(errorMessage)
}
