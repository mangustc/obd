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
INSERT INTO %[1]s (
		%[1]sName,
		%[1]sAccessUser,
		%[1]sAccessJob,
		%[1]sAccessStudent,
		%[1]sAccessUniGroup,
		%[1]sAccessFinhelpCtg,
		%[1]sAccessFinhelpStage,
		%[1]sAccessFinhelpProc,
		%[1]sAccessBuilding,
		%[1]sAccessCabinetType,
		%[1]sAccessCabinet,
		%[1]sAccessClassType,
		%[1]sAccessProf,
		%[1]sAccessCourseType,
		%[1]sAccessCourse,
		%[1]sAccessPerf,
		%[1]sAccessSkip,
		%[1]sAccessClass,
		%[1]sAccessSession
	)
	VALUES (
		"%[2]s",
		"%[3]t",
		"%[4]t",
		"%[5]t",
		"%[6]t",
		"%[7]t",
		"%[8]t",
		"%[9]t",
		"%[10]t",
		"%[11]t",
		"%[12]t",
		"%[13]t",
		"%[14]t",
		"%[15]t",
		"%[16]t",
		"%[17]t",
		"%[18]t",
		"%[19]t",
		"%[20]t"
	)
RETURNING *`,
		js.jobTN,
		data.JobName,
		data.JobAccessUser,
		data.JobAccessJob,
		data.JobAccessStudent,
		data.JobAccessUniGroup,
		data.JobAccessFinhelpCtg,
		data.JobAccessFinhelpStage,
		data.JobAccessFinhelpProc,
		data.JobAccessBuilding,
		data.JobAccessCabinetType,
		data.JobAccessCabinet,
		data.JobAccessClassType,
		data.JobAccessProf,
		data.JobAccessCourseType,
		data.JobAccessCourse,
		data.JobAccessPerf,
		data.JobAccessSkip,
		data.JobAccessClass,
		data.JobAccessSession,
	)

	stmt, err := js.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	jobDB = &jobschema.JobDB{}
	err = stmt.QueryRow().Scan(
		&jobDB.JobID,
		&jobDB.JobName,
		&jobDB.JobAccessUser,
		&jobDB.JobAccessJob,
		&jobDB.JobAccessStudent,
		&jobDB.JobAccessUniGroup,
		&jobDB.JobAccessFinhelpCtg,
		&jobDB.JobAccessFinhelpStage,
		&jobDB.JobAccessFinhelpProc,
		&jobDB.JobAccessBuilding,
		&jobDB.JobAccessCabinetType,
		&jobDB.JobAccessCabinet,
		&jobDB.JobAccessClassType,
		&jobDB.JobAccessProf,
		&jobDB.JobAccessCourseType,
		&jobDB.JobAccessCourse,
		&jobDB.JobAccessPerf,
		&jobDB.JobAccessSkip,
		&jobDB.JobAccessClass,
		&jobDB.JobAccessSession,
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
	%[1]sName = "%[3]s",
	%[1]sAccessUser = "%[4]t",
	%[1]sAccessJob = "%[5]t",
	%[1]sAccessStudent = "%[6]t",
	%[1]sAccessUniGroup = "%[7]t",
	%[1]sAccessFinhelpCtg = "%[8]t",
	%[1]sAccessFinhelpStage = "%[9]t",
	%[1]sAccessFinhelpProc = "%[10]t",
	%[1]sAccessBuilding = "%[11]t",
	%[1]sAccessCabinetType = "%[12]t",
	%[1]sAccessCabinet = "%[13]t",
	%[1]sAccessClassType = "%[14]t",
	%[1]sAccessProf = "%[15]t",
	%[1]sAccessCourseType = "%[16]t",
	%[1]sAccessCourse = "%[17]t",
	%[1]sAccessPerf = "%[18]t",
	%[1]sAccessSkip = "%[19]t",
	%[1]sAccessClass = "%[20]t",
	%[1]sAccessSession = "%[21]t"
WHERE %[1]sID = %[2]d
RETURNING *`,
		js.jobTN,
		data.JobID,
		data.JobName,
		data.JobAccessUser,
		data.JobAccessJob,
		data.JobAccessStudent,
		data.JobAccessUniGroup,
		data.JobAccessFinhelpCtg,
		data.JobAccessFinhelpStage,
		data.JobAccessFinhelpProc,
		data.JobAccessBuilding,
		data.JobAccessCabinetType,
		data.JobAccessCabinet,
		data.JobAccessClassType,
		data.JobAccessProf,
		data.JobAccessCourseType,
		data.JobAccessCourse,
		data.JobAccessPerf,
		data.JobAccessSkip,
		data.JobAccessClass,
		data.JobAccessSession,
	)

	stmt, err := js.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	jobDB = &jobschema.JobDB{}
	err = stmt.QueryRow().Scan(
		&jobDB.JobID,
		&jobDB.JobName,
		&jobDB.JobAccessUser,
		&jobDB.JobAccessJob,
		&jobDB.JobAccessStudent,
		&jobDB.JobAccessUniGroup,
		&jobDB.JobAccessFinhelpCtg,
		&jobDB.JobAccessFinhelpStage,
		&jobDB.JobAccessFinhelpProc,
		&jobDB.JobAccessBuilding,
		&jobDB.JobAccessCabinetType,
		&jobDB.JobAccessCabinet,
		&jobDB.JobAccessClassType,
		&jobDB.JobAccessProf,
		&jobDB.JobAccessCourseType,
		&jobDB.JobAccessCourse,
		&jobDB.JobAccessPerf,
		&jobDB.JobAccessSkip,
		&jobDB.JobAccessClass,
		&jobDB.JobAccessSession,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ErrNotFound
		}
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

