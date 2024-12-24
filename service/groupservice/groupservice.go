package groupservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/groupschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type GroupService struct {
	db      *sql.DB
	groupTN string
}

func NewGroupService(db *sql.DB, groupTN string) (grs *GroupService) {
	if db == nil || groupTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &GroupService{
		db:      db,
		groupTN: groupTN,
	}
}

func (grs *GroupService) InsertGroup(data *groupschema.GroupInsert) (groupDB *groupschema.GroupDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[1]sYear,
		%[1]sNumber,
		%[1]sCourseName
	)
	VALUES (
		%[2]d,
		"%[3]s",
		"%[4]s"
	)
RETURNING
	%[1]sID,
	%[1]sYear,
	%[1]sNumber,
	%[1]sCourseName,
	%[1]sIsHidden`,
		grs.groupTN,
		data.GroupYear,
		data.GroupNumber,
		data.GroupCourseName,
	)

	stmt, err := grs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	groupDB = &groupschema.GroupDB{}
	err = stmt.QueryRow().Scan(
		&groupDB.GroupID,
		&groupDB.GroupYear,
		&groupDB.GroupNumber,
		&groupDB.GroupCourseName,
		&groupDB.GroupIsHidden,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return groupDB, nil
}

func (grs *GroupService) UpdateGroup(data *groupschema.GroupUpdate) (groupDB *groupschema.GroupDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[1]sYear = %[3]d,
	%[1]sNumber = "%[4]s",
	%[1]sCourseName = "%[5]s"
WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sYear,
	%[1]sNumber,
	%[1]sCourseName,
	%[1]sIsHidden`,
		grs.groupTN,
		data.GroupID,
		data.GroupYear,
		data.GroupNumber,
		data.GroupCourseName,
	)

	stmt, err := grs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	groupDB = &groupschema.GroupDB{}
	err = stmt.QueryRow().Scan(
		&groupDB.GroupID,
		&groupDB.GroupYear,
		&groupDB.GroupNumber,
		&groupDB.GroupCourseName,
		&groupDB.GroupIsHidden,
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

	return groupDB, nil
}

// Toggle IsHidden
func (grs *GroupService) DeleteGroup(data *groupschema.GroupDelete) (groupDB *groupschema.GroupDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s
	SET %[1]sIsHidden = 1 - %[1]sIsHidden
	WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sYear,
	%[1]sNumber,
	%[1]sCourseName,
	%[1]sIsHidden`,
		grs.groupTN,
		data.GroupID,
	)

	stmt, err := grs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	groupDB = &groupschema.GroupDB{}
	err = stmt.QueryRow().Scan(
		&groupDB.GroupID,
		&groupDB.GroupYear,
		&groupDB.GroupNumber,
		&groupDB.GroupCourseName,
		&groupDB.GroupIsHidden,
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

	return groupDB, nil
}

func (grs *GroupService) GetGroup(data *groupschema.GroupGet) (groupDB *groupschema.GroupDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sYear,
	%[1]sNumber,
	%[1]sCourseName,
	%[1]sIsHidden
FROM %[1]s
WHERE %[1]sID = %[2]d`,
		grs.groupTN,
		data.GroupID,
	)

	stmt, err := grs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	groupDB = &groupschema.GroupDB{}
	err = stmt.QueryRow().Scan(
		&groupDB.GroupID,
		&groupDB.GroupYear,
		&groupDB.GroupNumber,
		&groupDB.GroupCourseName,
		&groupDB.GroupIsHidden,
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

	return groupDB, nil
}

func (grs *GroupService) GetGroups(data *groupschema.GroupsGet) (groupsDB []*groupschema.GroupDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sYear,
	%[1]sNumber,
	%[1]sCourseName,
	%[1]sIsHidden
FROM %[1]s
ORDER BY %[1]sID DESC`,
		grs.groupTN,
	)

	stmt, err := grs.db.Prepare(query)
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

	groupsDB = []*groupschema.GroupDB{}
	groupDB := &groupschema.GroupDB{}
	for rows.Next() {
		err = rows.Scan(
			&groupDB.GroupID,
			&groupDB.GroupYear,
			&groupDB.GroupNumber,
			&groupDB.GroupCourseName,
			&groupDB.GroupIsHidden,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		groupsDB = append(groupsDB, groupDB)
		groupDB = &groupschema.GroupDB{}
	}

	return groupsDB, nil
}
