package jobhandler

import (
	"net/http"
	"strconv"

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

	in.JobName = r.Form.Get("JobName")
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

	in.JobID, err = strconv.Atoi(r.Form.Get("JobID"))
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

	in.JobID, err = strconv.Atoi(r.Form.Get("JobID"))
	if err != nil {
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}
	in.JobName = r.Form.Get("JobName")
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

	in.JobID, err = strconv.Atoi(r.Form.Get("JobID"))
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
