package courseschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type CourseDB struct {
	CourseID       int    `json:"CourseID" server:"y"`
	CourseTypeID   int    `json:"CourseTypeID" server:"y"`
	CourseName     string `json:"CourseName"`
	CourseIsHidden bool   `json:"CourseIsHidden"`
}

type CourseInsert struct {
	CourseTypeID int    `json:"CourseTypeID" server:"y"`
	CourseName   string `json:"CourseName"`
}

type CourseUpdate struct {
	CourseID     int    `json:"CourseID" server:"y"`
	CourseTypeID int    `json:"CourseTypeID" server:"y"`
	CourseName   string `json:"CourseName"`
}

type CourseDelete struct {
	CourseID int `json:"CourseID" server:"y"`
}

type CourseGet struct {
	CourseID int `json:"CourseID" server:"y"`
}

type CoursesGet struct{}

func GetCourseInputOptionsFromCoursesDB(coursesDB []*CourseDB) []*schema.InputOption {
	notHiddenCoursesDB := GetNotHiddenCoursesDB(coursesDB)
	inputOptions := []*schema.InputOption{}
	for _, courseDB := range notHiddenCoursesDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%s", courseDB.CourseName),
			InputOptionValue: fmt.Sprintf("%d", courseDB.CourseID),
		})
	}
	return inputOptions
}

func GetNotHiddenCoursesDB(coursesDB []*CourseDB) []*CourseDB {
	notHiddenCoursesDB := []*CourseDB{}
	for _, courseDB := range coursesDB {
		if !courseDB.CourseIsHidden {
			notHiddenCoursesDB = append(notHiddenCoursesDB, courseDB)
		}
	}
	return notHiddenCoursesDB
}

func ValidateCourseDB(courseDB *CourseDB) (err error) {
	if courseDB == nil {
		panic("Object is nil")
	}
	if courseDB.CourseID <= 0 || courseDB.CourseTypeID <= 0 || courseDB.CourseName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCourseInsert(courseInsert *CourseInsert) (err error) {
	if courseInsert == nil {
		panic("Object is nil")
	}
	if courseInsert.CourseTypeID <= 0 || courseInsert.CourseName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCourseUpdate(courseUpdate *CourseUpdate) (err error) {
	if courseUpdate == nil {
		panic("Object is nil")
	}
	if courseUpdate.CourseID <= 0 || courseUpdate.CourseTypeID <= 0 || courseUpdate.CourseName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCourseDelete(courseDelete *CourseDelete) (err error) {
	if courseDelete == nil {
		panic("Object is nil")
	}
	if courseDelete.CourseID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCourseGet(courseGet *CourseGet) (err error) {
	if courseGet == nil {
		panic("Object is nil")
	}
	if courseGet.CourseID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateCoursesGet(courseGets *CoursesGet) (err error) {
	if courseGets == nil {
		panic("Object is nil")
	}
	return nil
}
