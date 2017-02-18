package apperr

// Access stores info about Access error
type Access struct {
	Where      string
	Permission string
}

func (a Access) Error() string {
	return "err#access#" + a.Where + "#" + a.Permission
}

// NewAccess returns access error
func NewAccess(where, permission string) Access {
	return Access{
		Where:      where,
		Permission: permission,
	}
}

// Authentication provides authentication error
type Authentication struct {
	Where string
	Cause string
}

func (a Authentication) Error() string {
	return "#err#authentication#" + a.Where + "#" + a.Cause
}

// NewAuthentication returns authentication error
func NewAuthentication(where, cause string) Authentication {
	return Authentication{
		Where: where,
		Cause: cause,
	}
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

// NewDatabase return Database error
func NewDatabase(base error, where, action string) Database {
	return Database{
		Base:   base,
		Where:  where,
		Action: action,
	}
}

// Internal stores information about internal error
type Internal struct {
	Where string
}

// Error returns formated error string
func (i Internal) Error() string {
	return "err#internal#" + i.Where
}

// NewInternal return Internal server error
func NewInternal(where string) Internal {
	return Internal{
		Where: where,
	}
}

// Notfound represents not found error
type Notfound struct {
	Where string
	What  string
}

func (n Notfound) Error() string {
	return "err#notfound#" + n.Where + "#" + n.What
}

// NewNotfound returns notfound error
func NewNotfound(where, what string) Notfound {
	return Notfound{
		Where: where,
		What:  what,
	}
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

// NewValidation returns validation error
func NewValidation(where, field, cause string) Validation {
	return Validation{
		Where: where,
		Field: field,
		Cause: cause,
	}
}

// Parse representes parse error
type Parse struct {
	Where string
	What  string
}

func (p Parse) Error() string {
	return "err#parse#" + p.Where + "#" + p.What
}

// NewParse returns parse error
func NewParse(where, what string) Parse {
	return Parse{
		Where: where,
		What:  what,
	}
}
