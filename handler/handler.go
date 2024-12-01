package handler

import (
	"github.com/mangustc/obd/schema/jobschema"
)

type JobService interface {
	InsertJob(data *jobschema.JobInsert) (jobDB *jobschema.JobDB, err error)
	UpdateJob(data *jobschema.JobUpdate) (jobDB *jobschema.JobDB, err error)
	DeleteJob(data *jobschema.JobDelete) (jobDB *jobschema.JobDB, err error)
	GetJob(data *jobschema.JobGet) (jobDB *jobschema.JobDB, err error)
	GetJobs(data *jobschema.JobsGet) (jobsDB []*jobschema.JobDB, err error)
}
