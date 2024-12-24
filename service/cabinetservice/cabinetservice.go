package cabinetservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/cabinetschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type CabinetService struct {
	db            *sql.DB
	cabinetTN     string
	buildingTN    string
	cabinetTypeTN string
}

func NewCabinetService(db *sql.DB,
	cabinetTN string,
	buildingTN string,
	cabinetTypeTN string,
) (cs *CabinetService) {
	if db == nil ||
		cabinetTN == "" ||
		buildingTN == "" ||
		cabinetTypeTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &CabinetService{
		db:            db,
		cabinetTN:     cabinetTN,
		buildingTN:    buildingTN,
		cabinetTypeTN: cabinetTypeTN,
	}
}

func (cs *CabinetService) InsertCabinet(data *cabinetschema.CabinetInsert) (cabinetDB *cabinetschema.CabinetDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[2]sID,
		%[3]sID,
		%[1]sNumber
	)
	VALUES (
		%[4]d,
		%[5]d,
		"%[6]s"
	)
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[1]sNumber,
	%[1]sIsHidden`,
		cs.cabinetTN,
		cs.buildingTN,
		cs.cabinetTypeTN,
		data.BuildingID,
		data.CabinetTypeID,
		data.CabinetNumber,
	)

	stmt, err := cs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	cabinetDB = &cabinetschema.CabinetDB{}
	err = stmt.QueryRow().Scan(
		&cabinetDB.CabinetID,
		&cabinetDB.BuildingID,
		&cabinetDB.CabinetTypeID,
		&cabinetDB.CabinetNumber,
		&cabinetDB.CabinetIsHidden,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return cabinetDB, nil
}

func (cs *CabinetService) UpdateCabinet(data *cabinetschema.CabinetUpdate) (cabinetDB *cabinetschema.CabinetDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[2]sID = %[5]d,
	%[3]sID = %[6]d,
	%[1]sNumber = "%[7]s"
WHERE %[1]sID = %[4]d
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[1]sNumber,
	%[1]sIsHidden`,
		cs.cabinetTN,
		cs.buildingTN,
		cs.cabinetTypeTN,
		data.CabinetID,
		data.BuildingID,
		data.CabinetTypeID,
		data.CabinetNumber,
	)

	stmt, err := cs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	cabinetDB = &cabinetschema.CabinetDB{}
	err = stmt.QueryRow().Scan(
		&cabinetDB.CabinetID,
		&cabinetDB.BuildingID,
		&cabinetDB.CabinetTypeID,
		&cabinetDB.CabinetNumber,
		&cabinetDB.CabinetIsHidden,
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

	return cabinetDB, nil
}

// Toggle IsHidden
func (cs *CabinetService) DeleteCabinet(data *cabinetschema.CabinetDelete) (cabinetDB *cabinetschema.CabinetDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s
	SET %[1]sIsHidden = 1 - %[1]sIsHidden
	WHERE %[1]sID = %[4]d
RETURNING
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[1]sNumber,
	%[1]sIsHidden`,
		cs.cabinetTN,
		cs.buildingTN,
		cs.cabinetTypeTN,
		data.CabinetID,
	)

	stmt, err := cs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	cabinetDB = &cabinetschema.CabinetDB{}
	err = stmt.QueryRow().Scan(
		&cabinetDB.CabinetID,
		&cabinetDB.BuildingID,
		&cabinetDB.CabinetTypeID,
		&cabinetDB.CabinetNumber,
		&cabinetDB.CabinetIsHidden,
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

	return cabinetDB, nil
}

func (cs *CabinetService) GetCabinet(data *cabinetschema.CabinetGet) (cabinetDB *cabinetschema.CabinetDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[1]sNumber,
	%[1]sIsHidden
FROM %[1]s
WHERE %[1]sID = %[4]d`,
		cs.cabinetTN,
		cs.buildingTN,
		cs.cabinetTypeTN,
		data.CabinetID,
	)

	stmt, err := cs.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	cabinetDB = &cabinetschema.CabinetDB{}
	err = stmt.QueryRow().Scan(
		&cabinetDB.CabinetID,
		&cabinetDB.BuildingID,
		&cabinetDB.CabinetTypeID,
		&cabinetDB.CabinetNumber,
		&cabinetDB.CabinetIsHidden,
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

	return cabinetDB, nil
}

func (cs *CabinetService) GetCabinets(data *cabinetschema.CabinetsGet) (cabinetsDB []*cabinetschema.CabinetDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[2]sID,
	%[3]sID,
	%[1]sNumber,
	%[1]sIsHidden
FROM %[1]s
ORDER BY %[1]sID DESC`,
		cs.cabinetTN,
		cs.buildingTN,
		cs.cabinetTypeTN,
	)

	stmt, err := cs.db.Prepare(query)
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

	cabinetsDB = []*cabinetschema.CabinetDB{}
	cabinetDB := &cabinetschema.CabinetDB{}
	for rows.Next() {
		err = rows.Scan(
			&cabinetDB.CabinetID,
			&cabinetDB.BuildingID,
			&cabinetDB.CabinetTypeID,
			&cabinetDB.CabinetNumber,
			&cabinetDB.CabinetIsHidden,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		cabinetsDB = append(cabinetsDB, cabinetDB)
		cabinetDB = &cabinetschema.CabinetDB{}
	}

	return cabinetsDB, nil
}
