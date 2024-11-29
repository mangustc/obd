package userschema

import (
	"fmt"
	"time"

	"github.com/mangustc/obd/util"
)

type UserDB struct {
	UserID         int
	UserLastname   string
	UserFirstname  string
	UserMiddlename string
	UserPassword   string
	UserCreatedAt  time.Time
	UserIsHidden   bool
	JobID          int
}

type UserInsert struct {
	UserLastname   string
	UserFirstname  string
	UserMiddlename string
	UserPassword   string
	JobID          int
}

type UserUpdate struct {
	UserID         int
	UserLastname   string
	UserFirstname  string
	UserMiddlename string
	UserPassword   string
	JobID          int
}

type UserDelete struct {
	UserID int
}

func ValidateUserDB(userDB *UserDB) (err error) {
	if userDB == nil {
		return fmt.Errorf("Object is nil")
	}
	if userDB.UserID <= 0 || userDB.UserLastname == "" || userDB.UserFirstname == "" || userDB.UserPassword == "" ||
		util.IsZero(userDB.UserCreatedAt) || userDB.JobID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateUserInsert(userInsert *UserInsert) (err error) {
	if userInsert == nil {
		return fmt.Errorf("Object is nil")
	}
	if userInsert.UserLastname == "" || userInsert.UserFirstname == "" ||
		userInsert.UserPassword == "" || userInsert.JobID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateUserUpdate(userUpdate *UserUpdate) (err error) {
	if userUpdate == nil {
		return fmt.Errorf("Object is nil")
	}
	if userUpdate.UserID <= 0 || userUpdate.UserLastname == "" || userUpdate.UserFirstname == "" ||
		userUpdate.UserPassword == "" || userUpdate.JobID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateUserDelete(userDelete *UserDelete) (err error) {
	if userDelete == nil {
		return fmt.Errorf("Object is nil")
	}
	if userDelete.UserID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}
