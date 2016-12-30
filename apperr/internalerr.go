package apperr

// Internal stores information about internal error
type Internal struct {
	Where string
}

// Error returns formated error string
func (i Internal) Error() string {
	return "err#internal#" + i.Where
}
