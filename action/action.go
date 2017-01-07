package action

import "io"

// JSONAction provides functionalites for json action
type JSONAction interface {
	Do()
	Result() (io.Reader, error)
}
