package finhelpprocschema

import (
	"fmt"
)

type FinhelpProcDB struct {
	FinhelpProcID        int    `json:"FinhelpProcID" server:"y"`
	FinhelpProcCreatedAt string `json:"FinhelpProcCreatedAt"`
	UserID               int    `json:"UserID" cookie:"y" server:"y"`
	StudentID            int    `json:"StudentID" server:"y"`
	FinhelpCtgID         int    `json:"FinhelpCtgID" server:"y"`
	FinhelpStageID       int    `json:"FinhelpStageID" server:"y"`
}

type FinhelpProcInsert struct {
	UserID         int `json:"UserID" cookie:"y" server:"y"`
	StudentID      int `json:"StudentID" server:"y"`
	FinhelpCtgID   int `json:"FinhelpCtgID" server:"y"`
	FinhelpStageID int `json:"FinhelpStageID" server:"y"`
}

type FinhelpProcUpdate struct {
	FinhelpProcID  int `json:"FinhelpProcID" server:"y"`
	UserID         int `json:"UserID" cookie:"y" server:"y"`
	StudentID      int `json:"StudentID" server:"y"`
	FinhelpCtgID   int `json:"FinhelpCtgID" server:"y"`
	FinhelpStageID int `json:"FinhelpStageID" server:"y"`
}

type FinhelpProcDelete struct {
	FinhelpProcID int `json:"FinhelpProcID" server:"y"`
}

type FinhelpProcGet struct {
	FinhelpProcID int `json:"FinhelpProcID" server:"y"`
}

type FinhelpProcsGet struct{}

func ValidateFinhelpProcDB(finhelpProcDB *FinhelpProcDB) (err error) {
	if finhelpProcDB == nil {
		panic("Object is nil")
	}
	if finhelpProcDB.FinhelpProcID <= 0 || finhelpProcDB.StudentID <= 0 || finhelpProcDB.UserID <= 0 ||
		finhelpProcDB.FinhelpCtgID <= 0 || finhelpProcDB.FinhelpStageID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpProcInsert(finhelpProcInsert *FinhelpProcInsert) (err error) {
	if finhelpProcInsert == nil {
		panic("Object is nil")
	}
	if finhelpProcInsert.StudentID <= 0 || finhelpProcInsert.UserID <= 0 ||
		finhelpProcInsert.FinhelpCtgID <= 0 || finhelpProcInsert.FinhelpStageID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateFinhelpProcUpdate(finhelpProcUpdate *FinhelpProcUpdate) (err error) {
	if finhelpProcUpdate == nil {
		panic("Object is nil")
	}
	if finhelpProcUpdate.FinhelpProcID <= 0 || finhelpProcUpdate.StudentID <= 0 || finhelpProcUpdate.UserID <= 0 ||
		finhelpProcUpdate.FinhelpCtgID <= 0 || finhelpProcUpdate.FinhelpStageID <= 0 {
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
