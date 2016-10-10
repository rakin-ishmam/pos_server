package data

// Language represents pos language
type Language string

const (
	// BanglaLanguage represents pos bangla language
	BanglaLanguage Language = "Bangla"
	// EnglishLanguage represents english language
	EnglishLanguage Language = "English"
)

// Validate returns true if language is valid
func (l Language) Validate() bool {

	switch l {
	case BanglaLanguage:
		return true
	case EnglishLanguage:
		return true
	}

	return false
}
