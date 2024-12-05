package userschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type UserDB struct {
	UserID         int    `json:"UserID"`
	UserLastname   string `json:"UserLastname"`
	UserFirstname  string `json:"UserFirstname"`
	UserMiddlename string `json:"UserMiddlename"`
	UserPassword   string `json:"UserPassword"`
	UserIsHidden   bool   `json:"UserIsHidden"`
	JobID          int    `json:"JobID"`
}

type UserInsert struct {
	UserLastname   string `json:"UserLastname"`
	UserFirstname  string `json:"UserFirstname"`
	UserMiddlename string `json:"UserMiddlename"`
	UserPassword   string `json:"UserPassword"`
	JobID          int    `json:"JobID"`
}

type UserUpdate struct {
	UserID         int    `json:"UserID"`
	UserLastname   string `json:"UserLastname"`
	UserFirstname  string `json:"UserFirstname"`
	UserMiddlename string `json:"UserMiddlename"`
	UserPassword   string `json:"UserPassword"`
	JobID          int    `json:"JobID"`
}

type UserDelete struct {
	UserID int `json:"UserID"`
}

type UserGet struct {
	UserID int `json:"UserID"`
}

type UsersGet struct{}

func GetUserInputOptionsFromUsersDB(usersDB []*UserDB) []*schema.InputOption {
	notHiddenUsersDB := GetNotHiddenUsersDB(usersDB)
	inputOptions := []*schema.InputOption{}
	for _, userDB := range notHiddenUsersDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%s %s %s", userDB.UserLastname, userDB.UserFirstname, userDB.UserMiddlename),
			InputOptionValue: fmt.Sprintf("%d", userDB.UserID),
		})
	}
	return inputOptions
}

func GetNotHiddenUsersDB(usersDB []*UserDB) []*UserDB {
	notHiddenUsersDB := []*UserDB{}
	for _, userDB := range usersDB {
		if !userDB.UserIsHidden {
			notHiddenUsersDB = append(notHiddenUsersDB, userDB)
		}
	}
	return notHiddenUsersDB
}

func ValidateUserDB(userDB *UserDB) (err error) {
	if userDB == nil {
		return fmt.Errorf("Object is nil")
	}
	if userDB.UserID <= 0 || userDB.UserLastname == "" || userDB.UserFirstname == "" ||
		userDB.UserPassword == "" || userDB.JobID <= 0 {
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

func ValidateUserGet(userGet *UserGet) (err error) {
	if userGet == nil {
		return fmt.Errorf("Object is nil")
	}
	if userGet.UserID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateUsersGet(userGets *UsersGet) (err error) {
	if userGets == nil {
		return fmt.Errorf("Object is nil")
	}
	return nil
}
