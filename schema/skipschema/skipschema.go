package skipschema

import (
	"fmt"
)

type SkipDB struct {
	SkipID    int `json:"SkipID" server:"y"`
	ClassID   int `json:"ClassID" server:"y"`
	StudentID int `json:"StudentID" server:"y"`
}

type SkipInsert struct {
	ClassID   int `json:"ClassID" server:"y"`
	StudentID int `json:"StudentID" server:"y"`
}

type SkipUpdate struct {
	SkipID    int `json:"SkipID" server:"y"`
	ClassID   int `json:"ClassID" server:"y"`
	StudentID int `json:"StudentID" server:"y"`
}

type SkipDelete struct {
	SkipID int `json:"SkipID" server:"y"`
}

type SkipGet struct {
	SkipID int `json:"SkipID" server:"y"`
}

type SkipsGet struct{}

func ValidateSkipDB(skipDB *SkipDB) (err error) {
	if skipDB == nil {
		panic("Object is nil")
	}
	if skipDB.SkipID <= 0 ||
		skipDB.ClassID <= 0 ||
		skipDB.StudentID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateSkipInsert(skipInsert *SkipInsert) (err error) {
	if skipInsert == nil {
		panic("Object is nil")
	}
	if skipInsert.ClassID <= 0 ||
		skipInsert.StudentID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateSkipUpdate(skipUpdate *SkipUpdate) (err error) {
	if skipUpdate == nil {
		panic("Object is nil")
	}
	if skipUpdate.SkipID <= 0 ||
		skipUpdate.ClassID <= 0 ||
		skipUpdate.StudentID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateSkipDelete(skipDelete *SkipDelete) (err error) {
	if skipDelete == nil {
		panic("Object is nil")
	}
	if skipDelete.SkipID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateSkipGet(skipGet *SkipGet) (err error) {
	if skipGet == nil {
		panic("Object is nil")
	}
	if skipGet.SkipID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateSkipsGet(skipGets *SkipsGet) (err error) {
	if skipGets == nil {
		panic("Object is nil")
	}
	return nil
}
