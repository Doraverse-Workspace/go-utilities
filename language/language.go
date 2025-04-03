package language

type Language string

const (
	Unknown    Language = ""
	English    Language = "en"
	Vietnamese Language = "vi"
)

func (l Language) String() string {
	return string(l)
}

func (l Language) IsValid() bool {
	return l != Unknown
}

func (l Language) IsEnglish() bool {
	return l == English
}

func (l Language) IsVietnamese() bool {
	return l == Vietnamese
}

func ToLanguage(lang string) Language {
	switch lang {
	case English.String():
		return English
	case Vietnamese.String():
		return Vietnamese
	default:
		return Unknown
	}
}
