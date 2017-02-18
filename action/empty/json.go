package empty

import "io"

// JSON representes empty json action
type JSON struct {
	err error
}

// Do generate token
func (j *JSON) Do() {
}

// Result returns result of the action
func (j JSON) Result() (io.Reader, error) {
	return nil, j.err
}

// NewJSON return empty JSON action
func NewJSON(err error) *JSON {
	return &JSON{err: err}
}
