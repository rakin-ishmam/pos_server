package action

import "io"

// JSONAction provides functionalites for json action
type JSONAction interface {
	Do()
	AccessValidate() error
	Result() (io.Reader, error)
}
