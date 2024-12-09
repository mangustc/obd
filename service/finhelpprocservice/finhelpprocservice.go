package finhelpprocservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/finhelpprocschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type FinhelpProcService struct {
	db             *sql.DB
	finhelpProcTN  string
	userTN         string
	studentTN      string
	finhelpCtgTN   string
	finhelpStageTN string
}

func NewFinhelpProcService(db *sql.DB,
	finhelpProcTN string,
	userTN string,
	studentTN string,
	finhelpCtgTN string,
	finhelpStageTN string,
) (fprs *FinhelpProcService) {
	if db == nil ||
		finhelpProcTN == "" ||
		userTN == "" ||
		studentTN == "" ||
		finhelpCtgTN == "" ||
		finhelpStageTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &FinhelpProcService{
		db:             db,
		finhelpProcTN:  finhelpProcTN,
		userTN:         userTN,
		studentTN:      studentTN,
		finhelpCtgTN:   finhelpCtgTN,
		finhelpStageTN: finhelpStageTN,
	}
}

func (fprs *FinhelpProcService) InsertFinhelpProc(data *finhelpprocschema.FinhelpProcInsert) (finhelpProcDB *finhelpprocschema.FinhelpProcDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[2]sID,
		%[3]sID,
		%[4]sID,
		%[5]sID,
	)
	VALUES (
		%[6]d,
		%[7]d,
		%[8]d,
		%[9]d
	)
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[4]sID,
	%[5]sID,
	%[1]sCreatedAt`,
		fprs.finhelpProcTN,
		fprs.userTN,
		fprs.studentTN,
		fprs.finhelpCtgTN,
		fprs.finhelpStageTN,
		data.UserID,
		data.StudentID,
		data.FinhelpCtgID,
		data.FinhelpStageID,
	)

	stmt, err := fprs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	finhelpProcDB = &finhelpprocschema.FinhelpProcDB{}
	err = stmt.QueryRow().Scan(
		&finhelpProcDB.FinhelpProcID,
		&finhelpProcDB.UserID,
		&finhelpProcDB.StudentID,
		&finhelpProcDB.FinhelpCtgID,
		&finhelpProcDB.FinhelpStageID,
		&finhelpProcDB.FinhelpProcCreatedAt,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return finhelpProcDB, nil
}

func (fprs *FinhelpProcService) UpdateFinhelpProc(data *finhelpprocschema.FinhelpProcUpdate) (finhelpProcDB *finhelpprocschema.FinhelpProcDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[2]sID = %[7]d,
	%[3]sID = %[8]d,
	%[4]sID = %[9]d,
	%[5]sID = %[10]d
WHERE %[1]sID = %[6]d
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[4]sID,
	%[5]sID,
	%[1]sCreatedAt`,
		fprs.finhelpProcTN,
		fprs.userTN,
		fprs.studentTN,
		fprs.finhelpCtgTN,
		fprs.finhelpStageTN,
		data.FinhelpProcID,
		data.UserID,
		data.StudentID,
		data.FinhelpCtgID,
		data.FinhelpStageID,
	)

	stmt, err := fprs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	finhelpProcDB = &finhelpprocschema.FinhelpProcDB{}
	err = stmt.QueryRow().Scan(
		&finhelpProcDB.FinhelpProcID,
		&finhelpProcDB.UserID,
		&finhelpProcDB.StudentID,
		&finhelpProcDB.FinhelpCtgID,
		&finhelpProcDB.FinhelpStageID,
		&finhelpProcDB.FinhelpProcCreatedAt,
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

	return finhelpProcDB, nil
}

// Toggle IsHidden
func (fprs *FinhelpProcService) DeleteFinhelpProc(data *finhelpprocschema.FinhelpProcDelete) (finhelpProcDB *finhelpprocschema.FinhelpProcDB, err error) {
	query := fmt.Sprintf(`
DELETE FROM %[1]s WHERE %[1]sID = %[6]d
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[4]sID,
	%[5]sID,
	%[1]sCreatedAt`,
		fprs.finhelpProcTN,
		fprs.userTN,
		fprs.studentTN,
		fprs.finhelpCtgTN,
		fprs.finhelpStageTN,
		data.FinhelpProcID,
	)

	stmt, err := fprs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	finhelpProcDB = &finhelpprocschema.FinhelpProcDB{}
	err = stmt.QueryRow().Scan(
		&finhelpProcDB.FinhelpProcID,
		&finhelpProcDB.UserID,
		&finhelpProcDB.StudentID,
		&finhelpProcDB.FinhelpCtgID,
		&finhelpProcDB.FinhelpStageID,
		&finhelpProcDB.FinhelpProcCreatedAt,
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

	return finhelpProcDB, nil
}

func (fprs *FinhelpProcService) GetFinhelpProc(data *finhelpprocschema.FinhelpProcGet) (finhelpProcDB *finhelpprocschema.FinhelpProcDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[4]sID,
	%[5]sID,
	%[1]sCreatedAt
FROM %[1]s
WHERE %[1]sID = %[6]d`,
		fprs.finhelpProcTN,
		fprs.userTN,
		fprs.studentTN,
		fprs.finhelpCtgTN,
		fprs.finhelpStageTN,
		data.FinhelpProcID,
	)

	stmt, err := fprs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	finhelpProcDB = &finhelpprocschema.FinhelpProcDB{}
	err = stmt.QueryRow().Scan(
		&finhelpProcDB.FinhelpProcID,
		&finhelpProcDB.UserID,
		&finhelpProcDB.StudentID,
		&finhelpProcDB.FinhelpCtgID,
		&finhelpProcDB.FinhelpStageID,
		&finhelpProcDB.FinhelpProcCreatedAt,
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

	return finhelpProcDB, nil
}

func (fprs *FinhelpProcService) GetFinhelpProcs(data *finhelpprocschema.FinhelpProcsGet) (finhelpProcsDB []*finhelpprocschema.FinhelpProcDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[4]sID,
	%[5]sID,
	%[1]sCreatedAt
FROM %[1]s
ORDER BY %[1]sID DESC`,
		fprs.finhelpProcTN,
		fprs.userTN,
		fprs.studentTN,
		fprs.finhelpCtgTN,
		fprs.finhelpStageTN,
	)

	stmt, err := fprs.db.Prepare(query)
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

	finhelpProcsDB = []*finhelpprocschema.FinhelpProcDB{}
	finhelpProcDB := &finhelpprocschema.FinhelpProcDB{}
	for rows.Next() {
		err = rows.Scan(
			&finhelpProcDB.FinhelpProcID,
			&finhelpProcDB.UserID,
			&finhelpProcDB.StudentID,
			&finhelpProcDB.FinhelpCtgID,
			&finhelpProcDB.FinhelpStageID,
			&finhelpProcDB.FinhelpProcCreatedAt,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		finhelpProcsDB = append(finhelpProcsDB, finhelpProcDB)
		finhelpProcDB = &finhelpprocschema.FinhelpProcDB{}
	}

	return finhelpProcsDB, nil
}
