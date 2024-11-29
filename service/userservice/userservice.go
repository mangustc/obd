package userservice

import "database/sql"

type UserService struct {
	db        *sql.DB
	userTN    string
	sessionTN string
	jobTN     string
}

func NewUserService(db *sql.DB, userTN string, sessionTN string, jobTN string) (us *UserService) {
	if db == nil || userTN == "" || sessionTN == "" || jobTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &UserService{
		db:        db,
		userTN:    userTN,
		sessionTN: sessionTN,
		jobTN:     jobTN,
	}
}
