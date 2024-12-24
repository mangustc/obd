package cabinettypeschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type CabinetTypeDB struct {
	CabinetTypeID       int    `json:"CabinetTypeID" server:"y"`
	CabinetTypeName     string `json:"CabinetTypeName"`
	CabinetTypeIsHidden bool   `json:"CabinetTypeIsHidden"`
}

type CabinetTypeInsert struct {
	CabinetTypeName string `json:"CabinetTypeName"`
}

type CabinetTypeUpdate struct {
	CabinetTypeID   int    `json:"CabinetTypeID" server:"y"`
	CabinetTypeName string `json:"CabinetTypeName"`
}

type CabinetTypeDelete struct {
	CabinetTypeID int `json:"CabinetTypeID" server:"y"`
}

type CabinetTypeGet struct {
	CabinetTypeID int `json:"CabinetTypeID" server:"y"`
}

type CabinetTypesGet struct{}

func GetCabinetTypeInputOptionsFromCabinetTypesDB(cabinetTypesDB []*CabinetTypeDB) []*schema.InputOption {
	notHiddenCabinetTypesDB := GetNotHiddenCabinetTypesDB(cabinetTypesDB)
	inputOptions := []*schema.InputOption{}
	for _, cabinetTypeDB := range notHiddenCabinetTypesDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%s", cabinetTypeDB.CabinetTypeName),
			InputOptionValue: fmt.Sprintf("%d", cabinetTypeDB.CabinetTypeID),
		})
	}
	return inputOptions
}

func GetNotHiddenCabinetTypesDB(cabinetTypesDB []*CabinetTypeDB) []*CabinetTypeDB {
	notHiddenCabinetTypesDB := []*CabinetTypeDB{}
	for _, cabinetTypeDB := range cabinetTypesDB {
		if !cabinetTypeDB.CabinetTypeIsHidden {
			notHiddenCabinetTypesDB = append(notHiddenCabinetTypesDB, cabinetTypeDB)
		}
	}
	return notHiddenCabinetTypesDB
}

func ValidateCabinetTypeDB(cabinetTypeDB *CabinetTypeDB) (err error) {
	if cabinetTypeDB == nil {
		panic("Object is nil")
	}
	if cabinetTypeDB.CabinetTypeID <= 0 || cabinetTypeDB.CabinetTypeName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCabinetTypeInsert(cabinetTypeInsert *CabinetTypeInsert) (err error) {
	if cabinetTypeInsert == nil {
		panic("Object is nil")
	}
	if cabinetTypeInsert.CabinetTypeName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCabinetTypeUpdate(cabinetTypeUpdate *CabinetTypeUpdate) (err error) {
	if cabinetTypeUpdate == nil {
		panic("Object is nil")
	}
	if cabinetTypeUpdate.CabinetTypeID <= 0 || cabinetTypeUpdate.CabinetTypeName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCabinetTypeDelete(cabinetTypeDelete *CabinetTypeDelete) (err error) {
	if cabinetTypeDelete == nil {
		panic("Object is nil")
	}
	if cabinetTypeDelete.CabinetTypeID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCabinetTypeGet(cabinetTypeGet *CabinetTypeGet) (err error) {
	if cabinetTypeGet == nil {
		panic("Object is nil")
	}
	if cabinetTypeGet.CabinetTypeID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCabinetTypesGet(cabinetTypeGets *CabinetTypesGet) (err error) {
	if cabinetTypeGets == nil {
		panic("Object is nil")
	}
	return nil
}
