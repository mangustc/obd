package sessionservice

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/sessionschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type SessionService struct {
	db        *sql.DB
	sessionTN string
	userTN    string
}

func NewSessionService(db *sql.DB, sessionTN string, userTN string) (ss *SessionService) {
	if db == nil || sessionTN == "" || userTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &SessionService{
		db:        db,
		sessionTN: sessionTN,
		userTN:    userTN,
	}
}

func (ss *SessionService) InsertSession(data *sessionschema.SessionInsert) (sessionDB *sessionschema.SessionDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[1]sUUID,
		%[2]sID
	)
	VALUES (
		"%[3]s",
		%[4]d
	)
RETURNING
	%[1]sID,
	%[1]sUUID,
	%[2]sID,
	%[1]sCreatedAt`,
		ss.sessionTN,
		ss.userTN,
		uuid.New().String(),
		data.UserID,
	)

	stmt, err := ss.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	sessionDB = &sessionschema.SessionDB{}
	err = stmt.QueryRow().Scan(
		&sessionDB.SessionID,
		&sessionDB.SessionUUID,
		&sessionDB.UserID,
		&sessionDB.SessionCreatedAt,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return sessionDB, nil
}

func (ss *SessionService) DeleteSession(data *sessionschema.SessionDelete) (sessionDB *sessionschema.SessionDB, err error) {
	query := fmt.Sprintf(`
DELETE FROM %[1]s WHERE %[1]sUUID = "%[2]s"
RETURNING *`,
		ss.sessionTN,
		data.SessionUUID,
	)

	stmt, err := ss.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	sessionDB = &sessionschema.SessionDB{}
	err = stmt.QueryRow().Scan(
		&sessionDB.SessionID,
		&sessionDB.SessionUUID,
		&sessionDB.UserID,
		&sessionDB.SessionCreatedAt,
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

	err = sessionschema.ValidateSessionDB(sessionDB)
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return sessionDB, nil
}

func (ss *SessionService) GetSession(data *sessionschema.SessionGet) (sessionDB *sessionschema.SessionDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sUUID,
	%[2]sID,
	%[1]sCreatedAt
FROM %[1]s
WHERE %[1]sUUID = "%[3]s"`,
		ss.sessionTN,
		ss.userTN,
		data.SessionUUID,
	)

	stmt, err := ss.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	sessionDB = &sessionschema.SessionDB{}
	err = stmt.QueryRow().Scan(
		&sessionDB.SessionID,
		&sessionDB.SessionUUID,
		&sessionDB.UserID,
		&sessionDB.SessionCreatedAt,
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

	return sessionDB, nil
}

func (ss *SessionService) GetSessions(data *sessionschema.SessionsGet) (sessionsDB []*sessionschema.SessionDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sUUID,
	%[2]sID,
	%[1]sCreatedAt
FROM %[1]s
WHERE %[2]sID = %[3]d
ORDER BY %[1]sID DESC`,
		ss.sessionTN,
		ss.userTN,
		data.UserID,
	)

	stmt, err := ss.db.Prepare(query)
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

	sessionsDB = []*sessionschema.SessionDB{}
	sessionDB := &sessionschema.SessionDB{}
	for rows.Next() {
		err = rows.Scan(
			&sessionDB.SessionID,
			&sessionDB.SessionUUID,
			&sessionDB.UserID,
			&sessionDB.SessionCreatedAt,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		sessionsDB = append(sessionsDB, sessionDB)
		sessionDB = &sessionschema.SessionDB{}
	}

	return sessionsDB, nil
}
