package jobhandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/jobschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/jobview"
)

func NewJobHandler(ss handler.SessionService, us handler.UserService, js handler.JobService) *JobHandler {
	return &JobHandler{
		SessionService: ss,
		UserService:    us,
		JobService:     js,
	}
}

type JobHandler struct {
	SessionService handler.SessionService
	UserService    handler.UserService
	JobService     handler.JobService
}

func (jh *JobHandler) Job(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, jobview.Job())
}

func (jh *JobHandler) JobPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, jobview.JobPage())
}

func (jh *JobHandler) GetJobs(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &jobschema.JobsGet{}

	err = jobschema.ValidateJobsGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	jobsDB, err := jh.JobService.GetJobs(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, jobview.JobTableRows(jobsDB))
}

func (jh *JobHandler) InsertJob(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &jobschema.JobInsert{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		jh.SessionService.GetSession,
		jh.UserService.GetUser,
		jh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessJob {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	in.JobName = util.GetStringFromForm(r, "JobName")
	in.JobAccessUser, _ = util.GetBoolFromForm(r, "JobAccessUser")
	in.JobAccessJob, _ = util.GetBoolFromForm(r, "JobAccessJob")
	in.JobAccessStudent, _ = util.GetBoolFromForm(r, "JobAccessStudent")
	in.JobAccessGroup, _ = util.GetBoolFromForm(r, "JobAccessGroup")
	in.JobAccessFinhelpCtg, _ = util.GetBoolFromForm(r, "JobAccessFinhelpCtg")
	in.JobAccessFinhelpStage, _ = util.GetBoolFromForm(r, "JobAccessFinhelpStage")
	in.JobAccessFinhelpProc, _ = util.GetBoolFromForm(r, "JobAccessFinhelpProc")
	in.JobAccessBuilding, _ = util.GetBoolFromForm(r, "JobAccessBuilding")
	in.JobAccessCabinetType, _ = util.GetBoolFromForm(r, "JobAccessCabinetType")
	in.JobAccessCabinet, _ = util.GetBoolFromForm(r, "JobAccessCabinet")
	in.JobAccessClassType, _ = util.GetBoolFromForm(r, "JobAccessClassType")
	in.JobAccessProf, _ = util.GetBoolFromForm(r, "JobAccessProf")
	in.JobAccessCourseType, _ = util.GetBoolFromForm(r, "JobAccessCourseType")
	in.JobAccessCourse, _ = util.GetBoolFromForm(r, "JobAccessCourse")
	in.JobAccessPerf, _ = util.GetBoolFromForm(r, "JobAccessPerf")
	in.JobAccessSkip, _ = util.GetBoolFromForm(r, "JobAccessSkip")
	in.JobAccessClass, _ = util.GetBoolFromForm(r, "JobAccessClass")
	err = jobschema.ValidateJobInsert(in)
	if err != nil {
		message = msg.JobNameEmpty
		logger.Error.Print(err.Error())
		return
	}

	jobDB, err := jh.JobService.InsertJob(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.JobExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		}
	}

	util.RenderComponent(r, &out, jobview.JobTableRow(jobDB))
}

func (jh *JobHandler) EditJob(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &jobschema.JobGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		jh.SessionService.GetSession,
		jh.UserService.GetUser,
		jh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessJob {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	in.JobID, err = util.GetIntFromForm(r, "JobID")
	err = jobschema.ValidateJobGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	jobDB, err := jh.JobService.GetJob(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, jobview.JobTableRowEdit(jobDB))
}

func (jh *JobHandler) UpdateJob(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &jobschema.JobUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		jh.SessionService.GetSession,
		jh.UserService.GetUser,
		jh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessJob {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	in.JobID, err = util.GetIntFromForm(r, "JobID")
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	in.JobName = util.GetStringFromForm(r, "JobName")
	in.JobAccessUser, _ = util.GetBoolFromForm(r, "JobAccessUser")
	in.JobAccessJob, _ = util.GetBoolFromForm(r, "JobAccessJob")
	in.JobAccessStudent, _ = util.GetBoolFromForm(r, "JobAccessStudent")
	in.JobAccessGroup, _ = util.GetBoolFromForm(r, "JobAccessGroup")
	in.JobAccessFinhelpCtg, _ = util.GetBoolFromForm(r, "JobAccessFinhelpCtg")
	in.JobAccessFinhelpStage, _ = util.GetBoolFromForm(r, "JobAccessFinhelpStage")
	in.JobAccessFinhelpProc, _ = util.GetBoolFromForm(r, "JobAccessFinhelpProc")
	in.JobAccessBuilding, _ = util.GetBoolFromForm(r, "JobAccessBuilding")
	in.JobAccessCabinetType, _ = util.GetBoolFromForm(r, "JobAccessCabinetType")
	in.JobAccessCabinet, _ = util.GetBoolFromForm(r, "JobAccessCabinet")
	in.JobAccessClassType, _ = util.GetBoolFromForm(r, "JobAccessClassType")
	in.JobAccessProf, _ = util.GetBoolFromForm(r, "JobAccessProf")
	in.JobAccessCourseType, _ = util.GetBoolFromForm(r, "JobAccessCourseType")
	in.JobAccessCourse, _ = util.GetBoolFromForm(r, "JobAccessCourse")
	in.JobAccessPerf, _ = util.GetBoolFromForm(r, "JobAccessPerf")
	in.JobAccessSkip, _ = util.GetBoolFromForm(r, "JobAccessSkip")
	in.JobAccessClass, _ = util.GetBoolFromForm(r, "JobAccessClass")
	err = jobschema.ValidateJobUpdate(in)
	if err != nil {
		message = msg.JobNameEmpty
		logger.Error.Print(err.Error())
		return
	}

	jobDB, err := jh.JobService.UpdateJob(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.JobExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		}
	}

	util.RenderComponent(r, &out, jobview.JobTableRow(jobDB))
}

func (jh *JobHandler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &jobschema.JobDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		jh.SessionService.GetSession,
		jh.UserService.GetUser,
		jh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessJob {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	in.JobID, err = util.GetIntFromForm(r, "JobID")
	err = jobschema.ValidateJobDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	_, err = jh.JobService.DeleteJob(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
}
