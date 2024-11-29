package jobservice

import (
	"database/sql"
	"fmt"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/jobschema"
	"github.com/mangustc/obd/util"
	"github.com/mattn/go-sqlite3"
)

type JobService struct {
	db    *sql.DB
	jobTN string
}

func NewJobService(db *sql.DB, jobTN string) (us *JobService) {
	if db == nil || jobTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &JobService{
		db:    db,
		jobTN: jobTN,
	}
}

func (js *JobService) InsertJob(data *jobschema.JobInsert) (jobDB *jobschema.JobDB, err error) {
	query := fmt.Sprintf(`
INSERT INTO %[1]s
	(%[1]sName)
	VALUES (%[2]s)
	RETURNING *`,
		js.jobTN,
		data.JobName,
	)

	stmt, err := js.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	err = stmt.QueryRow().Scan(
		&jobDB.JobID,
		&jobDB.JobName,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	err = jobschema.ValidateJobDB(jobDB)
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return jobDB, nil
}

func (js *JobService) UpdateJob(data *jobschema.JobUpdate) (jobDB *jobschema.JobDB, err error) {
	query := fmt.Sprintf(`
UPDATE %[1]s SET
	%[1]sName = %[3]s
	WHERE %[1]sID = %[2]d
	RETURNING *`,
		js.jobTN,
		data.JobID,
		data.JobName,
	)

	stmt, err := js.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	err = stmt.QueryRow().Scan(
		&jobDB.JobID,
		&jobDB.JobName,
	)
	if err != nil {
		if util.IsErrorSQL(err, sqlite3.ErrConstraint) {
			return nil, errs.ErrUnprocessableEntity
		}
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	err = jobschema.ValidateJobDB(jobDB)
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	return jobDB, nil
}
