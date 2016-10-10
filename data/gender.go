package data

// Gender representes gender of use or customer or others
type Gender string

const (
	// GenderMale represents male gender
	GenderMale Gender = "male"

	// GenderFemale represents female gender
	GenderFemale Gender = "female"
)

// Validate valid gender type
func (g Gender) Validate() bool {
	switch g {
	case GenderMale:
		return true
	case GenderFemale:
		return true
	}

	return false
}
