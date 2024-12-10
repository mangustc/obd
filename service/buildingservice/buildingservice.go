package buildingservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/buildingschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type BuildingService struct {
	db         *sql.DB
	buildingTN string
}

func NewBuildingService(db *sql.DB, buildingTN string) (bs *BuildingService) {
	if db == nil || buildingTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &BuildingService{
		db:         db,
		buildingTN: buildingTN,
	}
}

func (bs *BuildingService) InsertBuilding(data *buildingschema.BuildingInsert) (buildingDB *buildingschema.BuildingDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[1]sName,
		%[1]sAddress
	)
	VALUES (
		"%[2]s",
		"%[3]s"
	)
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sAddress,
	%[1]sIsHidden`,
		bs.buildingTN,
		data.BuildingName,
		data.BuildingAddress,
	)

	stmt, err := bs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	buildingDB = &buildingschema.BuildingDB{}
	err = stmt.QueryRow().Scan(
		&buildingDB.BuildingID,
		&buildingDB.BuildingName,
		&buildingDB.BuildingAddress,
		&buildingDB.BuildingIsHidden,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return buildingDB, nil
}

func (bs *BuildingService) UpdateBuilding(data *buildingschema.BuildingUpdate) (buildingDB *buildingschema.BuildingDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[1]sName = "%[3]s",
	%[1]sAddress = "%[4]s"
WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sAddress,
	%[1]sIsHidden`,
		bs.buildingTN,
		data.BuildingID,
		data.BuildingName,
		data.BuildingAddress,
	)

	stmt, err := bs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	buildingDB = &buildingschema.BuildingDB{}
	err = stmt.QueryRow().Scan(
		&buildingDB.BuildingID,
		&buildingDB.BuildingName,
		&buildingDB.BuildingAddress,
		&buildingDB.BuildingIsHidden,
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

	return buildingDB, nil
}

// Toggle IsHidden
func (bs *BuildingService) DeleteBuilding(data *buildingschema.BuildingDelete) (buildingDB *buildingschema.BuildingDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s
	SET %[1]sIsHidden = 1 - %[1]sIsHidden
	WHERE %[1]sID = %[2]d
RETURNING
	%[1]sID,
	%[1]sName,
	%[1]sAddress,
	%[1]sIsHidden`,
		bs.buildingTN,
		data.BuildingID,
	)

	stmt, err := bs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	buildingDB = &buildingschema.BuildingDB{}
	err = stmt.QueryRow().Scan(
		&buildingDB.BuildingID,
		&buildingDB.BuildingName,
		&buildingDB.BuildingAddress,
		&buildingDB.BuildingIsHidden,
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

	return buildingDB, nil
}

func (bs *BuildingService) GetBuilding(data *buildingschema.BuildingGet) (buildingDB *buildingschema.BuildingDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sName,
	%[1]sAddress,
	%[1]sIsHidden
FROM %[1]s
WHERE %[1]sID = %[2]d`,
		bs.buildingTN,
		data.BuildingID,
	)

	stmt, err := bs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	buildingDB = &buildingschema.BuildingDB{}
	err = stmt.QueryRow().Scan(
		&buildingDB.BuildingID,
		&buildingDB.BuildingName,
		&buildingDB.BuildingAddress,
		&buildingDB.BuildingIsHidden,
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

	return buildingDB, nil
}

func (bs *BuildingService) GetBuildings(data *buildingschema.BuildingsGet) (buildingsDB []*buildingschema.BuildingDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sName,
	%[1]sAddress,
	%[1]sIsHidden
FROM %[1]s
ORDER BY %[1]sID DESC`,
		bs.buildingTN,
	)

	stmt, err := bs.db.Prepare(query)
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

	buildingsDB = []*buildingschema.BuildingDB{}
	buildingDB := &buildingschema.BuildingDB{}
	for rows.Next() {
		err = rows.Scan(
			&buildingDB.BuildingID,
			&buildingDB.BuildingName,
			&buildingDB.BuildingAddress,
			&buildingDB.BuildingIsHidden,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		buildingsDB = append(buildingsDB, buildingDB)
		buildingDB = &buildingschema.BuildingDB{}
	}

	return buildingsDB, nil
}
