package sessionschema

import (
	"fmt"

	"github.com/google/uuid"
)

type SessionDB struct {
	SessionID        int    `json:"SessionID"`
	SessionUUID      string `json:"SessionUUID"`
	UserID           int    `json:"UserID"`
	SessionCreatedAt string `json:"SessionCreatedAt"`
}

type SessionInsert struct {
	UserID int `json:"UserID"`
}

type SessionDelete struct {
	SessionUUID string `json:"SessionUUID"`
}

type SessionGet struct {
	SessionUUID string `json:"SessionUUID"`
}

type SessionsGet struct {
	UserID int `json:"UserID"`
}

func ValidateSessionDB(sessionDB *SessionDB) (err error) {
	if sessionDB == nil {
		return fmt.Errorf("Object is nil")
	}
	err = uuid.Validate(sessionDB.SessionUUID)
	if sessionDB.SessionID <= 0 || sessionDB.UserID <= 0 || err != nil {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateSessionInsert(sessionInsert *SessionInsert) (err error) {
	if sessionInsert == nil {
		return fmt.Errorf("Object is nil")
	}
	return nil
}

func ValidateSessionDelete(sessionDelete *SessionDelete) (err error) {
	if sessionDelete == nil {
		return fmt.Errorf("Object is nil")
	}
	err = uuid.Validate(sessionDelete.SessionUUID)
	if err != nil {
		return fmt.Errorf("SessionUUID is not valid UUID")
	}
	return nil
}

func ValidateSessionGet(sessionGet *SessionGet) (err error) {
	if sessionGet == nil {
		return fmt.Errorf("Object is nil")
	}
	err = uuid.Validate(sessionGet.SessionUUID)
	if err != nil {
		return fmt.Errorf("SessionUUID is not valid UUID")
	}
	return nil
}

func ValidateSessionsGet(sessionsGet *SessionsGet) (err error) {
	if sessionsGet == nil {
		return fmt.Errorf("Object is nil")
	}
	if sessionsGet.UserID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}
