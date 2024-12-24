package finhelpctgservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/finhelpctgschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type FinhelpCtgService struct {
	db           *sql.DB
	finhelpCtgTN string
}

func NewFinhelpCtgService(db *sql.DB, finhelpCtgTN string) (fctgs *FinhelpCtgService) {
	if db == nil || finhelpCtgTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &FinhelpCtgService{
		db:           db,
		finhelpCtgTN: finhelpCtgTN,
	}
}

func (fctgs *FinhelpCtgService) InsertFinhelpCtg(data *finhelpctgschema.FinhelpCtgInsert) (finhelpCtgDB *finhelpctgschema.FinhelpCtgDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[1]sDescription,
		%[1]sPayment
	)
	VALUES (
		"%[2]s",
		%[3]d
	)
RETURNING
	%[1]sID,
	%[1]sDescription,
	%[1]sPayment,
	%[1]sIsHidden`,
		fctgs.finhelpCtgTN,
		data.FinhelpCtgDescription,
		data.FinhelpCtgPayment,
	)

	stmt, err := fctgs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	finhelpCtgDB = &finhelpctgschema.FinhelpCtgDB{}
	err = stmt.QueryRow().Scan(
		&finhelpCtgDB.FinhelpCtgID,
		&finhelpCtgDB.FinhelpCtgDescription,
		&finhelpCtgDB.FinhelpCtgPayment,
		&finhelpCtgDB.FinhelpCtgIsHidden,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return finhelpCtgDB, nil
}

func (fctgs *FinhelpCtgService) UpdateFinhelpCtg(data *finhelpctgschema.FinhelpCtgUpdate) (finhelpCtgDB *finhelpctgschema.FinhelpCtgDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[1]sDescription = "%[3]s",
	%[1]sPayment = %[4]d
WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sDescription,
	%[1]sPayment,
	%[1]sIsHidden`,
		fctgs.finhelpCtgTN,
		data.FinhelpCtgID,
		data.FinhelpCtgDescription,
		data.FinhelpCtgPayment,
	)

	stmt, err := fctgs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	finhelpCtgDB = &finhelpctgschema.FinhelpCtgDB{}
	err = stmt.QueryRow().Scan(
		&finhelpCtgDB.FinhelpCtgID,
		&finhelpCtgDB.FinhelpCtgDescription,
		&finhelpCtgDB.FinhelpCtgPayment,
		&finhelpCtgDB.FinhelpCtgIsHidden,
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

	return finhelpCtgDB, nil
}

// Toggle IsHidden
func (fctgs *FinhelpCtgService) DeleteFinhelpCtg(data *finhelpctgschema.FinhelpCtgDelete) (finhelpCtgDB *finhelpctgschema.FinhelpCtgDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s
	SET %[1]sIsHidden = 1 - %[1]sIsHidden
	WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sDescription,
	%[1]sPayment,
	%[1]sIsHidden`,
		fctgs.finhelpCtgTN,
		data.FinhelpCtgID,
	)

	stmt, err := fctgs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	finhelpCtgDB = &finhelpctgschema.FinhelpCtgDB{}
	err = stmt.QueryRow().Scan(
		&finhelpCtgDB.FinhelpCtgID,
		&finhelpCtgDB.FinhelpCtgDescription,
		&finhelpCtgDB.FinhelpCtgPayment,
		&finhelpCtgDB.FinhelpCtgIsHidden,
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

	return finhelpCtgDB, nil
}

func (fctgs *FinhelpCtgService) GetFinhelpCtg(data *finhelpctgschema.FinhelpCtgGet) (finhelpCtgDB *finhelpctgschema.FinhelpCtgDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sDescription,
	%[1]sPayment,
	%[1]sIsHidden
FROM %[1]s
WHERE %[1]sID = %[2]d`,
		fctgs.finhelpCtgTN,
		data.FinhelpCtgID,
	)

	stmt, err := fctgs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	finhelpCtgDB = &finhelpctgschema.FinhelpCtgDB{}
	err = stmt.QueryRow().Scan(
		&finhelpCtgDB.FinhelpCtgID,
		&finhelpCtgDB.FinhelpCtgDescription,
		&finhelpCtgDB.FinhelpCtgPayment,
		&finhelpCtgDB.FinhelpCtgIsHidden,
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

	return finhelpCtgDB, nil
}

func (fctgs *FinhelpCtgService) GetFinhelpCtgs(data *finhelpctgschema.FinhelpCtgsGet) (finhelpCtgsDB []*finhelpctgschema.FinhelpCtgDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sDescription,
	%[1]sPayment,
	%[1]sIsHidden
FROM %[1]s
ORDER BY %[1]sID DESC`,
		fctgs.finhelpCtgTN,
	)

	stmt, err := fctgs.db.Prepare(query)
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

	finhelpCtgsDB = []*finhelpctgschema.FinhelpCtgDB{}
	finhelpCtgDB := &finhelpctgschema.FinhelpCtgDB{}
	for rows.Next() {
		err = rows.Scan(
			&finhelpCtgDB.FinhelpCtgID,
			&finhelpCtgDB.FinhelpCtgDescription,
			&finhelpCtgDB.FinhelpCtgPayment,
			&finhelpCtgDB.FinhelpCtgIsHidden,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		finhelpCtgsDB = append(finhelpCtgsDB, finhelpCtgDB)
		finhelpCtgDB = &finhelpctgschema.FinhelpCtgDB{}
	}

	return finhelpCtgsDB, nil
}
