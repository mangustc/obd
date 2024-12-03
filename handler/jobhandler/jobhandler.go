package jobhandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/schema/jobschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view"
	"github.com/mangustc/obd/view/jobview"
)

func NewJobHandler(js handler.JobService) *JobHandler {
	return &JobHandler{
		JobService: js,
	}
}

type JobHandler struct {
	JobService handler.JobService
}

func (jh *JobHandler) Job(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	defer util.RespondHTTP(w, &code, &out)

	util.RenderComponent(r, &out, jobview.Job())
}

func (jh *JobHandler) JobPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	defer util.RespondHTTP(w, &code, &out)

	util.RenderComponent(r, &out, jobview.JobPage())
}

func (jh *JobHandler) GetJobs(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	in := &jobschema.JobsGet{}
	defer util.RespondHTTP(w, &code, &out)

	err = jobschema.ValidateJobsGet(in)
	if err != nil {
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	jobsDB, err := jh.JobService.GetJobs(in)
	if err != nil {
		if err == errs.ErrInternalServer {
			code = http.StatusInternalServerError
			util.RenderComponent(r, &out, view.ErrorIndex(code))
			return
		}
	}

	util.RenderComponent(r, &out, jobview.JobTableRows(jobsDB))
}

func (jh *JobHandler) InsertJob(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	in := &jobschema.JobInsert{}
	defer util.RespondHTTP(w, &code, &out)

	in.JobName = util.GetStringFromForm(r, "JobName")
	in.JobAccessUser, _ = util.GetBoolFromForm(r, "JobAccessUser")
	in.JobAccessJob, _ = util.GetBoolFromForm(r, "JobAccessJob")
	in.JobAccessStudent, _ = util.GetBoolFromForm(r, "JobAccessStudent")
	in.JobAccessUniGroup, _ = util.GetBoolFromForm(r, "JobAccessUniGroup")
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
	in.JobAccessSession, _ = util.GetBoolFromForm(r, "JobAccessSession")
	err = jobschema.ValidateJobInsert(in)
	if err != nil {
		code = http.StatusUnprocessableEntity
		// TODO: Handle error somehow (?)
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	jobDB, err := jh.JobService.InsertJob(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			code = http.StatusInternalServerError
			util.RenderComponent(r, &out, view.ErrorIndex(code))
			return
		} else {
			code = http.StatusUnprocessableEntity
			// TODO: Handle error somehow (?)
			util.RenderComponent(r, &out, view.ErrorIndex(code))
			return

		}
	}

	// util.RenderComponent(r, &out, jobview.JobAddForm())
	util.RenderComponent(r, &out, jobview.JobTableRow(jobDB))
}

func (jh *JobHandler) EditJob(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	in := &jobschema.JobGet{}
	defer util.RespondHTTP(w, &code, &out)

	in.JobID, err = util.GetIntFromForm(r, "JobID")
	err = jobschema.ValidateJobGet(in)
	if err != nil {
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	jobDB, err := jh.JobService.GetJob(in)
	if err != nil {
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	util.RenderComponent(r, &out, jobview.JobTableRowEdit(jobDB))
}

func (jh *JobHandler) UpdateJob(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	in := &jobschema.JobUpdate{}
	defer util.RespondHTTP(w, &code, &out)

	in.JobID, err = util.GetIntFromForm(r, "JobID")
	if err != nil {
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}
	in.JobName = util.GetStringFromForm(r, "JobName")
	in.JobAccessUser, _ = util.GetBoolFromForm(r, "JobAccessUser")
	in.JobAccessJob, _ = util.GetBoolFromForm(r, "JobAccessJob")
	in.JobAccessStudent, _ = util.GetBoolFromForm(r, "JobAccessStudent")
	in.JobAccessUniGroup, _ = util.GetBoolFromForm(r, "JobAccessUniGroup")
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
	in.JobAccessSession, _ = util.GetBoolFromForm(r, "JobAccessSession")
	err = jobschema.ValidateJobUpdate(in)
	if err != nil {
		code = http.StatusUnprocessableEntity
		// TODO: Handle error somehow (?)
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	jobDB, err := jh.JobService.UpdateJob(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			code = http.StatusInternalServerError
			util.RenderComponent(r, &out, view.ErrorIndex(code))
			return
		} else {
			code = http.StatusUnprocessableEntity
			// TODO: Handle error somehow (?)
			util.RenderComponent(r, &out, view.ErrorIndex(code))
			return

		}
	}

	util.RenderComponent(r, &out, jobview.JobTableRow(jobDB))
}

func (jh *JobHandler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	in := &jobschema.JobDelete{}
	defer util.RespondHTTP(w, &code, &out)

	in.JobID, err = util.GetIntFromForm(r, "JobID")
	err = jobschema.ValidateJobDelete(in)
	if err != nil {
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	_, err = jh.JobService.DeleteJob(in)
	if err != nil {
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	// util.RenderComponent(r, &out, jobview.JobTableRowEdit(jobDB))
}
