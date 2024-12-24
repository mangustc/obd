package studentservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/studentschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type StudentService struct {
	db        *sql.DB
	studentTN string
	groupTN   string
}

func NewStudentService(db *sql.DB, studentTN string, groupTN string) (sts *StudentService) {
	if db == nil || studentTN == "" || groupTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &StudentService{
		db:        db,
		studentTN: studentTN,
		groupTN:   groupTN,
	}
}

func (sts *StudentService) InsertStudent(data *studentschema.StudentInsert) (studentDB *studentschema.StudentDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[1]sLastname,
		%[1]sFirstname,
		%[1]sMiddlename,
		%[1]sPhoneNumber,
		%[2]sID
	)
	VALUES (
		"%[3]s",
		"%[4]s",
		"%[5]s",
		"%[6]s",
		"%[7]d"
	)
RETURNING
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPhoneNumber,
	%[1]sIsHidden,
	%[2]sID`,
		sts.studentTN,
		sts.groupTN,
		data.StudentLastname,
		data.StudentFirstname,
		data.StudentMiddlename,
		data.StudentPhoneNumber,
		data.GroupID,
	)

	stmt, err := sts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	studentDB = &studentschema.StudentDB{}
	err = stmt.QueryRow().Scan(
		&studentDB.StudentID,
		&studentDB.StudentLastname,
		&studentDB.StudentFirstname,
		&studentDB.StudentMiddlename,
		&studentDB.StudentPhoneNumber,
		&studentDB.StudentIsHidden,
		&studentDB.GroupID,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return studentDB, nil
}

func (sts *StudentService) UpdateStudent(data *studentschema.StudentUpdate) (studentDB *studentschema.StudentDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[1]sLastname = "%[4]s",
	%[1]sFirstname = "%[5]s",
	%[1]sMiddlename = "%[6]s",
	%[1]sPhoneNumber = "%[7]s",
	%[2]sID = "%[8]d"
WHERE %[1]sID = %[3]d
RETURNING
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPhoneNumber,
	%[1]sIsHidden,
	%[2]sID`,
		sts.studentTN,
		sts.groupTN,
		data.StudentID,
		data.StudentLastname,
		data.StudentFirstname,
		data.StudentMiddlename,
		data.StudentPhoneNumber,
		data.GroupID,
	)

	stmt, err := sts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	studentDB = &studentschema.StudentDB{}
	err = stmt.QueryRow().Scan(
		&studentDB.StudentID,
		&studentDB.StudentLastname,
		&studentDB.StudentFirstname,
		&studentDB.StudentMiddlename,
		&studentDB.StudentPhoneNumber,
		&studentDB.StudentIsHidden,
		&studentDB.GroupID,
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

	return studentDB, nil
}

// Toggle IsHidden
func (sts *StudentService) DeleteStudent(data *studentschema.StudentDelete) (studentDB *studentschema.StudentDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s
	SET %[1]sIsHidden = 1 - %[1]sIsHidden
	WHERE %[1]sID = %[3]d
RETURNING
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPhoneNumber,
	%[1]sIsHidden,
	%[2]sID`,
		sts.studentTN,
		sts.groupTN,
		data.StudentID,
	)

	stmt, err := sts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	studentDB = &studentschema.StudentDB{}
	err = stmt.QueryRow().Scan(
		&studentDB.StudentID,
		&studentDB.StudentLastname,
		&studentDB.StudentFirstname,
		&studentDB.StudentMiddlename,
		&studentDB.StudentPhoneNumber,
		&studentDB.StudentIsHidden,
		&studentDB.GroupID,
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

	return studentDB, nil
}

func (sts *StudentService) GetStudent(data *studentschema.StudentGet) (studentDB *studentschema.StudentDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPhoneNumber,
	%[1]sIsHidden,
	%[2]sID
FROM %[1]s
WHERE %[1]sID = %[3]d`,
		sts.studentTN,
		sts.groupTN,
		data.StudentID,
	)

	stmt, err := sts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	studentDB = &studentschema.StudentDB{}
	err = stmt.QueryRow().Scan(
		&studentDB.StudentID,
		&studentDB.StudentLastname,
		&studentDB.StudentFirstname,
		&studentDB.StudentMiddlename,
		&studentDB.StudentPhoneNumber,
		&studentDB.StudentIsHidden,
		&studentDB.GroupID,
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

	return studentDB, nil
}

func (sts *StudentService) GetStudents(data *studentschema.StudentsGet) (studentsDB []*studentschema.StudentDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPhoneNumber,
	%[1]sIsHidden,
	%[2]sID
FROM %[1]s
ORDER BY %[1]sID DESC`,
		sts.studentTN,
		sts.groupTN,
	)

	stmt, err := sts.db.Prepare(query)
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

	studentsDB = []*studentschema.StudentDB{}
	studentDB := &studentschema.StudentDB{}
	for rows.Next() {
		err = rows.Scan(
			&studentDB.StudentID,
			&studentDB.StudentLastname,
			&studentDB.StudentFirstname,
			&studentDB.StudentMiddlename,
			&studentDB.StudentPhoneNumber,
			&studentDB.StudentIsHidden,
			&studentDB.GroupID,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		studentsDB = append(studentsDB, studentDB)
		studentDB = &studentschema.StudentDB{}
	}

	return studentsDB, nil
}
