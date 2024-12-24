package finhelpstageservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/finhelpstageschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type FinhelpStageService struct {
	db             *sql.DB
	finhelpStageTN string
}

func NewFinhelpStageService(db *sql.DB, finhelpStageTN string) (fsts *FinhelpStageService) {
	if db == nil || finhelpStageTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &FinhelpStageService{
		db:             db,
		finhelpStageTN: finhelpStageTN,
	}
}

func (fsts *FinhelpStageService) InsertFinhelpStage(data *finhelpstageschema.FinhelpStageInsert) (finhelpStageDB *finhelpstageschema.FinhelpStageDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[1]sName,
		%[1]sDescription
	)
	VALUES (
		"%[2]s",
		"%[3]s"
	)
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sDescription,
	%[1]sIsHidden`,
		fsts.finhelpStageTN,
		data.FinhelpStageName,
		data.FinhelpStageDescription,
	)

	stmt, err := fsts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	finhelpStageDB = &finhelpstageschema.FinhelpStageDB{}
	err = stmt.QueryRow().Scan(
		&finhelpStageDB.FinhelpStageID,
		&finhelpStageDB.FinhelpStageName,
		&finhelpStageDB.FinhelpStageDescription,
		&finhelpStageDB.FinhelpStageIsHidden,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return finhelpStageDB, nil
}

func (fsts *FinhelpStageService) UpdateFinhelpStage(data *finhelpstageschema.FinhelpStageUpdate) (finhelpStageDB *finhelpstageschema.FinhelpStageDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[1]sName = "%[3]s",
	%[1]sDescription = "%[4]s"
WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sDescription,
	%[1]sIsHidden`,
		fsts.finhelpStageTN,
		data.FinhelpStageID,
		data.FinhelpStageName,
		data.FinhelpStageDescription,
	)

	stmt, err := fsts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	finhelpStageDB = &finhelpstageschema.FinhelpStageDB{}
	err = stmt.QueryRow().Scan(
		&finhelpStageDB.FinhelpStageID,
		&finhelpStageDB.FinhelpStageName,
		&finhelpStageDB.FinhelpStageDescription,
		&finhelpStageDB.FinhelpStageIsHidden,
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

	return finhelpStageDB, nil
}

// Toggle IsHidden
func (fsts *FinhelpStageService) DeleteFinhelpStage(data *finhelpstageschema.FinhelpStageDelete) (finhelpStageDB *finhelpstageschema.FinhelpStageDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s
	SET %[1]sIsHidden = 1 - %[1]sIsHidden
	WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sDescription,
	%[1]sIsHidden`,
		fsts.finhelpStageTN,
		data.FinhelpStageID,
	)

	stmt, err := fsts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	finhelpStageDB = &finhelpstageschema.FinhelpStageDB{}
	err = stmt.QueryRow().Scan(
		&finhelpStageDB.FinhelpStageID,
		&finhelpStageDB.FinhelpStageName,
		&finhelpStageDB.FinhelpStageDescription,
		&finhelpStageDB.FinhelpStageIsHidden,
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

	return finhelpStageDB, nil
}

func (fsts *FinhelpStageService) GetFinhelpStage(data *finhelpstageschema.FinhelpStageGet) (finhelpStageDB *finhelpstageschema.FinhelpStageDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sName,
	%[1]sDescription,
	%[1]sIsHidden
FROM %[1]s
WHERE %[1]sID = %[2]d`,
		fsts.finhelpStageTN,
		data.FinhelpStageID,
	)

	stmt, err := fsts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	finhelpStageDB = &finhelpstageschema.FinhelpStageDB{}
	err = stmt.QueryRow().Scan(
		&finhelpStageDB.FinhelpStageID,
		&finhelpStageDB.FinhelpStageName,
		&finhelpStageDB.FinhelpStageDescription,
		&finhelpStageDB.FinhelpStageIsHidden,
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

	return finhelpStageDB, nil
}

func (fsts *FinhelpStageService) GetFinhelpStages(data *finhelpstageschema.FinhelpStagesGet) (finhelpStagesDB []*finhelpstageschema.FinhelpStageDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sName,
	%[1]sDescription,
	%[1]sIsHidden
FROM %[1]s
ORDER BY %[1]sID DESC`,
		fsts.finhelpStageTN,
	)

	stmt, err := fsts.db.Prepare(query)
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

	finhelpStagesDB = []*finhelpstageschema.FinhelpStageDB{}
	finhelpStageDB := &finhelpstageschema.FinhelpStageDB{}
	for rows.Next() {
		err = rows.Scan(
			&finhelpStageDB.FinhelpStageID,
			&finhelpStageDB.FinhelpStageName,
			&finhelpStageDB.FinhelpStageDescription,
			&finhelpStageDB.FinhelpStageIsHidden,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		finhelpStagesDB = append(finhelpStagesDB, finhelpStageDB)
		finhelpStageDB = &finhelpstageschema.FinhelpStageDB{}
	}

	return finhelpStagesDB, nil
}
