package action

import (
	"io"

	"github.com/rakin-ishmam/pos_server/data"
)

// JSONAction provides functionalites for json action
type JSONAction interface {
	Do()
	Result() (io.Reader, error)
}

// TokenAction provides functionalites for token actiion
type TokenAction interface {
	Do()
	Result() (data.User, data.Role, error)
}
