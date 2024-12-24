package coursetypeservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/coursetypeschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type CourseTypeService struct {
	db           *sql.DB
	courseTypeTN string
}

func NewCourseTypeService(db *sql.DB, courseTypeTN string) (cots *CourseTypeService) {
	if db == nil || courseTypeTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &CourseTypeService{
		db:           db,
		courseTypeTN: courseTypeTN,
	}
}

func (cots *CourseTypeService) InsertCourseType(data *coursetypeschema.CourseTypeInsert) (courseTypeDB *coursetypeschema.CourseTypeDB, err error) {
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
		cots.courseTypeTN,
		data.CourseTypeName,
	)

	stmt, err := cots.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	courseTypeDB = &coursetypeschema.CourseTypeDB{}
	err = stmt.QueryRow().Scan(
		&courseTypeDB.CourseTypeID,
		&courseTypeDB.CourseTypeName,
		&courseTypeDB.CourseTypeIsHidden,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return courseTypeDB, nil
}

func (cots *CourseTypeService) UpdateCourseType(data *coursetypeschema.CourseTypeUpdate) (courseTypeDB *coursetypeschema.CourseTypeDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[1]sName = "%[3]s"
WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden`,
		cots.courseTypeTN,
		data.CourseTypeID,
		data.CourseTypeName,
	)

	stmt, err := cots.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	courseTypeDB = &coursetypeschema.CourseTypeDB{}
	err = stmt.QueryRow().Scan(
		&courseTypeDB.CourseTypeID,
		&courseTypeDB.CourseTypeName,
		&courseTypeDB.CourseTypeIsHidden,
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

	return courseTypeDB, nil
}

// Toggle IsHidden
func (cots *CourseTypeService) DeleteCourseType(data *coursetypeschema.CourseTypeDelete) (courseTypeDB *coursetypeschema.CourseTypeDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s
	SET %[1]sIsHidden = 1 - %[1]sIsHidden
	WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden`,
		cots.courseTypeTN,
		data.CourseTypeID,
	)

	stmt, err := cots.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	courseTypeDB = &coursetypeschema.CourseTypeDB{}
	err = stmt.QueryRow().Scan(
		&courseTypeDB.CourseTypeID,
		&courseTypeDB.CourseTypeName,
		&courseTypeDB.CourseTypeIsHidden,
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

	return courseTypeDB, nil
}

func (cots *CourseTypeService) GetCourseType(data *coursetypeschema.CourseTypeGet) (courseTypeDB *coursetypeschema.CourseTypeDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden
FROM %[1]s
WHERE %[1]sID = %[2]d`,
		cots.courseTypeTN,
		data.CourseTypeID,
	)

	stmt, err := cots.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	courseTypeDB = &coursetypeschema.CourseTypeDB{}
	err = stmt.QueryRow().Scan(
		&courseTypeDB.CourseTypeID,
		&courseTypeDB.CourseTypeName,
		&courseTypeDB.CourseTypeIsHidden,
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

	return courseTypeDB, nil
}

func (cots *CourseTypeService) GetCourseTypes(data *coursetypeschema.CourseTypesGet) (courseTypesDB []*coursetypeschema.CourseTypeDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden
FROM %[1]s
ORDER BY %[1]sID DESC`,
		cots.courseTypeTN,
	)

	stmt, err := cots.db.Prepare(query)
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

	courseTypesDB = []*coursetypeschema.CourseTypeDB{}
	courseTypeDB := &coursetypeschema.CourseTypeDB{}
	for rows.Next() {
		err = rows.Scan(
			&courseTypeDB.CourseTypeID,
			&courseTypeDB.CourseTypeName,
			&courseTypeDB.CourseTypeIsHidden,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		courseTypesDB = append(courseTypesDB, courseTypeDB)
		courseTypeDB = &coursetypeschema.CourseTypeDB{}
	}

	return courseTypesDB, nil
}
