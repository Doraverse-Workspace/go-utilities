package language

import "github.com/goccy/go-json"

type Multilingual struct {
	English    string `json:"en,omitempty"` // English
	Vietnamese string `json:"vi,omitempty"` // Vietnamese
}

func (m Multilingual) GetLocalized(lang string) string {
	switch lang {
	case English.String():
		return m.English
	default:
		return m.Vietnamese
	}
}

func (m Multilingual) IsEmpty() bool {
	return m.English == "" && m.Vietnamese == ""
}

func (m Multilingual) ToString() (string, error) {
	if data, err := json.Marshal(m); err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}
