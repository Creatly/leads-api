package models

import "fmt"

type Source string

const (
	PDF      Source = "pdf"
	Beta     Source = "beta"
	DemoCall Source = "demo call"
)

type Lead struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Source    Source `json:"source"`
}

func (l Lead) Info() string {
	return fmt.Sprintf("Имя: %s\nФамилия: %s\nEmail: %s\nТелефон: %s\nИсточник: %s\n",
		l.FirstName, l.LastName, l.Email, l.Phone, l.Source)
}
