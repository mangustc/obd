package handler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/jobschema"
	"github.com/mangustc/obd/schema/sessionschema"
	"github.com/mangustc/obd/schema/userschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view"
)

type (
	JobService interface {
		InsertJob(data *jobschema.JobInsert) (jobDB *jobschema.JobDB, err error)
		UpdateJob(data *jobschema.JobUpdate) (jobDB *jobschema.JobDB, err error)
		DeleteJob(data *jobschema.JobDelete) (jobDB *jobschema.JobDB, err error)
		GetJob(data *jobschema.JobGet) (jobDB *jobschema.JobDB, err error)
		GetJobs(data *jobschema.JobsGet) (jobsDB []*jobschema.JobDB, err error)
	}
	UserService interface {
		InsertUser(data *userschema.UserInsert) (userDB *userschema.UserDB, err error)
		UpdateUser(data *userschema.UserUpdate) (userDB *userschema.UserDB, err error)
		DeleteUser(data *userschema.UserDelete) (userDB *userschema.UserDB, err error)
		GetUser(data *userschema.UserGet) (userDB *userschema.UserDB, err error)
		GetUsers(data *userschema.UsersGet) (usersDB []*userschema.UserDB, err error)
	}
	SessionService interface {
		InsertSession(data *sessionschema.SessionInsert) (sessionDB *sessionschema.SessionDB, err error)
		DeleteSession(data *sessionschema.SessionDelete) (sessionDB *sessionschema.SessionDB, err error)
		GetSession(data *sessionschema.SessionGet) (sessionDB *sessionschema.SessionDB, err error)
		GetSessions(data *sessionschema.SessionsGet) (sessionsDB []*sessionschema.SessionDB, err error)
	}
)

func NewDefaultHandler(ss SessionService, us UserService, js JobService) *DefaultHandler {
	return &DefaultHandler{
		SessionService: ss,
		UserService:    us,
		JobService:     js,
	}
}

type DefaultHandler struct {
	SessionService SessionService
	UserService    UserService
	JobService     JobService
}

func (dh *DefaultHandler) Default(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	defer util.RespondHTTP(w, &code, &out)

	util.RenderComponent(r, &out, view.Layout("OBD"))
}

func (dh *DefaultHandler) Navigation(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	defer util.RespondHTTP(w, &code, &out)

	jobDB, err := util.GetJobBySessionCookie(
		w, r,
		dh.SessionService.GetSession,
		dh.UserService.GetUser,
		dh.JobService.GetJob,
	)
	if err != nil {
		if err == errs.ErrNotFound {
			jobDB = nil
		} else {
			code, str := util.GetCodeByErr(err)
			logger.Error.Print(str)
			// TODO: Handle error somehow (?)
			util.RenderComponent(r, &out, view.ErrorIndex(code))
			return
		}
	}

	util.RenderComponent(r, &out, view.NavigationByJobDB(jobDB))
}
