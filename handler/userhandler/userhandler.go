package userhandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/jobschema"
	"github.com/mangustc/obd/schema/userschema"
	"github.com/mangustc/obd/util"
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
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	jobsGet := &jobschema.JobsGet{}
	jobsDB, _ := uh.JobService.GetJobs(jobsGet)
	jobInputOptions := jobschema.GetJobInputOptionsFromJobsDB(jobsDB)

	util.RenderComponent(r, &out, userview.User(jobInputOptions))
}

func (uh *UserHandler) UserPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, userview.UserPage())
}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &userschema.UsersGet{}

	err = userschema.ValidateUsersGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	usersDB, err := uh.UserService.GetUsers(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	jobsGet := &jobschema.JobsGet{}
	jobsDB, _ := uh.JobService.GetJobs(jobsGet)
	jobInputOptions := jobschema.GetJobInputOptionsFromJobsDB(jobsDB)

	util.RenderComponent(r, &out, userview.UserTableRows(usersDB, jobInputOptions))
}

func (uh *UserHandler) InsertUser(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &userschema.UserInsert{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessUser {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	in.UserLastname = util.GetStringFromForm(r, "UserLastname")
	in.UserFirstname = util.GetStringFromForm(r, "UserFirstname")
	in.UserMiddlename = util.GetStringFromForm(r, "UserMiddlename")
	in.UserPassword = util.GetStringFromForm(r, "UserPassword")
	in.JobID, err = util.GetIntFromForm(r, "JobID")
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	err = userschema.ValidateUserInsert(in)
	if err != nil {
		message = msg.UserWrong
		logger.Error.Print(err.Error())
		return
	}

	userDB, err := uh.UserService.InsertUser(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.UserExists
			logger.Error.Print(err.Error())
			return
		}
	}

	jobsGet := &jobschema.JobsGet{}
	jobsDB, _ := uh.JobService.GetJobs(jobsGet)
	jobInputOptions := jobschema.GetJobInputOptionsFromJobsDB(jobsDB)

	util.RenderComponent(r, &out, userview.UserTableRow(userDB, jobInputOptions))
}

func (uh *UserHandler) EditUser(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &userschema.UserGet{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessUser {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	in.UserID, err = util.GetIntFromForm(r, "UserID")
	err = userschema.ValidateUserGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	userDB, err := uh.UserService.GetUser(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
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
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &userschema.UserUpdate{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessUser {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	in.UserID, err = util.GetIntFromForm(r, "UserID")
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	in.UserLastname = util.GetStringFromForm(r, "UserLastname")
	in.UserFirstname = util.GetStringFromForm(r, "UserFirstname")
	in.UserMiddlename = util.GetStringFromForm(r, "UserMiddlename")
	in.UserPassword = util.GetStringFromForm(r, "UserPassword")
	in.JobID, err = util.GetIntFromForm(r, "JobID")
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	err = userschema.ValidateUserUpdate(in)
	if err != nil {
		message = msg.UserWrong
		logger.Error.Print(err.Error())
		return
	}

	userDB, err := uh.UserService.UpdateUser(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.UserExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
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
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &userschema.UserDelete{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessUser {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	in.UserID, err = util.GetIntFromForm(r, "UserID")
	err = userschema.ValidateUserDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	userDB, err := uh.UserService.DeleteUser(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	jobsGet := &jobschema.JobsGet{}
	jobsDB, _ := uh.JobService.GetJobs(jobsGet)
	jobInputOptions := jobschema.GetJobInputOptionsFromJobsDB(jobsDB)

	util.RenderComponent(r, &out, userview.UserTableRow(userDB, jobInputOptions))
}
