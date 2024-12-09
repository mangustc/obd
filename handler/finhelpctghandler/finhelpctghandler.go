package finhelpctghandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/finhelpctgschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/finhelpctgview"
)

func NewFinhelpCtgHandler(ss handler.SessionService, us handler.UserService, js handler.JobService, fctgs handler.FinhelpCtgService) *FinhelpCtgHandler {
	return &FinhelpCtgHandler{
		SessionService:    ss,
		UserService:       us,
		JobService:        js,
		FinhelpCtgService: fctgs,
	}
}

type FinhelpCtgHandler struct {
	SessionService    handler.SessionService
	UserService       handler.UserService
	JobService        handler.JobService
	FinhelpCtgService handler.FinhelpCtgService
}

func (fctgh *FinhelpCtgHandler) FinhelpCtg(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, finhelpctgview.FinhelpCtg())
}

func (fctgh *FinhelpCtgHandler) FinhelpCtgPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, finhelpctgview.FinhelpCtgPage())
}

func (fctgh *FinhelpCtgHandler) GetFinhelpCtgs(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpctgschema.FinhelpCtgsGet{}

	err = finhelpctgschema.ValidateFinhelpCtgsGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	finhelpCtgsDB, err := fctgh.FinhelpCtgService.GetFinhelpCtgs(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	println(util.PrettyPrint(finhelpCtgsDB))

	util.RenderComponent(r, &out, finhelpctgview.FinhelpCtgTableRows(finhelpCtgsDB))
}

func (fctgh *FinhelpCtgHandler) InsertFinhelpCtg(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpctgschema.FinhelpCtgInsert{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		fctgh.SessionService.GetSession,
		fctgh.UserService.GetUser,
		fctgh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpCtg {
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
	err = finhelpctgschema.ValidateFinhelpCtgInsert(in)
	if err != nil {
		message = msg.FinhelpCtgWrong
		logger.Error.Print(err.Error())
		return
	}

	finhelpCtgDB, err := fctgh.FinhelpCtgService.InsertFinhelpCtg(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.FinhelpCtgExists
			logger.Error.Print(err.Error())
			return
		}
	}

	util.RenderComponent(r, &out, finhelpctgview.FinhelpCtgTableRow(finhelpCtgDB))
}

func (fctgh *FinhelpCtgHandler) EditFinhelpCtg(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpctgschema.FinhelpCtgGet{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		fctgh.SessionService.GetSession,
		fctgh.UserService.GetUser,
		fctgh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpCtg {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = finhelpctgschema.ValidateFinhelpCtgGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	finhelpCtgDB, err := fctgh.FinhelpCtgService.GetFinhelpCtg(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, finhelpctgview.FinhelpCtgTableRowEdit(finhelpCtgDB))
}

func (fctgh *FinhelpCtgHandler) UpdateFinhelpCtg(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpctgschema.FinhelpCtgUpdate{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		fctgh.SessionService.GetSession,
		fctgh.UserService.GetUser,
		fctgh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpCtg {
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
	err = finhelpctgschema.ValidateFinhelpCtgUpdate(in)
	if err != nil {
		message = msg.FinhelpCtgWrong
		logger.Error.Print(err.Error())
		return
	}

	finhelpCtgDB, err := fctgh.FinhelpCtgService.UpdateFinhelpCtg(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.FinhelpCtgExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	util.RenderComponent(r, &out, finhelpctgview.FinhelpCtgTableRow(finhelpCtgDB))
}

func (fctgh *FinhelpCtgHandler) DeleteFinhelpCtg(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpctgschema.FinhelpCtgDelete{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		fctgh.SessionService.GetSession,
		fctgh.UserService.GetUser,
		fctgh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpCtg {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = finhelpctgschema.ValidateFinhelpCtgDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	finhelpCtgDB, err := fctgh.FinhelpCtgService.DeleteFinhelpCtg(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, finhelpctgview.FinhelpCtgTableRow(finhelpCtgDB))
}
