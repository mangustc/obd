package perfservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/perfschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type PerfService struct {
	db        *sql.DB
	perfTN    string
	courseTN  string
	studentTN string
}

func NewPerfService(db *sql.DB,
	perfTN string,
	courseTN string,
	studentTN string,
) (ps *PerfService) {
	if db == nil ||
		perfTN == "" ||
		courseTN == "" ||
		studentTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &PerfService{
		db:        db,
		perfTN:    perfTN,
		courseTN:  courseTN,
		studentTN: studentTN,
	}
}

func (ps *PerfService) InsertPerf(data *perfschema.PerfInsert) (perfDB *perfschema.PerfDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[2]sID,
		%[3]sID,
		%[1]sGrade
	)
	VALUES (
		%[4]d,
		%[5]d,
		%[6]d
	)
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[1]sGrade`,
		ps.perfTN,
		ps.courseTN,
		ps.studentTN,
		data.CourseID,
		data.StudentID,
		data.PerfGrade,
	)

	stmt, err := ps.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	perfDB = &perfschema.PerfDB{}
	err = stmt.QueryRow().Scan(
		&perfDB.PerfID,
		&perfDB.CourseID,
		&perfDB.StudentID,
		&perfDB.PerfGrade,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return perfDB, nil
}

func (ps *PerfService) UpdatePerf(data *perfschema.PerfUpdate) (perfDB *perfschema.PerfDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
		%[2]sID = %[5]d,
		%[3]sID = %[6]d,
		%[1]sGrade = %[7]d
WHERE %[1]sID = %[4]d
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[1]sGrade`,
		ps.perfTN,
		ps.courseTN,
		ps.studentTN,
		data.PerfID,
		data.CourseID,
		data.StudentID,
		data.PerfGrade,
	)

	stmt, err := ps.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	perfDB = &perfschema.PerfDB{}
	err = stmt.QueryRow().Scan(
		&perfDB.PerfID,
		&perfDB.CourseID,
		&perfDB.StudentID,
		&perfDB.PerfGrade,
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

	return perfDB, nil
}

// Toggle IsHidden
func (ps *PerfService) DeletePerf(data *perfschema.PerfDelete) (perfDB *perfschema.PerfDB, err error) {
	query := fmt.Sprintf(`
DELETE FROM %[1]s WHERE %[1]sID = %[4]d
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[1]sGrade`,
		ps.perfTN,
		ps.courseTN,
		ps.studentTN,
		data.PerfID,
	)

	stmt, err := ps.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	perfDB = &perfschema.PerfDB{}
	err = stmt.QueryRow().Scan(
		&perfDB.PerfID,
		&perfDB.CourseID,
		&perfDB.StudentID,
		&perfDB.PerfGrade,
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

	return perfDB, nil
}

func (ps *PerfService) GetPerf(data *perfschema.PerfGet) (perfDB *perfschema.PerfDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[1]sGrade
FROM %[1]s
WHERE %[1]sID = %[4]d`,
		ps.perfTN,
		ps.courseTN,
		ps.studentTN,
		data.PerfID,
	)

	stmt, err := ps.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	perfDB = &perfschema.PerfDB{}
	err = stmt.QueryRow().Scan(
		&perfDB.PerfID,
		&perfDB.CourseID,
		&perfDB.StudentID,
		&perfDB.PerfGrade,
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

	return perfDB, nil
}

func (ps *PerfService) GetPerfs(data *perfschema.PerfsGet) (perfsDB []*perfschema.PerfDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[1]sGrade
FROM %[1]s
ORDER BY %[1]sID DESC`,
		ps.perfTN,
		ps.courseTN,
		ps.studentTN,
	)

	stmt, err := ps.db.Prepare(query)
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

	perfsDB = []*perfschema.PerfDB{}
	perfDB := &perfschema.PerfDB{}
	for rows.Next() {
		err = rows.Scan(
			&perfDB.PerfID,
			&perfDB.CourseID,
			&perfDB.StudentID,
			&perfDB.PerfGrade,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		perfsDB = append(perfsDB, perfDB)
		perfDB = &perfschema.PerfDB{}
	}

	return perfsDB, nil
}
