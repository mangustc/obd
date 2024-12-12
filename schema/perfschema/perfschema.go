package perfschema

import (
	"fmt"
)

type PerfDB struct {
	PerfID    int `json:"PerfID" server:"y"`
	CourseID  int `json:"CourseID" server:"y"`
	StudentID int `json:"StudentID" server:"y"`
	PerfGrade int `json:"PerfGrade"`
}

type PerfInsert struct {
	CourseID  int `json:"CourseID" server:"y"`
	StudentID int `json:"StudentID" server:"y"`
	PerfGrade int `json:"PerfGrade"`
}

type PerfUpdate struct {
	PerfID    int `json:"PerfID" server:"y"`
	CourseID  int `json:"CourseID" server:"y"`
	StudentID int `json:"StudentID" server:"y"`
	PerfGrade int `json:"PerfGrade"`
}

type PerfDelete struct {
	PerfID int `json:"PerfID" server:"y"`
}

type PerfGet struct {
	PerfID int `json:"PerfID" server:"y"`
}

type PerfsGet struct{}

func ValidatePerfDB(perfDB *PerfDB) (err error) {
	if perfDB == nil {
		panic("Object is nil")
	}
	if perfDB.PerfGrade < 0 || perfDB.PerfGrade > 5 {
		return fmt.Errorf("Perf grade should be in range of [0, 5]")
	}
	if perfDB.PerfID <= 0 ||
		perfDB.CourseID <= 0 ||
		perfDB.StudentID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidatePerfInsert(perfInsert *PerfInsert) (err error) {
	if perfInsert == nil {
		panic("Object is nil")
	}
	if perfInsert.PerfGrade < 0 || perfInsert.PerfGrade > 5 {
		return fmt.Errorf("Perf grade should be in range of [0, 5]")
	}
	if perfInsert.CourseID <= 0 ||
		perfInsert.StudentID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidatePerfUpdate(perfUpdate *PerfUpdate) (err error) {
	if perfUpdate == nil {
		panic("Object is nil")
	}
	if perfUpdate.PerfGrade < 0 || perfUpdate.PerfGrade > 5 {
		return fmt.Errorf("Perf grade should be in range of [0, 5]")
	}
	if perfUpdate.PerfID <= 0 ||
		perfUpdate.CourseID <= 0 ||
		perfUpdate.StudentID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidatePerfDelete(perfDelete *PerfDelete) (err error) {
	if perfDelete == nil {
		panic("Object is nil")
	}
	if perfDelete.PerfID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidatePerfGet(perfGet *PerfGet) (err error) {
	if perfGet == nil {
		panic("Object is nil")
	}
	if perfGet.PerfID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidatePerfsGet(perfGets *PerfsGet) (err error) {
	if perfGets == nil {
		panic("Object is nil")
	}
	return nil
}