func (js *JobService) DeleteJob(data *jobschema.JobDelete) (jobDB *jobschema.JobDB, err error) {
	query := fmt.Sprintf(`
DELETE FROM %[1]s WHERE %[1]sID = %[2]d
RETURNING *`,
		js.jobTN,
		data.JobID,
	)

	stmt, err := js.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	jobDB = &jobschema.JobDB{}
	err = stmt.QueryRow().Scan(
		&jobDB.JobID,
		&jobDB.JobName,
		&jobDB.JobAccessUser,
		&jobDB.JobAccessJob,
		&jobDB.JobAccessStudent,
		&jobDB.JobAccessUniGroup,
		&jobDB.JobAccessFinhelpCtg,
		&jobDB.JobAccessFinhelpStage,
		&jobDB.JobAccessFinhelpProc,
		&jobDB.JobAccessBuilding,
		&jobDB.JobAccessCabinetType,
		&jobDB.JobAccessCabinet,
		&jobDB.JobAccessClassType,
		&jobDB.JobAccessProf,
		&jobDB.JobAccessCourseType,
		&jobDB.JobAccessCourse,
		&jobDB.JobAccessPerf,
		&jobDB.JobAccessSkip,
		&jobDB.JobAccessClass,
		&jobDB.JobAccessSession,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ErrNotFound
		}
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

func (js *JobService) GetJob(data *jobschema.JobGet) (jobDB *jobschema.JobDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sName,
	%[1]sAccessUser,
	%[1]sAccessJob,
	%[1]sAccessStudent,
	%[1]sAccessUniGroup,
	%[1]sAccessFinhelpCtg,
	%[1]sAccessFinhelpStage,
	%[1]sAccessFinhelpProc,
	%[1]sAccessBuilding,
	%[1]sAccessCabinetType,
	%[1]sAccessCabinet,
	%[1]sAccessClassType,
	%[1]sAccessProf,
	%[1]sAccessCourseType,
	%[1]sAccessCourse,
	%[1]sAccessPerf,
	%[1]sAccessSkip,
	%[1]sAccessClass,
	%[1]sAccessSession
FROM %[1]s
WHERE %[1]sID = %[2]d`,
		js.jobTN,
		data.JobID,
	)

	stmt, err := js.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	jobDB = &jobschema.JobDB{}
	err = stmt.QueryRow().Scan(
		&jobDB.JobID,
		&jobDB.JobName,
		&jobDB.JobAccessUser,
		&jobDB.JobAccessJob,
		&jobDB.JobAccessStudent,
		&jobDB.JobAccessUniGroup,
		&jobDB.JobAccessFinhelpCtg,
		&jobDB.JobAccessFinhelpStage,
		&jobDB.JobAccessFinhelpProc,
		&jobDB.JobAccessBuilding,
		&jobDB.JobAccessCabinetType,
		&jobDB.JobAccessCabinet,
		&jobDB.JobAccessClassType,
		&jobDB.JobAccessProf,
		&jobDB.JobAccessCourseType,
		&jobDB.JobAccessCourse,
		&jobDB.JobAccessPerf,
		&jobDB.JobAccessSkip,
		&jobDB.JobAccessClass,
		&jobDB.JobAccessSession,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.ErrUnprocessableEntity
		}
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

func (js *JobService) GetJobs(data *jobschema.JobsGet) (jobsDB []*jobschema.JobDB, err error) {
	query := fmt.Sprintf(`
SELECT
	%[1]sID,
	%[1]sName,
	%[1]sAccessUser,
	%[1]sAccessJob,
	%[1]sAccessStudent,
	%[1]sAccessUniGroup,
	%[1]sAccessFinhelpCtg,
	%[1]sAccessFinhelpStage,
	%[1]sAccessFinhelpProc,
	%[1]sAccessBuilding,
	%[1]sAccessCabinetType,
	%[1]sAccessCabinet,
	%[1]sAccessClassType,
	%[1]sAccessProf,
	%[1]sAccessCourseType,
	%[1]sAccessCourse,
	%[1]sAccessPerf,
	%[1]sAccessSkip,
	%[1]sAccessClass,
	%[1]sAccessSession
FROM %[1]s
ORDER BY %[1]sID DESC`,
		js.jobTN,
	)

	stmt, err := js.db.Prepare(query)
	defer stmt.Close()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	rows, err := stmt.Query()
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		return nil, errs.ErrInternalServer
	}

	jobsDB = []*jobschema.JobDB{}
	jobDB := &jobschema.JobDB{}
	for rows.Next() {
		err = rows.Scan(
			&jobDB.JobID,
			&jobDB.JobName,
			&jobDB.JobAccessUser,
			&jobDB.JobAccessJob,
			&jobDB.JobAccessStudent,
			&jobDB.JobAccessUniGroup,
			&jobDB.JobAccessFinhelpCtg,
			&jobDB.JobAccessFinhelpStage,
			&jobDB.JobAccessFinhelpProc,
			&jobDB.JobAccessBuilding,
			&jobDB.JobAccessCabinetType,
			&jobDB.JobAccessCabinet,
			&jobDB.JobAccessClassType,
			&jobDB.JobAccessProf,
			&jobDB.JobAccessCourseType,
			&jobDB.JobAccessCourse,
			&jobDB.JobAccessPerf,
			&jobDB.JobAccessSkip,
			&jobDB.JobAccessClass,
			&jobDB.JobAccessSession,
		)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		err = jobschema.ValidateJobDB(jobDB)
		if err != nil {
			logger.Error.Printf("Internal server error (%s)", err)
			return nil, errs.ErrInternalServer
		}
		jobsDB = append(jobsDB, jobDB)
		jobDB = &jobschema.JobDB{}
	}

	return jobsDB, nil
}
