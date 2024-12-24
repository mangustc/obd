package classservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/classschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type ClassService struct {
	db          *sql.DB
	classTN     string
	classTypeTN string
	profTN      string
	cabinetTN   string
	courseTN    string
	groupTN     string
}

func NewClassService(db *sql.DB,
	classTN string,
	classTypeTN string,
	profTN string,
	cabinetTN string,
	courseTN string,
	groupTN string,
) (cls *ClassService) {
	if db == nil ||
		classTN == "" ||
		classTypeTN == "" ||
		profTN == "" ||
		cabinetTN == "" ||
		courseTN == "" ||
		groupTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &ClassService{
		db:          db,
		classTN:     classTN,
		classTypeTN: classTypeTN,
		profTN:      profTN,
		cabinetTN:   cabinetTN,
		courseTN:    courseTN,
		groupTN:     groupTN,
	}
}

func (cls *ClassService) InsertClass(data *classschema.ClassInsert) (classDB *classschema.ClassDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[2]sID,
		%[3]sID,
		%[4]sID,
		%[5]sID,
		%[6]sID,
		%[1]sStart,
		%[1]sNumber
	)
	VALUES (
		%[7]d,
		%[8]d,
		%[9]d,
		%[10]d,
		%[11]d,
		"%[12]s",
		%[13]d
	)
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[4]sID,
	%[5]sID,
	%[6]sID,
	%[1]sStart,
	%[1]sNumber`,
		cls.classTN,
		cls.classTypeTN,
		cls.profTN,
		cls.cabinetTN,
		cls.courseTN,
		cls.groupTN,
		data.ClassTypeID,
		data.ProfID,
		data.CabinetID,
		data.CourseID,
		data.GroupID,
		data.ClassStart,
		data.ClassNumber,
	)

	stmt, err := cls.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	classDB = &classschema.ClassDB{}
	err = stmt.QueryRow().Scan(
		&classDB.ClassID,
		&classDB.ClassTypeID,
		&classDB.ProfID,
		&classDB.CabinetID,
		&classDB.CourseID,
		&classDB.GroupID,
		&classDB.ClassStart,
		&classDB.ClassNumber,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return classDB, nil
}

func (cls *ClassService) UpdateClass(data *classschema.ClassUpdate) (classDB *classschema.ClassDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
		%[2]sID = %[8]d,
		%[3]sID = %[9]d,
		%[4]sID = %[10]d,
		%[5]sID = %[11]d,
		%[6]sID = %[12]d,
		%[1]sStart = "%[13]s",
		%[1]sNumber = %[14]d
WHERE %[1]sID = %[7]d
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[4]sID,
	%[5]sID,
	%[6]sID,
	%[1]sStart,
	%[1]sNumber`,
		cls.classTN,
		cls.classTypeTN,
		cls.profTN,
		cls.cabinetTN,
		cls.courseTN,
		cls.groupTN,
		data.ClassID,
		data.ClassTypeID,
		data.ProfID,
		data.CabinetID,
		data.CourseID,
		data.GroupID,
		data.ClassStart,
		data.ClassNumber,
	)

	stmt, err := cls.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	classDB = &classschema.ClassDB{}
	err = stmt.QueryRow().Scan(
		&classDB.ClassID,
		&classDB.ClassTypeID,
		&classDB.ProfID,
		&classDB.CabinetID,
		&classDB.CourseID,
		&classDB.GroupID,
		&classDB.ClassStart,
		&classDB.ClassNumber,
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

	return classDB, nil
}

// Toggle IsHidden
func (cls *ClassService) DeleteClass(data *classschema.ClassDelete) (classDB *classschema.ClassDB, err error) {
	query := fmt.Sprintf(`
DELETE FROM %[1]s WHERE %[1]sID = %[7]d
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[4]sID,
	%[5]sID,
	%[6]sID,
	%[1]sStart,
	%[1]sNumber`,
		cls.classTN,
		cls.classTypeTN,
		cls.profTN,
		cls.cabinetTN,
		cls.courseTN,
		cls.groupTN,
		data.ClassID,
	)

	stmt, err := cls.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	classDB = &classschema.ClassDB{}
	err = stmt.QueryRow().Scan(
		&classDB.ClassID,
		&classDB.ClassTypeID,
		&classDB.ProfID,
		&classDB.CabinetID,
		&classDB.CourseID,
		&classDB.GroupID,
		&classDB.ClassStart,
		&classDB.ClassNumber,
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

	return classDB, nil
}

func (cls *ClassService) GetClass(data *classschema.ClassGet) (classDB *classschema.ClassDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[4]sID,
	%[5]sID,
	%[6]sID,
	%[1]sStart,
	%[1]sNumber
FROM %[1]s
WHERE %[1]sID = %[7]d`,
		cls.classTN,
		cls.classTypeTN,
		cls.profTN,
		cls.cabinetTN,
		cls.courseTN,
		cls.groupTN,
		data.ClassID,
	)

	stmt, err := cls.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	classDB = &classschema.ClassDB{}
	err = stmt.QueryRow().Scan(
		&classDB.ClassID,
		&classDB.ClassTypeID,
		&classDB.ProfID,
		&classDB.CabinetID,
		&classDB.CourseID,
		&classDB.GroupID,
		&classDB.ClassStart,
		&classDB.ClassNumber,
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

	return classDB, nil
}

func (cls *ClassService) GetClasss(data *classschema.ClasssGet) (classsDB []*classschema.ClassDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[4]sID,
	%[5]sID,
	%[6]sID,
	%[1]sStart,
	%[1]sNumber
FROM %[1]s
ORDER BY %[1]sID DESC`,
		cls.classTN,
		cls.classTypeTN,
		cls.profTN,
		cls.cabinetTN,
		cls.courseTN,
		cls.groupTN,
	)

	stmt, err := cls.db.Prepare(query)
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

	classsDB = []*classschema.ClassDB{}
	classDB := &classschema.ClassDB{}
	for rows.Next() {
		err = rows.Scan(
			&classDB.ClassID,
			&classDB.ClassTypeID,
			&classDB.ProfID,
			&classDB.CabinetID,
			&classDB.CourseID,
			&classDB.GroupID,
			&classDB.ClassStart,
			&classDB.ClassNumber,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		classsDB = append(classsDB, classDB)
		classDB = &classschema.ClassDB{}
	}

	return classsDB, nil
}
