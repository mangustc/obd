package jobschema

import (
	"fmt"

	"github.com/mangustc/obd/schema"
)

type JobDB struct {
	JobID                 int    `json:"JobID"`
	JobName               string `json:"JobName"`
	JobAccessUser         bool   `json:"JobAccessUser"`
	JobAccessJob          bool   `json:"JobAccessJob"`
	JobAccessStudent      bool   `json:"JobAccessStudent"`
	JobAccessGroup        bool   `json:"JobAccessGroup"`
	JobAccessFinhelpCtg   bool   `json:"JobAccessFinhelpCtg"`
	JobAccessFinhelpStage bool   `json:"JobAccessFinhelpStage"`
	JobAccessFinhelpProc  bool   `json:"JobAccessProc"`
	JobAccessBuilding     bool   `json:"JobAccessBuilding"`
	JobAccessCabinetType  bool   `json:"JobAccessCabinetType"`
	JobAccessCabinet      bool   `json:"JobAccessCabinet"`
	JobAccessClassType    bool   `json:"JobAccessClassType"`
	JobAccessProf         bool   `json:"JobAccessProf"`
	JobAccessCourseType   bool   `json:"JobAccessCourseType"`
	JobAccessCourse       bool   `json:"JobAccessCourse"`
	JobAccessPerf         bool   `json:"JobAccessPerf"`
	JobAccessSkip         bool   `json:"JobAccessSkip"`
	JobAccessClass        bool   `json:"JobAccessClass"`
}

type JobInsert struct {
	JobName               string `json:"JobName"`
	JobAccessUser         bool   `json:"JobAccessUser"`
	JobAccessJob          bool   `json:"JobAccessJob"`
	JobAccessStudent      bool   `json:"JobAccessStudent"`
	JobAccessGroup        bool   `json:"JobAccessGroup"`
	JobAccessFinhelpCtg   bool   `json:"JobAccessFinhelpCtg"`
	JobAccessFinhelpStage bool   `json:"JobAccessFinhelpStage"`
	JobAccessFinhelpProc  bool   `json:"JobAccessProc"`
	JobAccessBuilding     bool   `json:"JobAccessBuilding"`
	JobAccessCabinetType  bool   `json:"JobAccessCabinetType"`
	JobAccessCabinet      bool   `json:"JobAccessCabinet"`
	JobAccessClassType    bool   `json:"JobAccessClassType"`
	JobAccessProf         bool   `json:"JobAccessProf"`
	JobAccessCourseType   bool   `json:"JobAccessCourseType"`
	JobAccessCourse       bool   `json:"JobAccessCourse"`
	JobAccessPerf         bool   `json:"JobAccessPerf"`
	JobAccessSkip         bool   `json:"JobAccessSkip"`
	JobAccessClass        bool   `json:"JobAccessClass"`
}

type JobUpdate struct {
	JobID                 int    `json:"JobID"`
	JobName               string `json:"JobName"`
	JobAccessUser         bool   `json:"JobAccessUser"`
	JobAccessJob          bool   `json:"JobAccessJob"`
	JobAccessStudent      bool   `json:"JobAccessStudent"`
	JobAccessGroup        bool   `json:"JobAccessGroup"`
	JobAccessFinhelpCtg   bool   `json:"JobAccessFinhelpCtg"`
	JobAccessFinhelpStage bool   `json:"JobAccessFinhelpStage"`
	JobAccessFinhelpProc  bool   `json:"JobAccessProc"`
	JobAccessBuilding     bool   `json:"JobAccessBuilding"`
	JobAccessCabinetType  bool   `json:"JobAccessCabinetType"`
	JobAccessCabinet      bool   `json:"JobAccessCabinet"`
	JobAccessClassType    bool   `json:"JobAccessClassType"`
	JobAccessProf         bool   `json:"JobAccessProf"`
	JobAccessCourseType   bool   `json:"JobAccessCourseType"`
	JobAccessCourse       bool   `json:"JobAccessCourse"`
	JobAccessPerf         bool   `json:"JobAccessPerf"`
	JobAccessSkip         bool   `json:"JobAccessSkip"`
	JobAccessClass        bool   `json:"JobAccessClass"`
}

type JobDelete struct {
	JobID int `json:"JobID"`
}

type JobGet struct {
	JobID int `json:"JobID"`
}

type JobsGet struct{}

func GetJobInputOptionsFromJobsDB(jobsDB []*JobDB) []*schema.InputOption {
	inputOptions := []*schema.InputOption{}
	for _, jobDB := range jobsDB {
		inputOptions = append(inputOptions, &schema.InputOption{
			InputOptionLabel: fmt.Sprintf("%s", jobDB.JobName),
			InputOptionValue: fmt.Sprintf("%d", jobDB.JobID),
		})
	}
	return inputOptions
}

func ValidateJobDB(jobDB *JobDB) (err error) {
	if jobDB == nil {
		return fmt.Errorf("Object is nil")
	}
	if jobDB.JobID <= 0 || jobDB.JobName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateJobInsert(jobInsert *JobInsert) (err error) {
	if jobInsert == nil {
		return fmt.Errorf("Object is nil")
	}
	if jobInsert.JobName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateJobUpdate(jobUpdate *JobUpdate) (err error) {
	if jobUpdate == nil {
		return fmt.Errorf("Object is nil")
	}
	if jobUpdate.JobID <= 0 || jobUpdate.JobName == "" {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateJobDelete(jobDelete *JobDelete) (err error) {
	if jobDelete == nil {
		return fmt.Errorf("Object is nil")
	}
	if jobDelete.JobID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateJobGet(jobGet *JobGet) (err error) {
	if jobGet == nil {
		return fmt.Errorf("Object is nil")
	}
	if jobGet.JobID <= 0 {
		return fmt.Errorf("One or more neccesary arguments are zero")
	}
	return nil
}

func ValidateJobsGet(jobsGet *JobsGet) (err error) {
	if jobsGet == nil {
		return fmt.Errorf("Object is nil")
	}
	return nil
}
