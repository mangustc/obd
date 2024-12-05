package userhandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/schema/jobschema"
	"github.com/mangustc/obd/schema/userschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view"
	"github.com/mangustc/obd/view/userview"
)

func NewUserHandler(ss handler.SessionService, us handler.UserService, js handler.JobService) *UserHandler {
	return &UserHandler{
		SessionService: ss,
		UserService:    us,
		JobService:     js,
	}
}

type UserHandler struct {
	SessionService handler.SessionService
	UserService    handler.UserService
	JobService     handler.JobService
}

func (uh *UserHandler) User(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	defer util.RespondHTTP(w, &code, &out)

	jobsGet := &jobschema.JobsGet{}
	jobsDB, _ := uh.JobService.GetJobs(jobsGet)
	jobInputOptions := jobschema.GetJobInputOptionsFromJobsDB(jobsDB)

	util.RenderComponent(r, &out, userview.User(jobInputOptions))
}

func (uh *UserHandler) UserPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	defer util.RespondHTTP(w, &code, &out)

	util.RenderComponent(r, &out, userview.UserPage())
}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	in := &userschema.UsersGet{}
	defer util.RespondHTTP(w, &code, &out)

	err = userschema.ValidateUsersGet(in)
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	usersDB, err := uh.UserService.GetUsers(in)
	if err != nil {
		if err == errs.ErrInternalServer {
			logger.Error.Printf("Internal server error (%s)", err)
			code = http.StatusInternalServerError
			util.RenderComponent(r, &out, view.ErrorIndex(code))
			return
		}
	}

	jobsGet := &jobschema.JobsGet{}
	jobsDB, _ := uh.JobService.GetJobs(jobsGet)
	jobInputOptions := jobschema.GetJobInputOptionsFromJobsDB(jobsDB)

	util.RenderComponent(r, &out, userview.UserTableRows(usersDB, jobInputOptions))
}

func (uh *UserHandler) InsertUser(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	in := &userschema.UserInsert{}
	defer util.RespondHTTP(w, &code, &out)

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessUser {
		code, str := util.GetCodeByErr(errs.ErrUnauthorized)
		logger.Error.Print(str)
		// TODO: Handle error somehow (?)
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	in.UserLastname = util.GetStringFromForm(r, "UserLastname")
	in.UserFirstname = util.GetStringFromForm(r, "UserFirstname")
	in.UserMiddlename = util.GetStringFromForm(r, "UserMiddlename")
	in.UserPassword = util.GetStringFromForm(r, "UserPassword")
	in.JobID, err = util.GetIntFromForm(r, "JobID")
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}
	err = userschema.ValidateUserInsert(in)
	if err != nil {
		logger.Error.Printf("Unprocessable Entity (%s)", err)
		code = http.StatusUnprocessableEntity
		// TODO: Handle error somehow (?)
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	userDB, err := uh.UserService.InsertUser(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			logger.Error.Printf("Internal server error (%s)", err)
			code = http.StatusInternalServerError
			util.RenderComponent(r, &out, view.ErrorIndex(code))
			return
		} else {
			logger.Error.Printf("Unprocessable Entity (%s)", err)
			code = http.StatusUnprocessableEntity
			// TODO: Handle error somehow (?)
			util.RenderComponent(r, &out, view.ErrorIndex(code))
			return

		}
	}

	jobsGet := &jobschema.JobsGet{}
	jobsDB, _ := uh.JobService.GetJobs(jobsGet)
	jobInputOptions := jobschema.GetJobInputOptionsFromJobsDB(jobsDB)

	// util.RenderComponent(r, &out, userview.UserAddForm())
	util.RenderComponent(r, &out, userview.UserTableRow(userDB, jobInputOptions))
}

func (uh *UserHandler) EditUser(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	in := &userschema.UserGet{}
	defer util.RespondHTTP(w, &code, &out)

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessUser {
		code, str := util.GetCodeByErr(errs.ErrUnauthorized)
		logger.Error.Print(str)
		// TODO: Handle error somehow (?)
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	in.UserID, err = util.GetIntFromForm(r, "UserID")
	err = userschema.ValidateUserGet(in)
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	userDB, err := uh.UserService.GetUser(in)
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	jobsGet := &jobschema.JobsGet{}
	jobsDB, _ := uh.JobService.GetJobs(jobsGet)
	jobInputOptions := jobschema.GetJobInputOptionsFromJobsDB(jobsDB)

	util.RenderComponent(r, &out, userview.UserTableRowEdit(userDB, jobInputOptions))
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	in := &userschema.UserUpdate{}
	defer util.RespondHTTP(w, &code, &out)

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessUser {
		code, str := util.GetCodeByErr(errs.ErrUnauthorized)
		logger.Error.Print(str)
		// TODO: Handle error somehow (?)
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	in.UserID, err = util.GetIntFromForm(r, "UserID")
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}
	in.UserLastname = util.GetStringFromForm(r, "UserLastname")
	in.UserFirstname = util.GetStringFromForm(r, "UserFirstname")
	in.UserMiddlename = util.GetStringFromForm(r, "UserMiddlename")
	in.UserPassword = util.GetStringFromForm(r, "UserPassword")
	in.JobID, err = util.GetIntFromForm(r, "JobID")
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}
	err = userschema.ValidateUserUpdate(in)
	if err != nil {
		logger.Error.Printf("Unprocessable Entity (%s)", err)
		code = http.StatusUnprocessableEntity
		// TODO: Handle error somehow (?)
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	userDB, err := uh.UserService.UpdateUser(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			logger.Error.Printf("Internal server error (%s)", err)
			code = http.StatusInternalServerError
			util.RenderComponent(r, &out, view.ErrorIndex(code))
			return
		} else {
			logger.Error.Printf("Unprocessable Entity (%s)", err)
			code = http.StatusUnprocessableEntity
			// TODO: Handle error somehow (?)
			util.RenderComponent(r, &out, view.ErrorIndex(code))
			return

		}
	}

	jobsGet := &jobschema.JobsGet{}
	jobsDB, _ := uh.JobService.GetJobs(jobsGet)
	jobInputOptions := jobschema.GetJobInputOptionsFromJobsDB(jobsDB)

	util.RenderComponent(r, &out, userview.UserTableRow(userDB, jobInputOptions))
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	in := &userschema.UserDelete{}
	defer util.RespondHTTP(w, &code, &out)

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessUser {
		code, str := util.GetCodeByErr(errs.ErrUnauthorized)
		logger.Error.Print(str)
		// TODO: Handle error somehow (?)
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	in.UserID, err = util.GetIntFromForm(r, "UserID")
	err = userschema.ValidateUserDelete(in)
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	userDB, err := uh.UserService.DeleteUser(in)
	if err != nil {
		logger.Error.Printf("Internal server error (%s)", err)
		code = http.StatusInternalServerError
		util.RenderComponent(r, &out, view.ErrorIndex(code))
		return
	}

	jobsGet := &jobschema.JobsGet{}
	jobsDB, _ := uh.JobService.GetJobs(jobsGet)
	jobInputOptions := jobschema.GetJobInputOptionsFromJobsDB(jobsDB)

	util.RenderComponent(r, &out, userview.UserTableRow(userDB, jobInputOptions))
}
