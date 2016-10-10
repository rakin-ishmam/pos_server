package apperr

// Access stores info about Access error
type Access struct {
	Where      string
	Permission string
}

func (a Access) Error() string {
	return "err#access#" + a.Where + "#" + a.Permission
}
