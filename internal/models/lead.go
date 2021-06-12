package models

import "fmt"

type Lead struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
	Source    string `json:"source" binding:"required"`
}

func (l Lead) Info() string {
	return fmt.Sprintf("Имя: %s\nФамилия: %s\nEmail: %s\nТелефон: %s\nИсточник: %s\n",
		l.FirstName, l.LastName, l.Email, l.Phone, l.Source)
}

func (l Lead) CardTitle() string {
	return fmt.Sprintf("%s | ИСТОЧНИК - %s", l.Email, l.Source)
}
