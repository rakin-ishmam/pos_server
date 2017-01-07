package admin

import (
	"time"

	"github.com/rakin-ishmam/pos_server/data"
	"github.com/rakin-ishmam/pos_server/db"
	"github.com/rakin-ishmam/pos_server/db/query"
	mgo "gopkg.in/mgo.v2"
)

// Create create admin user and role for the first time
type Create struct {
	Session *mgo.Session
	Err     error
}

// Do proceed action
func (c *Create) Do() {
	if hasUser, err := c.hasAnyUser(); err != nil {
		c.Err = err
		return
	} else if hasUser {
		return
	}

	dbRole := db.Role{Session: c.Session}
	dtRole := c.adminRole()
	if err := dbRole.Put(&dtRole); err != nil {
		c.Err = err
		return
	}

	dbUser := db.User{Session: c.Session}
	dtUser := c.adminUser(dtRole)
	if err := dbUser.Put(&dtUser); err != nil {
		c.Err = err
		return
	}

}

// Error returns erro of the action
func (c Create) Error() error {
	return c.Err
}

// Result reutns results of the action
func (c Create) Result() interface{} {
	return nil
}

func (c Create) adminUser(role data.Role) data.User {
	user := data.User{
		RoleID:   role.ID,
		Name:     "Admin",
		UserName: "admin",
		Password: "admin",
	}

	user.CreatedAt = time.Now()
	user.ModifiedAt = user.CreatedAt

	return user
}

func (c Create) adminRole() data.Role {
	role := data.Role{
		Name:            "Admin",
		UserAccess:      data.AccessList{data.AccessRead, data.AccessWrite, data.AccessUpdate},
		RoleAccess:      data.AccessList{data.AccessRead, data.AccessWrite, data.AccessUpdate},
		CategoryAccess:  data.AccessList{data.AccessRead, data.AccessWrite, data.AccessUpdate},
		CustomerAccess:  data.AccessList{data.AccessRead, data.AccessWrite, data.AccessUpdate},
		InventoryAccess: data.AccessList{data.AccessRead, data.AccessWrite, data.AccessUpdate},
		ProductAccess:   data.AccessList{data.AccessRead, data.AccessWrite, data.AccessUpdate},
		SellAccess:      data.AccessList{data.AccessRead, data.AccessWrite, data.AccessUpdate},
		PaymentAccess:   data.AccessList{data.AccessRead, data.AccessWrite, data.AccessUpdate},
		FileAccess:      data.AccessList{data.AccessRead, data.AccessWrite, data.AccessUpdate},
	}

	role.CreatedAt = time.Now()
	role.ModifiedAt = role.CreatedAt

	return role
}

func (c Create) hasAnyUser() (bool, error) {
	dbUser := db.User{Session: c.Session}
	users, err := dbUser.List(0, 1, []query.Applier{})

	if err != nil {
		return false, err
	}

	if len(users) > 0 {
		return true, nil
	}

	return false, nil
}
