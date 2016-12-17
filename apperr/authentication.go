package apperr

// Authentication provides authentication error
type Authentication struct {
	Where string
	Cause string
}

func (a Authentication) Error() string {
	return "#err#authentication#" + a.Where + "#" + a.Cause
}
