package studentschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
	"github.com/mangustc/obd/schema/groupschema"
)

type StudentDB struct {
	StudentID          int    `json:"StudentID" server:"y"`
	StudentLastname    string `json:"StudentLastname"`
	StudentFirstname   string `json:"StudentFirstname"`
	StudentMiddlename  string `json:"StudentMiddlename"`
	StudentPhoneNumber string `json:"StudentPhoneNumber"`
	StudentIsHidden    bool   `json:"StudentIsHidden"`
	GroupID            int    `json:"GroupID" server:"y"`
}

type StudentInsert struct {
	StudentLastname    string `json:"StudentLastname"`
	StudentFirstname   string `json:"StudentFirstname"`
	StudentMiddlename  string `json:"StudentMiddlename"`
	StudentPhoneNumber string `json:"StudentPhoneNumber"`
	GroupID            int    `json:"GroupID" server:"y"`
}

type StudentUpdate struct {
	StudentID          int    `json:"StudentID" server:"y"`
	StudentLastname    string `json:"StudentLastname"`
	StudentFirstname   string `json:"StudentFirstname"`
	StudentMiddlename  string `json:"StudentMiddlename"`
	StudentPhoneNumber string `json:"StudentPhoneNumber"`
	GroupID            int    `json:"GroupID" server:"y"`
}

type StudentDelete struct {
	StudentID int `json:"StudentID" server:"y"`
}

type StudentGet struct {
	StudentID int `json:"StudentID" server:"y"`
}

type StudentsGet struct{}

func GetStudentInputOptionsFromStudentsDB(studentsDB []*StudentDB, groupsDB []*groupschema.GroupDB) []*schema.InputOption {
	notHiddenStudentsDB := GetNotHiddenStudentsDB(studentsDB)
	inputOptions := []*schema.InputOption{}
	for _, studentDB := range notHiddenStudentsDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%s %s %s %s", studentDB.StudentLastname, studentDB.StudentFirstname, studentDB.StudentMiddlename, groupschema.GetGroupByGroupID(groupsDB, studentDB.GroupID).GroupNumber),
			InputOptionValue: fmt.Sprintf("%d", studentDB.StudentID),
		})
	}
	return inputOptions
}

func GetNotHiddenStudentsDB(studentsDB []*StudentDB) []*StudentDB {
	notHiddenStudentsDB := []*StudentDB{}
	for _, studentDB := range studentsDB {
		if !studentDB.StudentIsHidden {
			notHiddenStudentsDB = append(notHiddenStudentsDB, studentDB)
		}
	}
	return notHiddenStudentsDB
}

func ValidateStudentDB(studentDB *StudentDB) (err error) {
	if studentDB == nil {
		return fmt.Errorf("Object is nil")
	}
	if studentDB.StudentID <= 0 || studentDB.StudentLastname == "" || studentDB.StudentFirstname == "" ||
		studentDB.GroupID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateStudentInsert(studentInsert *StudentInsert) (err error) {
	if studentInsert == nil {
		return fmt.Errorf("Object is nil")
	}
	if studentInsert.StudentLastname == "" || studentInsert.StudentFirstname == "" ||
		studentInsert.GroupID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateStudentUpdate(studentUpdate *StudentUpdate) (err error) {
	if studentUpdate == nil {
		return fmt.Errorf("Object is nil")
	}
	if studentUpdate.StudentID <= 0 || studentUpdate.StudentLastname == "" || studentUpdate.StudentFirstname == "" ||
		studentUpdate.GroupID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateStudentDelete(studentDelete *StudentDelete) (err error) {
	if studentDelete == nil {
		return fmt.Errorf("Object is nil")
	}
	if studentDelete.StudentID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateStudentGet(studentGet *StudentGet) (err error) {
	if studentGet == nil {
		return fmt.Errorf("Object is nil")
	}
	if studentGet.StudentID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateStudentsGet(studentGets *StudentsGet) (err error) {
	if studentGets == nil {
		return fmt.Errorf("Object is nil")
	}
	return nil
}
