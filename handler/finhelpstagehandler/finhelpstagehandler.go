package finhelpstagehandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/finhelpstageschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/finhelpstageview"
)

func NewFinhelpStageHandler(ss handler.SessionService, us handler.UserService, js handler.JobService, fstages handler.FinhelpStageService) *FinhelpStageHandler {
	return &FinhelpStageHandler{
		SessionService:      ss,
		UserService:         us,
		JobService:          js,
		FinhelpStageService: fstages,
	}
}

type FinhelpStageHandler struct {
	SessionService      handler.SessionService
	UserService         handler.UserService
	JobService          handler.JobService
	FinhelpStageService handler.FinhelpStageService
}

func (fsth *FinhelpStageHandler) FinhelpStage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, finhelpstageview.FinhelpStage())
}

func (fsth *FinhelpStageHandler) FinhelpStagePage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, finhelpstageview.FinhelpStagePage())
}

func (fsth *FinhelpStageHandler) GetFinhelpStages(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpstageschema.FinhelpStagesGet{}

	err = finhelpstageschema.ValidateFinhelpStagesGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	finhelpStagesDB, err := fsth.FinhelpStageService.GetFinhelpStages(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	println(util.PrettyPrint(finhelpStagesDB))

	util.RenderComponent(r, &out, finhelpstageview.FinhelpStageTableRows(finhelpStagesDB))
}

func (fsth *FinhelpStageHandler) InsertFinhelpStage(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpstageschema.FinhelpStageInsert{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		fsth.SessionService.GetSession,
		fsth.UserService.GetUser,
		fsth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpStage {
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
	err = finhelpstageschema.ValidateFinhelpStageInsert(in)
	if err != nil {
		message = msg.FinhelpStageWrong
		logger.Error.Print(err.Error())
		return
	}

	finhelpStageDB, err := fsth.FinhelpStageService.InsertFinhelpStage(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.FinhelpStageExists
			logger.Error.Print(err.Error())
			return
		}
	}

	util.RenderComponent(r, &out, finhelpstageview.FinhelpStageTableRow(finhelpStageDB))
}

func (fsth *FinhelpStageHandler) EditFinhelpStage(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpstageschema.FinhelpStageGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		fsth.SessionService.GetSession,
		fsth.UserService.GetUser,
		fsth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpStage {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = finhelpstageschema.ValidateFinhelpStageGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	finhelpStageDB, err := fsth.FinhelpStageService.GetFinhelpStage(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, finhelpstageview.FinhelpStageTableRowEdit(finhelpStageDB))
}

func (fsth *FinhelpStageHandler) UpdateFinhelpStage(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpstageschema.FinhelpStageUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		fsth.SessionService.GetSession,
		fsth.UserService.GetUser,
		fsth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpStage {
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
	err = finhelpstageschema.ValidateFinhelpStageUpdate(in)
	if err != nil {
		message = msg.FinhelpStageWrong
		logger.Error.Print(err.Error())
		return
	}

	finhelpStageDB, err := fsth.FinhelpStageService.UpdateFinhelpStage(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.FinhelpStageExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	util.RenderComponent(r, &out, finhelpstageview.FinhelpStageTableRow(finhelpStageDB))
}

func (fsth *FinhelpStageHandler) DeleteFinhelpStage(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpstageschema.FinhelpStageDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		fsth.SessionService.GetSession,
		fsth.UserService.GetUser,
		fsth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpStage {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = finhelpstageschema.ValidateFinhelpStageDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	finhelpStageDB, err := fsth.FinhelpStageService.DeleteFinhelpStage(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, finhelpstageview.FinhelpStageTableRow(finhelpStageDB))
}
