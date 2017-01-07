package empty

import "io"

// JSON representes empty json action
type JSON struct {
	Err error
}

// Do generate token
func (j *JSON) Do() {
}

// Result returns result of the action
func (j JSON) Result() (io.Reader, error) {
	return nil, j.Err
}
