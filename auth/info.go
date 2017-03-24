package auth

import "time"

// Info stores token information
type Info struct {
	Username string
	Exp      time.Time
}
