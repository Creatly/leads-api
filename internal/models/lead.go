package models

import "fmt"

type Lead struct {
	FirstName string `json:"firstname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
	Source    string `json:"source" binding:"required"`
	Niche     string `json:"niche"`
	UtmSource string `json:"utm_source"`
	UtmMedium string `json:"utm_medium"`
}

func (l Lead) Info() string {
	return fmt.Sprintf("Имя: **%s**\nEmail: **%s**\nТелефон: **%s**\nФорма: **%s**\nНиша: **%s**\nUTM Source: **%s**\nUTM Medium: **%s**\n",
		l.FirstName, l.Email, l.Phone, l.Source, l.Niche, l.UtmSource, l.UtmMedium)
}

func (l Lead) CardTitle() string {
	return fmt.Sprintf("%s | ИЗ ФОРМЫ - %s", l.Email, l.Source)
}
