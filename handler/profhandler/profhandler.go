package profhandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/profschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/profview"
)

func NewProfHandler(ss handler.SessionService, us handler.UserService, js handler.JobService, prs handler.ProfService) *ProfHandler {
	return &ProfHandler{
		SessionService: ss,
		UserService:    us,
		JobService:     js,
		ProfService:    prs,
	}
}

type ProfHandler struct {
	SessionService handler.SessionService
	UserService    handler.UserService
	JobService     handler.JobService
	ProfService    handler.ProfService
}

func (cth *ProfHandler) Prof(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, profview.Prof())
}

func (cth *ProfHandler) ProfPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, profview.ProfPage())
}

func (cth *ProfHandler) GetProfs(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &profschema.ProfsGet{}

	err = profschema.ValidateProfsGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	profsDB, err := cth.ProfService.GetProfs(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, profview.ProfTableRows(profsDB))
}

func (cth *ProfHandler) InsertProf(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &profschema.ProfInsert{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		cth.SessionService.GetSession,
		cth.UserService.GetUser,
		cth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessProf {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	err = profschema.ValidateProfInsert(in)
	if err != nil {
		message = msg.ProfWrong
		logger.Error.Print(err.Error())
		return
	}

	profDB, err := cth.ProfService.InsertProf(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.ProfExists
			logger.Error.Print(err.Error())
			return
		}
	}

	util.RenderComponent(r, &out, profview.ProfTableRow(profDB))
}

func (cth *ProfHandler) EditProf(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &profschema.ProfGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		cth.SessionService.GetSession,
		cth.UserService.GetUser,
		cth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessProf {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = profschema.ValidateProfGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	profDB, err := cth.ProfService.GetProf(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, profview.ProfTableRowEdit(profDB))
}

func (cth *ProfHandler) UpdateProf(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &profschema.ProfUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		cth.SessionService.GetSession,
		cth.UserService.GetUser,
		cth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessProf {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	err = profschema.ValidateProfUpdate(in)
	if err != nil {
		message = msg.ProfWrong
		logger.Error.Print(err.Error())
		return
	}

	profDB, err := cth.ProfService.UpdateProf(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.ProfExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	util.RenderComponent(r, &out, profview.ProfTableRow(profDB))
}

func (cth *ProfHandler) DeleteProf(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &profschema.ProfDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		cth.SessionService.GetSession,
		cth.UserService.GetUser,
		cth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessProf {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = profschema.ValidateProfDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	profDB, err := cth.ProfService.DeleteProf(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, profview.ProfTableRow(profDB))
}
