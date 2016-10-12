package data

// File provides file info
type File struct {
	Track

	Location string `bson:"location"`
}

// Validate valids File data
func (f File) Validate() error {
	return nil
}
