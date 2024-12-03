package userservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/userschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type UserService struct {
	db     *sql.DB
	userTN string
	jobTN  string
}

func NewUserService(db *sql.DB, userTN string, jobTN string) (us *UserService) {
	if db == nil || userTN == "" || jobTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &UserService{
		db:     db,
		userTN: userTN,
		jobTN:  jobTN,
	}
}

func (us *UserService) InsertUser(data *userschema.UserInsert) (userDB *userschema.UserDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[1]sLastname,
		%[1]sFirstname,
		%[1]sMiddlename,
		%[1]sPassword,
		%[2]sID
	)
	VALUES (
		"%[3]s",
		"%[4]s",
		"%[5]s",
		"%[6]s",
		"%[7]d"
	)
RETURNING
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPassword,
	%[1]sIsHidden,
	%[2]sID`,
		us.userTN,
		us.jobTN,
		data.UserLastname,
		data.UserFirstname,
		data.UserMiddlename,
		data.UserPassword,
		data.JobID,
	)

	stmt, err := us.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	userDB = &userschema.UserDB{}
	err = stmt.QueryRow().Scan(
		&userDB.UserID,
		&userDB.UserLastname,
		&userDB.UserFirstname,
		&userDB.UserMiddlename,
		&userDB.UserPassword,
		&userDB.UserIsHidden,
		&userDB.JobID,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return userDB, nil
}

func (us *UserService) UpdateUser(data *userschema.UserUpdate) (userDB *userschema.UserDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[1]sLastname = "%[4]s",
	%[1]sFirstname = "%[5]s",
	%[1]sMiddlename = "%[6]s",
	%[1]sPassword = "%[7]s",
	%[2]sID = "%[8]d"
WHERE %[1]sID = %[3]d
RETURNING
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPassword,
	%[1]sIsHidden,
	%[2]sID`,
		us.userTN,
		us.jobTN,
		data.UserID,
		data.UserLastname,
		data.UserFirstname,
		data.UserMiddlename,
		data.UserPassword,
		data.JobID,
	)

	stmt, err := us.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	userDB = &userschema.UserDB{}
	err = stmt.QueryRow().Scan(
		&userDB.UserID,
		&userDB.UserLastname,
		&userDB.UserFirstname,
		&userDB.UserMiddlename,
		&userDB.UserPassword,
		&userDB.UserIsHidden,
		&userDB.JobID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ErrNotFound
		}
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return userDB, nil
}

// Toggle IsHidden
func (us *UserService) DeleteUser(data *userschema.UserDelete) (userDB *userschema.UserDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s
	SET %[1]sIsHidden = 1 - %[1]sIsHidden
	WHERE %[1]sID = %[3]d
RETURNING
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPassword,
	%[1]sIsHidden,
	%[2]sID`,
		us.userTN,
		us.jobTN,
		data.UserID,
	)

	stmt, err := us.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	userDB = &userschema.UserDB{}
	err = stmt.QueryRow().Scan(
		&userDB.UserID,
		&userDB.UserLastname,
		&userDB.UserFirstname,
		&userDB.UserMiddlename,
		&userDB.UserPassword,
		&userDB.UserIsHidden,
		&userDB.JobID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ErrNotFound
		}
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return userDB, nil
}

func (us *UserService) GetUser(data *userschema.UserGet) (userDB *userschema.UserDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPassword,
	%[1]sIsHidden,
	%[2]sID
FROM %[1]s
WHERE %[1]sID = %[3]d`,
		us.userTN,
		us.jobTN,
		data.UserID,
	)

	stmt, err := us.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	userDB = &userschema.UserDB{}
	err = stmt.QueryRow().Scan(
		&userDB.UserID,
		&userDB.UserLastname,
		&userDB.UserFirstname,
		&userDB.UserMiddlename,
		&userDB.UserPassword,
		&userDB.UserIsHidden,
		&userDB.JobID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ErrUnprocessableEntity
		}
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return userDB, nil
}

func (us *UserService) GetUsers(data *userschema.UsersGet) (usersDB []*userschema.UserDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPassword,
	%[1]sIsHidden,
	%[2]sID
FROM %[1]s
ORDER BY %[1]sID DESC`,
		us.userTN,
		us.jobTN,
	)

	stmt, err := us.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	rows, err := stmt.Query()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	usersDB = []*userschema.UserDB{}
	userDB := &userschema.UserDB{}
	for rows.Next() {
		err = rows.Scan(
			&userDB.UserID,
			&userDB.UserLastname,
			&userDB.UserFirstname,
			&userDB.UserMiddlename,
			&userDB.UserPassword,
			&userDB.UserIsHidden,
			&userDB.JobID,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		usersDB = append(usersDB, userDB)
		userDB = &userschema.UserDB{}
	}

	return usersDB, nil
}
