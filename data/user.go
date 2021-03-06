package data

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"gopkg.in/mgo.v2/bson"
)

// User provides user information
type User struct {
	Track

	Name      string        `bson:"name"`
	UserName  string        `bson:"user_name"`
	Language  Language      `bson:"language"`
	AvtFileID string        `bson:"avt_file_id"`
	RoleID    bson.ObjectId `bson:"role_id, omitempty"`
	Email     string        `bson:"email"`
	Address   string        `bson:"address"`
	Gender    Gender        `bson:"gender"`
	Phone     []string      `bson:"phone"`
	Password  string        `bson:"password"`
}

// Validate valids user data
func (u *User) Validate() error {
	if !u.Language.Validate() {
		return apperr.Validation{
			Where: "User",
			Field: "language",
			Cause: apperr.ValidationInvalid,
		}
	}
	if !u.Gender.Validate() {
		return apperr.Validation{
			Where: "User",
			Field: "gender",
			Cause: apperr.ValidationInvalid,
		}
	}
	return nil
}
