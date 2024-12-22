package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	Url string `gorm:"unique;uniqueIndex"`
	Visits uint `gorm:"default:0"`
}

func (p *Page) CreateSVG() []byte {
	return []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100">
	<text x="10" y="20" font-family="Verdana" font-size="10">` + fmt.Sprint(p.Visits) + `</text>
</svg>`)
}