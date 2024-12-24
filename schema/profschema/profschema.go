package profschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type ProfDB struct {
	ProfID          int    `json:"ProfID" server:"y"`
	ProfLastname    string `json:"ProfLastname"`
	ProfFirstname   string `json:"ProfFirstname"`
	ProfMiddlename  string `json:"ProfMiddlename"`
	ProfPhoneNumber string `json:"ProfPhoneNumber"`
	ProfEmail       string `json:"ProfEmail"`
	ProfIsHidden    bool   `json:"ProfIsHidden"`
}

type ProfInsert struct {
	ProfLastname    string `json:"ProfLastname"`
	ProfFirstname   string `json:"ProfFirstname"`
	ProfMiddlename  string `json:"ProfMiddlename"`
	ProfPhoneNumber string `json:"ProfPhoneNumber"`
	ProfEmail       string `json:"ProfEmail"`
}

type ProfUpdate struct {
	ProfID          int    `json:"ProfID" server:"y"`
	ProfLastname    string `json:"ProfLastname"`
	ProfFirstname   string `json:"ProfFirstname"`
	ProfMiddlename  string `json:"ProfMiddlename"`
	ProfPhoneNumber string `json:"ProfPhoneNumber"`
	ProfEmail       string `json:"ProfEmail"`
}

type ProfDelete struct {
	ProfID int `json:"ProfID" server:"y"`
}

type ProfGet struct {
	ProfID int `json:"ProfID" server:"y"`
}

type ProfsGet struct{}

func GetProfInputOptionsFromProfsDB(profsDB []*ProfDB) []*schema.InputOption {
	notHiddenProfsDB := GetNotHiddenProfsDB(profsDB)
	inputOptions := []*schema.InputOption{}
	for _, profDB := range notHiddenProfsDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%s %s %s", profDB.ProfLastname, profDB.ProfFirstname, profDB.ProfMiddlename),
			InputOptionValue: fmt.Sprintf("%d", profDB.ProfID),
		})
	}
	return inputOptions
}

func GetNotHiddenProfsDB(profsDB []*ProfDB) []*ProfDB {
	notHiddenProfsDB := []*ProfDB{}
	for _, profDB := range profsDB {
		if !profDB.ProfIsHidden {
			notHiddenProfsDB = append(notHiddenProfsDB, profDB)
		}
	}
	return notHiddenProfsDB
}

func ValidateProfDB(profDB *ProfDB) (err error) {
	if profDB == nil {
		return fmt.Errorf("Object is nil")
	}
	if profDB.ProfID <= 0 || profDB.ProfLastname == "" || profDB.ProfFirstname == "" || profDB.ProfEmail == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateProfInsert(profInsert *ProfInsert) (err error) {
	if profInsert == nil {
		return fmt.Errorf("Object is nil")
	}
	if profInsert.ProfLastname == "" || profInsert.ProfFirstname == "" || profInsert.ProfEmail == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateProfUpdate(profUpdate *ProfUpdate) (err error) {
	if profUpdate == nil {
		return fmt.Errorf("Object is nil")
	}
	if profUpdate.ProfID <= 0 || profUpdate.ProfLastname == "" || profUpdate.ProfFirstname == "" || profUpdate.ProfEmail == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateProfDelete(profDelete *ProfDelete) (err error) {
	if profDelete == nil {
		return fmt.Errorf("Object is nil")
	}
	if profDelete.ProfID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateProfGet(profGet *ProfGet) (err error) {
	if profGet == nil {
		return fmt.Errorf("Object is nil")
	}
	if profGet.ProfID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateProfsGet(profGets *ProfsGet) (err error) {
	if profGets == nil {
		return fmt.Errorf("Object is nil")
	}
	return nil
}
