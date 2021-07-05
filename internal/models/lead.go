package models

import (
	"bytes"
	"fmt"
	"text/template"
)

const templateInfo = `Имя: {{MakeItBold .FirstName}}
Email: {{MakeItBold .Email}}
Телефон: {{MakeItBold .Phone}}
Форма: {{MakeItBold .Source}}
Ниша: {{MakeItBold .Niche}}
UTM Source: {{MakeItBold .UtmSource}}
UTM Medium: {{MakeItBold .UtmMedium}}
`

var objTemplateInfo = template.Must(
	template.
		New("Info").
		Funcs(template.FuncMap{
			"MakeItBold": func(s string) string { return "**" + s + "**" },
		}).
		Parse(templateInfo),
)

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
	return l.String()
}

func (l Lead) CardTitle() string {
	return fmt.Sprintf("%s | ИЗ ФОРМЫ - %s", l.Email, l.Source)
}

func (l Lead) String() string {
	var buf bytes.Buffer

	if err := objTemplateInfo.Execute(&buf, l); err != nil {
		return err.Error()
	}

	return buf.String()
}
