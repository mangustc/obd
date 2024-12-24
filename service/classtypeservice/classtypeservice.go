package classtypeservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/classtypeschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type ClassTypeService struct {
	db          *sql.DB
	classTypeTN string
}

func NewClassTypeService(db *sql.DB, classTypeTN string) (clts *ClassTypeService) {
	if db == nil || classTypeTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &ClassTypeService{
		db:          db,
		classTypeTN: classTypeTN,
	}
}

func (clts *ClassTypeService) InsertClassType(data *classtypeschema.ClassTypeInsert) (classTypeDB *classtypeschema.ClassTypeDB, err error) {
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
		clts.classTypeTN,
		data.ClassTypeName,
	)

	stmt, err := clts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	classTypeDB = &classtypeschema.ClassTypeDB{}
	err = stmt.QueryRow().Scan(
		&classTypeDB.ClassTypeID,
		&classTypeDB.ClassTypeName,
		&classTypeDB.ClassTypeIsHidden,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return classTypeDB, nil
}

func (clts *ClassTypeService) UpdateClassType(data *classtypeschema.ClassTypeUpdate) (classTypeDB *classtypeschema.ClassTypeDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[1]sName = "%[3]s"
WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden`,
		clts.classTypeTN,
		data.ClassTypeID,
		data.ClassTypeName,
	)

	stmt, err := clts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	classTypeDB = &classtypeschema.ClassTypeDB{}
	err = stmt.QueryRow().Scan(
		&classTypeDB.ClassTypeID,
		&classTypeDB.ClassTypeName,
		&classTypeDB.ClassTypeIsHidden,
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

	return classTypeDB, nil
}

// Toggle IsHidden
func (clts *ClassTypeService) DeleteClassType(data *classtypeschema.ClassTypeDelete) (classTypeDB *classtypeschema.ClassTypeDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s
	SET %[1]sIsHidden = 1 - %[1]sIsHidden
	WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden`,
		clts.classTypeTN,
		data.ClassTypeID,
	)

	stmt, err := clts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	classTypeDB = &classtypeschema.ClassTypeDB{}
	err = stmt.QueryRow().Scan(
		&classTypeDB.ClassTypeID,
		&classTypeDB.ClassTypeName,
		&classTypeDB.ClassTypeIsHidden,
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

	return classTypeDB, nil
}

func (clts *ClassTypeService) GetClassType(data *classtypeschema.ClassTypeGet) (classTypeDB *classtypeschema.ClassTypeDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden
FROM %[1]s
WHERE %[1]sID = %[2]d`,
		clts.classTypeTN,
		data.ClassTypeID,
	)

	stmt, err := clts.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	classTypeDB = &classtypeschema.ClassTypeDB{}
	err = stmt.QueryRow().Scan(
		&classTypeDB.ClassTypeID,
		&classTypeDB.ClassTypeName,
		&classTypeDB.ClassTypeIsHidden,
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

	return classTypeDB, nil
}

func (clts *ClassTypeService) GetClassTypes(data *classtypeschema.ClassTypesGet) (classTypesDB []*classtypeschema.ClassTypeDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sName,
	%[1]sIsHidden
FROM %[1]s
ORDER BY %[1]sID DESC`,
		clts.classTypeTN,
	)

	stmt, err := clts.db.Prepare(query)
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

	classTypesDB = []*classtypeschema.ClassTypeDB{}
	classTypeDB := &classtypeschema.ClassTypeDB{}
	for rows.Next() {
		err = rows.Scan(
			&classTypeDB.ClassTypeID,
			&classTypeDB.ClassTypeName,
			&classTypeDB.ClassTypeIsHidden,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		classTypesDB = append(classTypesDB, classTypeDB)
		classTypeDB = &classtypeschema.ClassTypeDB{}
	}

	return classTypesDB, nil
}
