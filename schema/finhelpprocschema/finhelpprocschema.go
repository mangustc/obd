package finhelpprocschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type FinhelpProcDB struct {
	FinhelpProcID          int    `json:"FinhelpProcID" server:"y"`
	FinhelpProcDescription string `json:"FinhelpProcDescription"`
	FinhelpProcPayment     int    `json:"FinhelpProcPayment"`
	FinhelpProcIsHidden    bool   `json:"FinhelpProcIsHidden"`
}

type FinhelpProcInsert struct {
	FinhelpProcDescription string `json:"FinhelpProcDescription"`
	FinhelpProcPayment     int    `json:"FinhelpProcPayment"`
}

type FinhelpProcUpdate struct {
	FinhelpProcID          int    `json:"FinhelpProcID" server:"y"`
	FinhelpProcDescription string `json:"FinhelpProcDescription"`
	FinhelpProcPayment     int    `json:"FinhelpProcPayment"`
}

type FinhelpProcDelete struct {
	FinhelpProcID int `json:"FinhelpProcID" server:"y"`
}

type FinhelpProcGet struct {
	FinhelpProcID int `json:"FinhelpProcID" server:"y"`
}

type FinhelpProcsGet struct{}

func GetFinhelpProcInputOptionsFromFinhelpProcsDB(finhelpProcsDB []*FinhelpProcDB) []*schema.InputOption {
	notHiddenFinhelpProcsDB := GetNotHiddenFinhelpProcsDB(finhelpProcsDB)
	inputOptions := []*schema.InputOption{}
	for _, finhelpProcDB := range notHiddenFinhelpProcsDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%d", finhelpProcDB.FinhelpProcID),
			InputOptionValue: fmt.Sprintf("%d", finhelpProcDB.FinhelpProcID),
		})
	}
	return inputOptions
}

func GetNotHiddenFinhelpProcsDB(finhelpProcsDB []*FinhelpProcDB) []*FinhelpProcDB {
	notHiddenFinhelpProcsDB := []*FinhelpProcDB{}
	for _, finhelpProcDB := range finhelpProcsDB {
		if !finhelpProcDB.FinhelpProcIsHidden {
			notHiddenFinhelpProcsDB = append(notHiddenFinhelpProcsDB, finhelpProcDB)
		}
	}
	return notHiddenFinhelpProcsDB
}

func ValidateFinhelpProcDB(finhelpProcDB *FinhelpProcDB) (err error) {
	if finhelpProcDB == nil {
		panic("Object is nil")
	}
	if finhelpProcDB.FinhelpProcID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpProcInsert(finhelpProcInsert *FinhelpProcInsert) (err error) {
	if finhelpProcInsert == nil {
		panic("Object is nil")
	}
	return nil
}

func ValidateFinhelpProcUpdate(finhelpProcUpdate *FinhelpProcUpdate) (err error) {
	if finhelpProcUpdate == nil {
		panic("Object is nil")
	}
	if finhelpProcUpdate.FinhelpProcID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpProcDelete(finhelpProcDelete *FinhelpProcDelete) (err error) {
	if finhelpProcDelete == nil {
		panic("Object is nil")
	}
	if finhelpProcDelete.FinhelpProcID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpProcGet(finhelpProcGet *FinhelpProcGet) (err error) {
	if finhelpProcGet == nil {
		panic("Object is nil")
	}
	if finhelpProcGet.FinhelpProcID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpProcsGet(finhelpProcGets *FinhelpProcsGet) (err error) {
	if finhelpProcGets == nil {
		panic("Object is nil")
	}
	return nil
}
