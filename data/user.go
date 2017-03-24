package data

import (
	"github.com/rakin-ishmam/pos_server/apperr"
	"gopkg.in/mgo.v2/bson"
)

// User provides user information
type User struct {
	Track

	Name      string        `bson:"name"`
	Username  string        `bson:"username"`
	Language  Language      `bson:"language"`
	AvtFileID string        `bson:"avt_file_id"`
	RoleID    bson.ObjectId `bson:"role_id,omitempty"`
	Email     string        `bson:"email"`
	Address   string        `bson:"address"`
	Gender    Gender        `bson:"gender"`
	Phone     []string      `bson:"phone"`
	Password  string        `bson:"password"`
}

// PreSave takes the necessary step before saving data
func (u *User) PreSave() {
	u.Track.Search = Spliter(u.Name, u.Username)
}

// Validate valids user data
func (u User) Validate() error {
	if !u.Language.Validate() {
		return apperr.Validation{
			Where: "User",
			Field: "Language",
			Cause: apperr.StrInvalid,
		}
	}
	if !u.Gender.Validate() {
		return apperr.Validation{
			Where: "User",
			Field: "Gender",
			Cause: apperr.StrInvalid,
		}
	}
	return nil
}
