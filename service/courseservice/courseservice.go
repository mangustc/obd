package courseservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/courseschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type CourseService struct {
	db           *sql.DB
	courseTN     string
	courseTypeTN string
}

func NewCourseService(db *sql.DB,
	courseTN string,
	courseTypeTN string,
) (cos *CourseService) {
	if db == nil ||
		courseTN == "" ||
		courseTypeTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &CourseService{
		db:           db,
		courseTN:     courseTN,
		courseTypeTN: courseTypeTN,
	}
}

func (cos *CourseService) InsertCourse(data *courseschema.CourseInsert) (courseDB *courseschema.CourseDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[2]sID,
		%[1]sName
	)
	VALUES (
		%[3]d,
		"%[4]s"
	)
RETURNING
	%[1]sID,
	%[2]sID,
	%[1]sName,
	%[1]sIsHidden`,
		cos.courseTN,
		cos.courseTypeTN,
		data.CourseTypeID,
		data.CourseName,
	)

	stmt, err := cos.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	courseDB = &courseschema.CourseDB{}
	err = stmt.QueryRow().Scan(
		&courseDB.CourseID,
		&courseDB.CourseTypeID,
		&courseDB.CourseName,
		&courseDB.CourseIsHidden,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return courseDB, nil
}

func (cos *CourseService) UpdateCourse(data *courseschema.CourseUpdate) (courseDB *courseschema.CourseDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[2]sID = %[4]d,
	%[1]sName = "%[5]s"
WHERE %[1]sID = %[3]d
RETURNING
	%[1]sID,
	%[2]sID,
	%[1]sName,
	%[1]sIsHidden`,
		cos.courseTN,
		cos.courseTypeTN,
		data.CourseID,
		data.CourseTypeID,
		data.CourseName,
	)

	stmt, err := cos.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	courseDB = &courseschema.CourseDB{}
	err = stmt.QueryRow().Scan(
		&courseDB.CourseID,
		&courseDB.CourseTypeID,
		&courseDB.CourseName,
		&courseDB.CourseIsHidden,
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

	return courseDB, nil
}

// Toggle IsHidden
func (cos *CourseService) DeleteCourse(data *courseschema.CourseDelete) (courseDB *courseschema.CourseDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s
	SET %[1]sIsHidden = 1 - %[1]sIsHidden
	WHERE %[1]sID = %[3]d
RETURNING
	%[1]sID,
	%[2]sID,
	%[1]sName,
	%[1]sIsHidden`,
		cos.courseTN,
		cos.courseTypeTN,
		data.CourseID,
	)

	stmt, err := cos.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	courseDB = &courseschema.CourseDB{}
	err = stmt.QueryRow().Scan(
		&courseDB.CourseID,
		&courseDB.CourseTypeID,
		&courseDB.CourseName,
		&courseDB.CourseIsHidden,
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

	return courseDB, nil
}

func (cos *CourseService) GetCourse(data *courseschema.CourseGet) (courseDB *courseschema.CourseDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[2]sID,
	%[1]sName,
	%[1]sIsHidden
FROM %[1]s
WHERE %[1]sID = %[3]d`,
		cos.courseTN,
		cos.courseTypeTN,
		data.CourseID,
	)

	stmt, err := cos.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	courseDB = &courseschema.CourseDB{}
	err = stmt.QueryRow().Scan(
		&courseDB.CourseID,
		&courseDB.CourseTypeID,
		&courseDB.CourseName,
		&courseDB.CourseIsHidden,
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

	return courseDB, nil
}

func (cos *CourseService) GetCourses(data *courseschema.CoursesGet) (coursesDB []*courseschema.CourseDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[2]sID,
	%[1]sName,
	%[1]sIsHidden
FROM %[1]s
ORDER BY %[1]sID DESC`,
		cos.courseTN,
		cos.courseTypeTN,
	)

	stmt, err := cos.db.Prepare(query)
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

	coursesDB = []*courseschema.CourseDB{}
	courseDB := &courseschema.CourseDB{}
	for rows.Next() {
		err = rows.Scan(
			&courseDB.CourseID,
			&courseDB.CourseTypeID,
			&courseDB.CourseName,
			&courseDB.CourseIsHidden,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		coursesDB = append(coursesDB, courseDB)
		courseDB = &courseschema.CourseDB{}
	}

	return coursesDB, nil
}
