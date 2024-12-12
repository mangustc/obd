package skipservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/skipschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type SkipService struct {
	db        *sql.DB
	skipTN    string
	classTN   string
	studentTN string
}

func NewSkipService(db *sql.DB,
	skipTN string,
	classTN string,
	studentTN string,
) (sks *SkipService) {
	if db == nil ||
		skipTN == "" ||
		classTN == "" ||
		studentTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &SkipService{
		db:        db,
		skipTN:    skipTN,
		classTN:   classTN,
		studentTN: studentTN,
	}
}

func (sks *SkipService) InsertSkip(data *skipschema.SkipInsert) (skipDB *skipschema.SkipDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[2]sID,
		%[3]sID
	)
	VALUES (
		%[4]d,
		%[5]d
	)
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID`,
		sks.skipTN,
		sks.classTN,
		sks.studentTN,
		data.ClassID,
		data.StudentID,
	)

	stmt, err := sks.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	skipDB = &skipschema.SkipDB{}
	err = stmt.QueryRow().Scan(
		&skipDB.SkipID,
		&skipDB.ClassID,
		&skipDB.StudentID,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return skipDB, nil
}

func (sks *SkipService) UpdateSkip(data *skipschema.SkipUpdate) (skipDB *skipschema.SkipDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
		%[2]sID = %[5]d,
		%[3]sID = %[6]d
WHERE %[1]sID = %[4]d
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID`,
		sks.skipTN,
		sks.classTN,
		sks.studentTN,
		data.SkipID,
		data.ClassID,
		data.StudentID,
	)

	stmt, err := sks.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	skipDB = &skipschema.SkipDB{}
	err = stmt.QueryRow().Scan(
		&skipDB.SkipID,
		&skipDB.ClassID,
		&skipDB.StudentID,
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

	return skipDB, nil
}

// Toggle IsHidden
func (sks *SkipService) DeleteSkip(data *skipschema.SkipDelete) (skipDB *skipschema.SkipDB, err error) {
	query := fmt.Sprintf(`
DELETE FROM %[1]s WHERE %[1]sID = %[4]d
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID`,
		sks.skipTN,
		sks.classTN,
		sks.studentTN,
		data.SkipID,
	)

	stmt, err := sks.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	skipDB = &skipschema.SkipDB{}
	err = stmt.QueryRow().Scan(
		&skipDB.SkipID,
		&skipDB.ClassID,
		&skipDB.StudentID,
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

	return skipDB, nil
}

func (sks *SkipService) GetSkip(data *skipschema.SkipGet) (skipDB *skipschema.SkipDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[2]sID,
	%[3]sID
FROM %[1]s
WHERE %[1]sID = %[4]d`,
		sks.skipTN,
		sks.classTN,
		sks.studentTN,
		data.SkipID,
	)

	stmt, err := sks.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	skipDB = &skipschema.SkipDB{}
	err = stmt.QueryRow().Scan(
		&skipDB.SkipID,
		&skipDB.ClassID,
		&skipDB.StudentID,
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

	return skipDB, nil
}

func (sks *SkipService) GetSkips(data *skipschema.SkipsGet) (skipsDB []*skipschema.SkipDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[2]sID,
	%[3]sID
FROM %[1]s
ORDER BY %[1]sID DESC`,
		sks.skipTN,
		sks.classTN,
		sks.studentTN,
	)

	stmt, err := sks.db.Prepare(query)
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

	skipsDB = []*skipschema.SkipDB{}
	skipDB := &skipschema.SkipDB{}
	for rows.Next() {
		err = rows.Scan(
			&skipDB.SkipID,
			&skipDB.ClassID,
			&skipDB.StudentID,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		skipsDB = append(skipsDB, skipDB)
		skipDB = &skipschema.SkipDB{}
	}

	return skipsDB, nil
}
