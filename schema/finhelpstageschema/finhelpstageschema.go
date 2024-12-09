package finhelpstageschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type FinhelpStageDB struct {
	FinhelpStageID          int    `json:"FinhelpStageID" server:"y"`
	FinhelpStageDescription string `json:"FinhelpStageDescription"`
	FinhelpStageIsHidden    bool   `json:"FinhelpStageIsHidden"`
}

type FinhelpStageInsert struct {
	FinhelpStageDescription string `json:"FinhelpStageDescription"`
}

type FinhelpStageUpdate struct {
	FinhelpStageID          int    `json:"FinhelpStageID" server:"y"`
	FinhelpStageDescription string `json:"FinhelpStageDescription"`
}

type FinhelpStageDelete struct {
	FinhelpStageID int `json:"FinhelpStageID" server:"y"`
}

type FinhelpStageGet struct {
	FinhelpStageID int `json:"FinhelpStageID" server:"y"`
}

type FinhelpStagesGet struct{}

func GetFinhelpStageInputOptionsFromFinhelpStagesDB(finhelpStagesDB []*FinhelpStageDB) []*schema.InputOption {
	notHiddenFinhelpStagesDB := GetNotHiddenFinhelpStagesDB(finhelpStagesDB)
	inputOptions := []*schema.InputOption{}
	for _, finhelpStageDB := range notHiddenFinhelpStagesDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%d", finhelpStageDB.FinhelpStageID),
			InputOptionValue: fmt.Sprintf("%d", finhelpStageDB.FinhelpStageID),
		})
	}
	return inputOptions
}

func GetNotHiddenFinhelpStagesDB(finhelpStagesDB []*FinhelpStageDB) []*FinhelpStageDB {
	notHiddenFinhelpStagesDB := []*FinhelpStageDB{}
	for _, finhelpStageDB := range finhelpStagesDB {
		if !finhelpStageDB.FinhelpStageIsHidden {
			notHiddenFinhelpStagesDB = append(notHiddenFinhelpStagesDB, finhelpStageDB)
		}
	}
	return notHiddenFinhelpStagesDB
}

func ValidateFinhelpStageDB(finhelpStageDB *FinhelpStageDB) (err error) {
	if finhelpStageDB == nil {
		panic("Object is nil")
	}
	if finhelpStageDB.FinhelpStageID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpStageInsert(finhelpStageInsert *FinhelpStageInsert) (err error) {
	if finhelpStageInsert == nil {
		panic("Object is nil")
	}
	return nil
}

func ValidateFinhelpStageUpdate(finhelpStageUpdate *FinhelpStageUpdate) (err error) {
	if finhelpStageUpdate == nil {
		panic("Object is nil")
	}
	if finhelpStageUpdate.FinhelpStageID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpStageDelete(finhelpStageDelete *FinhelpStageDelete) (err error) {
	if finhelpStageDelete == nil {
		panic("Object is nil")
	}
	if finhelpStageDelete.FinhelpStageID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpStageGet(finhelpStageGet *FinhelpStageGet) (err error) {
	if finhelpStageGet == nil {
		panic("Object is nil")
	}
	if finhelpStageGet.FinhelpStageID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpStagesGet(finhelpStageGets *FinhelpStagesGet) (err error) {
	if finhelpStageGets == nil {
		panic("Object is nil")
	}
	return nil
}
