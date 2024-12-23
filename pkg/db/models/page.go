package models

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"visit-counter/pkg/utils"

	"visit-counter/assets"

	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	Url    string `gorm:"unique;uniqueIndex"`
	Visits uint   `gorm:"default:0"`
}

var imagesBase64 []string = func() []string {
	data := make([]string, 10)

	for i := 0; i <= 9; i++ {
		filePath := fmt.Sprintf("%d.gif", i)
		fileData, err := assets.ImagesFS.ReadFile(filePath)
		if err != nil {
			continue
		}
		data[i] = base64.StdEncoding.EncodeToString(fileData)
	}

	return data
}()

func (p *Page) CreateSVG() string {
	parts := utils.PadLeft(strconv.Itoa(int(p.Visits+1)), 5, '0')

	var svg strings.Builder

	svg.WriteString(`<svg xmlns="http://www.w3.org/2000/svg" width="225" height="100">`)

	for i := 0; i < 5; i++ {
		digit := parts[i] - '0'
		x := i * 45
		svg.WriteString(fmt.Sprintf(
			`<image x="%d" y="0" width="45" height="100" href="data:image/gif;base64,%s" />`,
			x, imagesBase64[digit],
		))
	}

	svg.WriteString(`</svg>`)

	return svg.String()
}
