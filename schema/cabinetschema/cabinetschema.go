package cabinetschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type CabinetDB struct {
	CabinetID       int    `json:"CabinetID" server:"y"`
	CabinetTypeID   int    `json:"CabinetTypeID" server:"y"`
	BuildingID      int    `json:"BuildingID" server:"y"`
	CabinetNumber   string `json:"CabinetNumber"`
	CabinetIsHidden bool   `json:"CabinetIsHidden"`
}

type CabinetInsert struct {
	CabinetTypeID int    `json:"CabinetTypeID" server:"y"`
	BuildingID    int    `json:"BuildingID" server:"y"`
	CabinetNumber string `json:"CabinetNumber"`
}

type CabinetUpdate struct {
	CabinetID     int    `json:"CabinetID" server:"y"`
	CabinetTypeID int    `json:"CabinetTypeID" server:"y"`
	BuildingID    int    `json:"BuildingID" server:"y"`
	CabinetNumber string `json:"CabinetNumber"`
}

type CabinetDelete struct {
	CabinetID int `json:"CabinetID" server:"y"`
}

type CabinetGet struct {
	CabinetID int `json:"CabinetID" server:"y"`
}

type CabinetsGet struct{}

func GetCabinetInputOptionsFromCabinetsDB(cabinetsDB []*CabinetDB) []*schema.InputOption {
	notHiddenCabinetsDB := GetNotHiddenCabinetsDB(cabinetsDB)
	inputOptions := []*schema.InputOption{}
	for _, cabinetDB := range notHiddenCabinetsDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%s %d", cabinetDB.CabinetNumber, cabinetDB.BuildingID),
			InputOptionValue: fmt.Sprintf("%d", cabinetDB.CabinetID),
		})
	}
	return inputOptions
}

func GetNotHiddenCabinetsDB(cabinetsDB []*CabinetDB) []*CabinetDB {
	notHiddenCabinetsDB := []*CabinetDB{}
	for _, cabinetDB := range cabinetsDB {
		if !cabinetDB.CabinetIsHidden {
			notHiddenCabinetsDB = append(notHiddenCabinetsDB, cabinetDB)
		}
	}
	return notHiddenCabinetsDB
}

func ValidateCabinetDB(cabinetDB *CabinetDB) (err error) {
	if cabinetDB == nil {
		panic("Object is nil")
	}
	if cabinetDB.CabinetID <= 0 || cabinetDB.CabinetTypeID <= 0 || cabinetDB.BuildingID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCabinetInsert(cabinetInsert *CabinetInsert) (err error) {
	if cabinetInsert == nil {
		panic("Object is nil")
	}
	if cabinetInsert.CabinetTypeID <= 0 || cabinetInsert.BuildingID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCabinetUpdate(cabinetUpdate *CabinetUpdate) (err error) {
	if cabinetUpdate == nil {
		panic("Object is nil")
	}
	if cabinetUpdate.CabinetID <= 0 || cabinetUpdate.CabinetTypeID <= 0 || cabinetUpdate.BuildingID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCabinetDelete(cabinetDelete *CabinetDelete) (err error) {
	if cabinetDelete == nil {
		panic("Object is nil")
	}
	if cabinetDelete.CabinetID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCabinetGet(cabinetGet *CabinetGet) (err error) {
	if cabinetGet == nil {
		panic("Object is nil")
	}
	if cabinetGet.CabinetID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCabinetsGet(cabinetGets *CabinetsGet) (err error) {
	if cabinetGets == nil {
		panic("Object is nil")
	}
	return nil
}
