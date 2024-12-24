package coursetypeschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type CourseTypeDB struct {
	CourseTypeID       int    `json:"CourseTypeID" server:"y"`
	CourseTypeName     string `json:"CourseTypeName"`
	CourseTypeIsHidden bool   `json:"CourseTypeIsHidden"`
}

type CourseTypeInsert struct {
	CourseTypeName string `json:"CourseTypeName"`
}

type CourseTypeUpdate struct {
	CourseTypeID   int    `json:"CourseTypeID" server:"y"`
	CourseTypeName string `json:"CourseTypeName"`
}

type CourseTypeDelete struct {
	CourseTypeID int `json:"CourseTypeID" server:"y"`
}

type CourseTypeGet struct {
	CourseTypeID int `json:"CourseTypeID" server:"y"`
}

type CourseTypesGet struct{}

func GetCourseTypeInputOptionsFromCourseTypesDB(courseTypesDB []*CourseTypeDB) []*schema.InputOption {
	notHiddenCourseTypesDB := GetNotHiddenCourseTypesDB(courseTypesDB)
	inputOptions := []*schema.InputOption{}
	for _, courseTypeDB := range notHiddenCourseTypesDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%s", courseTypeDB.CourseTypeName),
			InputOptionValue: fmt.Sprintf("%d", courseTypeDB.CourseTypeID),
		})
	}
	return inputOptions
}

func GetNotHiddenCourseTypesDB(courseTypesDB []*CourseTypeDB) []*CourseTypeDB {
	notHiddenCourseTypesDB := []*CourseTypeDB{}
	for _, courseTypeDB := range courseTypesDB {
		if !courseTypeDB.CourseTypeIsHidden {
			notHiddenCourseTypesDB = append(notHiddenCourseTypesDB, courseTypeDB)
		}
	}
	return notHiddenCourseTypesDB
}

func ValidateCourseTypeDB(courseTypeDB *CourseTypeDB) (err error) {
	if courseTypeDB == nil {
		panic("Object is nil")
	}
	if courseTypeDB.CourseTypeID <= 0 || courseTypeDB.CourseTypeName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCourseTypeInsert(courseTypeInsert *CourseTypeInsert) (err error) {
	if courseTypeInsert == nil {
		panic("Object is nil")
	}
	if courseTypeInsert.CourseTypeName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCourseTypeUpdate(courseTypeUpdate *CourseTypeUpdate) (err error) {
	if courseTypeUpdate == nil {
		panic("Object is nil")
	}
	if courseTypeUpdate.CourseTypeID <= 0 || courseTypeUpdate.CourseTypeName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCourseTypeDelete(courseTypeDelete *CourseTypeDelete) (err error) {
	if courseTypeDelete == nil {
		panic("Object is nil")
	}
	if courseTypeDelete.CourseTypeID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCourseTypeGet(courseTypeGet *CourseTypeGet) (err error) {
	if courseTypeGet == nil {
		panic("Object is nil")
	}
	if courseTypeGet.CourseTypeID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCourseTypesGet(courseTypeGets *CourseTypesGet) (err error) {
	if courseTypeGets == nil {
		panic("Object is nil")
	}
	return nil
}
