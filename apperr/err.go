package apperr

// Access stores info about Access error
type Access struct {
	Where      string
	Permission string
}

func (a Access) Error() string {
	return "err#access#" + a.Where + "#" + a.Permission
}

// Authentication provides authentication error
type Authentication struct {
	Where string
	Cause string
}

func (a Authentication) Error() string {
	return "#err#authentication#" + a.Where + "#" + a.Cause
}

// Database stores database error
type Database struct {
	Base   error
	Where  string
	Action string
}

func (d Database) Error() string {
	return "err#database#" + d.Where + "#" + d.Action + "#" + d.Base.Error()
}

// Internal stores information about internal error
type Internal struct {
	Where string
}

// Error returns formated error string
func (i Internal) Error() string {
	return "err#internal#" + i.Where
}

// Notfound represents not found error
type Notfound struct {
	Where string
	What  string
}

func (n Notfound) Error() string {
	return "err#notfound#" + n.Where + "#" + n.What
}

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

// Parse representes parse error
type Parse struct {
	Where string
	What  string
}

func (p Parse) Error() string {
	return "err#parse#" + p.Where + "#" + p.What
}
