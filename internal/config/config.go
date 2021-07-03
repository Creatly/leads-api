package config

import "github.com/kelseyhightower/envconfig"

type Trello struct {
	ListName string `defaults:"Leads" split_words:"true"`
	BoardId  string `required:"true" split_words:"true"`
	ApiToken string `required:"true" split_words:"true"`
	ApiKey   string `required:"true" split_words:"true"`
}

type Config struct {
	Trello Trello
}

func New() (Config, error) {
	cfg := Config{}
	trello := Trello{}

	if err := envconfig.Process("TRELLO", &trello); err != nil {
		return cfg, err
	}

	cfg.Trello = trello

	return cfg, nil
}
