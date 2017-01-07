package apperr

// Notfound represents not found error
type Notfound struct {
	Where string
	What  string
}

func (n Notfound) Error() string {
	return "err#notfound#" + n.Where + "#" + n.What
}
