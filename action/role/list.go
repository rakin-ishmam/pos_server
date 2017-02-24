package role

import (
	"io"

	"github.com/rakin-ishmam/pos_server/apperr"
	"github.com/rakin-ishmam/pos_server/converter"
	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
)

// List provides action to get user list
type List struct {
	session    *mgo.Session
	reqQ       query.Role
	resPayload []DetailPayload
	who        data.User
	role       data.Role
	err        error
}

// Do takes necessary steps to fetch role data
func (l *List) Do() {
	if err := l.AccessValidate(); err != nil {
		l.err = err
		return
	}

	if err := l.Validate(); err != nil {
		l.err = err
		return
	}

	dbRole := db.Role{Session: l.session}
	dtRoles, err := dbRole.List(l.reqQ)
	if err != nil {
		l.err = apperr.Database{
			Base:   err,
			Where:  "Role",
			Action: "List",
		}
		return
	}

	l.resPayload = []DetailPayload{}
	for _, dt := range dtRoles {
		dtPayload := DetailPayload{}
		dtPayload.LoadFromData(&dt)
		l.resPayload = append(l.resPayload, dtPayload)
	}
}

// Result returns result of thte action
func (l List) Result() (io.Reader, error) {
	if l.err != nil {
		return nil, l.err
	}

	return converter.JSONtoBuff(l.resPayload)
}

// AccessValidate checks access permission
func (l *List) AccessValidate() error {
	if !l.role.RoleAccess.Can(data.AccessRead) {
		return apperr.Access{Where: "User", Permission: string(data.AccessRead)}
	}

	return nil
}

// Validate valids action data
func (l List) Validate() error {
	return nil
}

// NewList return pointer of List action
func NewList(session *mgo.Session, q query.Role, who data.User, role data.Role) *List {
	return &List{
		session: session,
		reqQ:    q,
		who:     who,
		role:    role,
	}
}
