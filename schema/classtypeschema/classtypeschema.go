package classtypeschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type ClassTypeDB struct {
	ClassTypeID       int    `json:"ClassTypeID" server:"y"`
	ClassTypeName     string `json:"ClassTypeName"`
	ClassTypeIsHidden bool   `json:"ClassTypeIsHidden"`
}

type ClassTypeInsert struct {
	ClassTypeName string `json:"ClassTypeName"`
}

type ClassTypeUpdate struct {
	ClassTypeID   int    `json:"ClassTypeID" server:"y"`
	ClassTypeName string `json:"ClassTypeName"`
}

type ClassTypeDelete struct {
	ClassTypeID int `json:"ClassTypeID" server:"y"`
}

type ClassTypeGet struct {
	ClassTypeID int `json:"ClassTypeID" server:"y"`
}

type ClassTypesGet struct{}

func GetClassTypeInputOptionsFromClassTypesDB(classTypesDB []*ClassTypeDB) []*schema.InputOption {
	notHiddenClassTypesDB := GetNotHiddenClassTypesDB(classTypesDB)
	inputOptions := []*schema.InputOption{}
	for _, classTypeDB := range notHiddenClassTypesDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%s", classTypeDB.ClassTypeName),
			InputOptionValue: fmt.Sprintf("%d", classTypeDB.ClassTypeID),
		})
	}
	return inputOptions
}

func GetNotHiddenClassTypesDB(classTypesDB []*ClassTypeDB) []*ClassTypeDB {
	notHiddenClassTypesDB := []*ClassTypeDB{}
	for _, classTypeDB := range classTypesDB {
		if !classTypeDB.ClassTypeIsHidden {
			notHiddenClassTypesDB = append(notHiddenClassTypesDB, classTypeDB)
		}
	}
	return notHiddenClassTypesDB
}

func ValidateClassTypeDB(classTypeDB *ClassTypeDB) (err error) {
	if classTypeDB == nil {
		panic("Object is nil")
	}
	if classTypeDB.ClassTypeID <= 0 || classTypeDB.ClassTypeName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateClassTypeInsert(classTypeInsert *ClassTypeInsert) (err error) {
	if classTypeInsert == nil {
		panic("Object is nil")
	}
	if classTypeInsert.ClassTypeName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateClassTypeUpdate(classTypeUpdate *ClassTypeUpdate) (err error) {
	if classTypeUpdate == nil {
		panic("Object is nil")
	}
	if classTypeUpdate.ClassTypeID <= 0 || classTypeUpdate.ClassTypeName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateClassTypeDelete(classTypeDelete *ClassTypeDelete) (err error) {
	if classTypeDelete == nil {
		panic("Object is nil")
	}
	if classTypeDelete.ClassTypeID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateClassTypeGet(classTypeGet *ClassTypeGet) (err error) {
	if classTypeGet == nil {
		panic("Object is nil")
	}
	if classTypeGet.ClassTypeID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateClassTypesGet(classTypeGets *ClassTypesGet) (err error) {
	if classTypeGets == nil {
		panic("Object is nil")
	}
	return nil
}
