package cabinettypeservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/cabinettypeschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type CabinetTypeService struct {
	db            *sql.DB
	cabinetTypeTN string
}

func NewCabinetTypeService(db *sql.DB, cabinetTypeTN string) (cts *CabinetTypeService) {
	if db == nil || cabinetTypeTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &CabinetTypeService{
		db:            db,
		cabinetTypeTN: cabinetTypeTN,
	}
}

func (cts *CabinetTypeService) InsertCabinetType(data *cabinettypeschema.CabinetTypeInsert) (cabinetTypeDB *cabinettypeschema.CabinetTypeDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[1]sName
	)
	VALUES (
		"%[2]s"
	)
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden`,
		cts.cabinetTypeTN,
		data.CabinetTypeName,
	)

	stmt, err := cts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	cabinetTypeDB = &cabinettypeschema.CabinetTypeDB{}
	err = stmt.QueryRow().Scan(
		&cabinetTypeDB.CabinetTypeID,
		&cabinetTypeDB.CabinetTypeName,
		&cabinetTypeDB.CabinetTypeIsHidden,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return cabinetTypeDB, nil
}

func (cts *CabinetTypeService) UpdateCabinetType(data *cabinettypeschema.CabinetTypeUpdate) (cabinetTypeDB *cabinettypeschema.CabinetTypeDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[1]sName = "%[3]s"
WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden`,
		cts.cabinetTypeTN,
		data.CabinetTypeID,
		data.CabinetTypeName,
	)

	stmt, err := cts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	cabinetTypeDB = &cabinettypeschema.CabinetTypeDB{}
	err = stmt.QueryRow().Scan(
		&cabinetTypeDB.CabinetTypeID,
		&cabinetTypeDB.CabinetTypeName,
		&cabinetTypeDB.CabinetTypeIsHidden,
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

	return cabinetTypeDB, nil
}

// Toggle IsHidden
func (cts *CabinetTypeService) DeleteCabinetType(data *cabinettypeschema.CabinetTypeDelete) (cabinetTypeDB *cabinettypeschema.CabinetTypeDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s
	SET %[1]sIsHidden = 1 - %[1]sIsHidden
	WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden`,
		cts.cabinetTypeTN,
		data.CabinetTypeID,
	)

	stmt, err := cts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	cabinetTypeDB = &cabinettypeschema.CabinetTypeDB{}
	err = stmt.QueryRow().Scan(
		&cabinetTypeDB.CabinetTypeID,
		&cabinetTypeDB.CabinetTypeName,
		&cabinetTypeDB.CabinetTypeIsHidden,
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

	return cabinetTypeDB, nil
}

func (cts *CabinetTypeService) GetCabinetType(data *cabinettypeschema.CabinetTypeGet) (cabinetTypeDB *cabinettypeschema.CabinetTypeDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden
FROM %[1]s
WHERE %[1]sID = %[2]d`,
		cts.cabinetTypeTN,
		data.CabinetTypeID,
	)

	stmt, err := cts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	cabinetTypeDB = &cabinettypeschema.CabinetTypeDB{}
	err = stmt.QueryRow().Scan(
		&cabinetTypeDB.CabinetTypeID,
		&cabinetTypeDB.CabinetTypeName,
		&cabinetTypeDB.CabinetTypeIsHidden,
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

	return cabinetTypeDB, nil
}

func (cts *CabinetTypeService) GetCabinetTypes(data *cabinettypeschema.CabinetTypesGet) (cabinetTypesDB []*cabinettypeschema.CabinetTypeDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden
FROM %[1]s
ORDER BY %[1]sID DESC`,
		cts.cabinetTypeTN,
	)

	stmt, err := cts.db.Prepare(query)
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

	cabinetTypesDB = []*cabinettypeschema.CabinetTypeDB{}
	cabinetTypeDB := &cabinettypeschema.CabinetTypeDB{}
	for rows.Next() {
		err = rows.Scan(
			&cabinetTypeDB.CabinetTypeID,
			&cabinetTypeDB.CabinetTypeName,
			&cabinetTypeDB.CabinetTypeIsHidden,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		cabinetTypesDB = append(cabinetTypesDB, cabinetTypeDB)
		cabinetTypeDB = &cabinettypeschema.CabinetTypeDB{}
	}

	return cabinetTypesDB, nil
}
