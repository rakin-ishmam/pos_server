package auth

import "time"

// Info stores token information
type Info struct {
	UserName string
	Exp      time.Time
}
