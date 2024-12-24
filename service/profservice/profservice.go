package profservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/profschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type ProfService struct {
	db     *sql.DB
	profTN string
}

func NewProfService(db *sql.DB, profTN string) (prs *ProfService) {
	if db == nil || profTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &ProfService{
		db:     db,
		profTN: profTN,
	}
}

func (prs *ProfService) InsertProf(data *profschema.ProfInsert) (profDB *profschema.ProfDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[1]sLastname,
		%[1]sFirstname,
		%[1]sMiddlename,
		%[1]sPhoneNumber,
		%[1]sEmail
	)
	VALUES (
		"%[2]s",
		"%[3]s",
		"%[4]s",
		"%[5]s",
		"%[6]s"
	)
RETURNING
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPhoneNumber,
	%[1]sIsHidden,
	%[1]sEmail`,
		prs.profTN,
		data.ProfLastname,
		data.ProfFirstname,
		data.ProfMiddlename,
		data.ProfPhoneNumber,
		data.ProfEmail,
	)

	stmt, err := prs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	profDB = &profschema.ProfDB{}
	err = stmt.QueryRow().Scan(
		&profDB.ProfID,
		&profDB.ProfLastname,
		&profDB.ProfFirstname,
		&profDB.ProfMiddlename,
		&profDB.ProfPhoneNumber,
		&profDB.ProfIsHidden,
		&profDB.ProfEmail,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return profDB, nil
}

func (prs *ProfService) UpdateProf(data *profschema.ProfUpdate) (profDB *profschema.ProfDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[1]sLastname = "%[3]s",
	%[1]sFirstname = "%[4]s",
	%[1]sMiddlename = "%[5]s",
	%[1]sPhoneNumber = "%[6]s",
	%[1]sEmail = "%[7]s"
WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPhoneNumber,
	%[1]sIsHidden,
	%[1]sEmail`,
		prs.profTN,
		data.ProfID,
		data.ProfLastname,
		data.ProfFirstname,
		data.ProfMiddlename,
		data.ProfPhoneNumber,
		data.ProfEmail,
	)

	stmt, err := prs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	profDB = &profschema.ProfDB{}
	err = stmt.QueryRow().Scan(
		&profDB.ProfID,
		&profDB.ProfLastname,
		&profDB.ProfFirstname,
		&profDB.ProfMiddlename,
		&profDB.ProfPhoneNumber,
		&profDB.ProfIsHidden,
		&profDB.ProfEmail,
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

	return profDB, nil
}

// Toggle IsHidden
func (prs *ProfService) DeleteProf(data *profschema.ProfDelete) (profDB *profschema.ProfDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s
	SET %[1]sIsHidden = 1 - %[1]sIsHidden
	WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPhoneNumber,
	%[1]sIsHidden,
	%[1]sEmail`,
		prs.profTN,
		data.ProfID,
	)

	stmt, err := prs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	profDB = &profschema.ProfDB{}
	err = stmt.QueryRow().Scan(
		&profDB.ProfID,
		&profDB.ProfLastname,
		&profDB.ProfFirstname,
		&profDB.ProfMiddlename,
		&profDB.ProfPhoneNumber,
		&profDB.ProfIsHidden,
		&profDB.ProfEmail,
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

	return profDB, nil
}

func (prs *ProfService) GetProf(data *profschema.ProfGet) (profDB *profschema.ProfDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPhoneNumber,
	%[1]sIsHidden,
	%[1]sEmail
FROM %[1]s
WHERE %[1]sID = %[2]d`,
		prs.profTN,
		data.ProfID,
	)

	stmt, err := prs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	profDB = &profschema.ProfDB{}
	err = stmt.QueryRow().Scan(
		&profDB.ProfID,
		&profDB.ProfLastname,
		&profDB.ProfFirstname,
		&profDB.ProfMiddlename,
		&profDB.ProfPhoneNumber,
		&profDB.ProfIsHidden,
		&profDB.ProfEmail,
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

	return profDB, nil
}

func (prs *ProfService) GetProfs(data *profschema.ProfsGet) (profsDB []*profschema.ProfDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPhoneNumber,
	%[1]sIsHidden,
	%[1]sEmail
FROM %[1]s
ORDER BY %[1]sID DESC`,
		prs.profTN,
	)

	stmt, err := prs.db.Prepare(query)
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

	profsDB = []*profschema.ProfDB{}
	profDB := &profschema.ProfDB{}
	for rows.Next() {
		err = rows.Scan(
			&profDB.ProfID,
			&profDB.ProfLastname,
			&profDB.ProfFirstname,
			&profDB.ProfMiddlename,
			&profDB.ProfPhoneNumber,
			&profDB.ProfIsHidden,
			&profDB.ProfEmail,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		profsDB = append(profsDB, profDB)
		profDB = &profschema.ProfDB{}
	}

	return profsDB, nil
}
