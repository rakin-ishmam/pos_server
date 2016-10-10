package db

import (
	r "gopkg.in/dancannon/gorethink.v2"
)

// SkipStep represents skip step
type SkipStep struct {
	Skip int
}

// Filter returns term with skip param
func (s *SkipStep) Filter(t r.Term) r.Term {
	return t.Skip(s.Skip)
}

// LimitStep represents limit step
type LimitStep struct {
	Limit int
}

// Filter returns term with lim9t param
func (l *LimitStep) Filter(t r.Term) r.Term {
	return t.Limit(l.Limit)
}
