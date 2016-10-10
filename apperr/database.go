package apperr

// Database stores database error
type Database struct {
	Base   error
	Where  string
	Action string
}

func (d Database) Error() string {
	return "err#database#" + d.Where + "#" + d.Action + "#" + d.Base.Error()
}
