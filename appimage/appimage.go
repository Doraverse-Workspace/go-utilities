package appimage

import (
	"fmt"

	"github.com/goccy/go-json"
)

type Image struct {
	ID string `json:"id"`
	Sm string `json:"sm"`
	Md string `json:"md"`
}

func (i Image) GetURL(fileHost string, isMobile bool) string {
	var name = i.Md
	if isMobile {
		name = i.Sm
	}
	return fmt.Sprintf("%s/%s", fileHost, name)
}

func (i Image) ToString() (string, error) {
	if data, err := json.Marshal(i); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}
