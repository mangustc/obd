package finhelpctgschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type FinhelpCtgDB struct {
	FinhelpCtgID          int    `json:"FinhelpCtgID" server:"y"`
	FinhelpCtgDescription string `json:"FinhelpCtgDescription"`
	FinhelpCtgPayment     int    `json:"FinhelpCtgPayment"`
	FinhelpCtgIsHidden    bool   `json:"FinhelpCtgIsHidden"`
}

type FinhelpCtgInsert struct {
	FinhelpCtgDescription string `json:"FinhelpCtgDescription"`
	FinhelpCtgPayment     int    `json:"FinhelpCtgPayment"`
}

type FinhelpCtgUpdate struct {
	FinhelpCtgID          int    `json:"FinhelpCtgID" server:"y"`
	FinhelpCtgDescription string `json:"FinhelpCtgDescription"`
	FinhelpCtgPayment     int    `json:"FinhelpCtgPayment"`
}

type FinhelpCtgDelete struct {
	FinhelpCtgID int `json:"FinhelpCtgID" server:"y"`
}

type FinhelpCtgGet struct {
	FinhelpCtgID int `json:"FinhelpCtgID" server:"y"`
}

type FinhelpCtgsGet struct{}

func GetFinhelpCtgInputOptionsFromFinhelpCtgsDB(finhelpCtgsDB []*FinhelpCtgDB) []*schema.InputOption {
	notHiddenFinhelpCtgsDB := GetNotHiddenFinhelpCtgsDB(finhelpCtgsDB)
	inputOptions := []*schema.InputOption{}
	for _, finhelpCtgDB := range notHiddenFinhelpCtgsDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%d", finhelpCtgDB.FinhelpCtgID),
			InputOptionValue: fmt.Sprintf("%d", finhelpCtgDB.FinhelpCtgID),
		})
	}
	return inputOptions
}

func GetNotHiddenFinhelpCtgsDB(finhelpCtgsDB []*FinhelpCtgDB) []*FinhelpCtgDB {
	notHiddenFinhelpCtgsDB := []*FinhelpCtgDB{}
	for _, finhelpCtgDB := range finhelpCtgsDB {
		if !finhelpCtgDB.FinhelpCtgIsHidden {
			notHiddenFinhelpCtgsDB = append(notHiddenFinhelpCtgsDB, finhelpCtgDB)
		}
	}
	return notHiddenFinhelpCtgsDB
}

func ValidateFinhelpCtgDB(finhelpCtgDB *FinhelpCtgDB) (err error) {
	if finhelpCtgDB == nil {
		panic("Object is nil")
	}
	if finhelpCtgDB.FinhelpCtgID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpCtgInsert(finhelpCtgInsert *FinhelpCtgInsert) (err error) {
	if finhelpCtgInsert == nil {
		panic("Object is nil")
	}
	return nil
}

func ValidateFinhelpCtgUpdate(finhelpCtgUpdate *FinhelpCtgUpdate) (err error) {
	if finhelpCtgUpdate == nil {
		panic("Object is nil")
	}
	if finhelpCtgUpdate.FinhelpCtgID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpCtgDelete(finhelpCtgDelete *FinhelpCtgDelete) (err error) {
	if finhelpCtgDelete == nil {
		panic("Object is nil")
	}
	if finhelpCtgDelete.FinhelpCtgID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpCtgGet(finhelpCtgGet *FinhelpCtgGet) (err error) {
	if finhelpCtgGet == nil {
		panic("Object is nil")
	}
	if finhelpCtgGet.FinhelpCtgID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpCtgsGet(finhelpCtgGets *FinhelpCtgsGet) (err error) {
	if finhelpCtgGets == nil {
		panic("Object is nil")
	}
	return nil
}
