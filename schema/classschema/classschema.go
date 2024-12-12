package classschema

import (
	"fmt"
	"time"

	"github.com/mangustc/obd/schema"
)

type ClassDB struct {
	ClassID     int    `json:"ClassID" server:"y"`
	ClassTypeID int    `json:"ClassTypeID" server:"y"`
	ProfID      int    `json:"ProfID" server:"y"`
	CabinetID   int    `json:"CabinetID" server:"y"`
	CourseID    int    `json:"CourseID" server:"y"`
	GroupID     int    `json:"GroupID" server:"y"`
	ClassStart  string `json:"ClassStart"`
	ClassNumber int    `json:"ClassNumber"`
}

type ClassInsert struct {
	ClassTypeID int    `json:"ClassTypeID" server:"y"`
	ProfID      int    `json:"ProfID" server:"y"`
	CabinetID   int    `json:"CabinetID" server:"y"`
	CourseID    int    `json:"CourseID" server:"y"`
	GroupID     int    `json:"GroupID" server:"y"`
	ClassStart  string `json:"ClassStart"`
	ClassNumber int    `json:"ClassNumber"`
}

type ClassUpdate struct {
	ClassID     int    `json:"ClassID" server:"y"`
	ClassTypeID int    `json:"ClassTypeID" server:"y"`
	ProfID      int    `json:"ProfID" server:"y"`
	CabinetID   int    `json:"CabinetID" server:"y"`
	CourseID    int    `json:"CourseID" server:"y"`
	GroupID     int    `json:"GroupID" server:"y"`
	ClassStart  string `json:"ClassStart"`
	ClassNumber int    `json:"ClassNumber"`
}

type ClassDelete struct {
	ClassID int `json:"ClassID" server:"y"`
}

type ClassGet struct {
	ClassID int `json:"ClassID" server:"y"`
}

type ClasssGet struct{}

func GetClassInputOptionsFromClasssDB(classsDB []*ClassDB) []*schema.InputOption {
	inputOptions := []*schema.InputOption{}
	for _, classDB := range classsDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%.10s Number %d", classDB.ClassStart, classDB.ClassNumber),
			InputOptionValue: fmt.Sprintf("%d", classDB.ClassID),
		})
	}
	return inputOptions
}

func ValidateClassDB(classDB *ClassDB) (err error) {
	if classDB == nil {
		panic("Object is nil")
	}
	_, err = time.Parse("2006-01-02", classDB.ClassStart)
	if err != nil {
		return fmt.Errorf("Time is not in the right format")
	}
	if classDB.ClassNumber < 1 || classDB.ClassNumber > 7 {
		return fmt.Errorf("Class number should be in range of [1, 7]")
	}
	if classDB.ClassID <= 0 ||
		classDB.ClassTypeID <= 0 ||
		classDB.ProfID <= 0 ||
		classDB.CabinetID <= 0 ||
		classDB.CourseID <= 0 ||
		classDB.GroupID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateClassInsert(classInsert *ClassInsert) (err error) {
	if classInsert == nil {
		panic("Object is nil")
	}
	_, err = time.Parse("2006-01-02", classInsert.ClassStart)
	if err != nil {
		return fmt.Errorf("Time is not in the right format")
	}
	if classInsert.ClassNumber < 1 || classInsert.ClassNumber > 7 {
		return fmt.Errorf("Class number should be in range of [1, 7]")
	}
	if classInsert.ClassTypeID <= 0 ||
		classInsert.ProfID <= 0 ||
		classInsert.CabinetID <= 0 ||
		classInsert.CourseID <= 0 ||
		classInsert.GroupID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateClassUpdate(classUpdate *ClassUpdate) (err error) {
	if classUpdate == nil {
		panic("Object is nil")
	}
	_, err = time.Parse("2006-01-02", classUpdate.ClassStart)
	if err != nil {
		return fmt.Errorf("Time is not in the right format")
	}
	if classUpdate.ClassNumber < 1 || classUpdate.ClassNumber > 7 {
		return fmt.Errorf("Class number should be in range of [1, 7]")
	}
	if classUpdate.ClassID <= 0 ||
		classUpdate.ClassTypeID <= 0 ||
		classUpdate.ProfID <= 0 ||
		classUpdate.CabinetID <= 0 ||
		classUpdate.CourseID <= 0 ||
		classUpdate.GroupID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateClassDelete(classDelete *ClassDelete) (err error) {
	if classDelete == nil {
		panic("Object is nil")
	}
	if classDelete.ClassID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateClassGet(classGet *ClassGet) (err error) {
	if classGet == nil {
		panic("Object is nil")
	}
	if classGet.ClassID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateClasssGet(classGets *ClasssGet) (err error) {
	if classGets == nil {
		panic("Object is nil")
	}
	return nil
}
