package apperr

// Validation stores information about validation error
type Validation struct {
	Where string
	Field string
	Cause string
}

// Error returns formated error string
func (v Validation) Error() string {
	return "err#validation#" + v.Where + "#" + v.Field + "#" + v.Cause
}
