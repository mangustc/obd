package jobschema

import (
	"fmt"
)

type JobDB struct {
	JobID   int
	JobName string
}

type JobInsert struct {
	JobName string
}

type JobUpdate struct {
	JobID   int
	JobName string
}

type JobDelete struct {
	JobID int
}

type JobGet struct {
	JobID int
}

type JobsGet struct{}

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
