package cabinettypehandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/cabinettypeschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/cabinettypeview"
)

func NewCabinetTypeHandler(ss handler.SessionService, us handler.UserService, js handler.JobService, fstages handler.CabinetTypeService) *CabinetTypeHandler {
	return &CabinetTypeHandler{
		SessionService:     ss,
		UserService:        us,
		JobService:         js,
		CabinetTypeService: fstages,
	}
}

type CabinetTypeHandler struct {
	SessionService     handler.SessionService
	UserService        handler.UserService
	JobService         handler.JobService
	CabinetTypeService handler.CabinetTypeService
}

func (cth *CabinetTypeHandler) CabinetType(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, cabinettypeview.CabinetType())
}

func (cth *CabinetTypeHandler) CabinetTypePage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, cabinettypeview.CabinetTypePage())
}

func (cth *CabinetTypeHandler) GetCabinetTypes(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &cabinettypeschema.CabinetTypesGet{}

	err = cabinettypeschema.ValidateCabinetTypesGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	cabinetTypesDB, err := cth.CabinetTypeService.GetCabinetTypes(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, cabinettypeview.CabinetTypeTableRows(cabinetTypesDB))
}

func (cth *CabinetTypeHandler) InsertCabinetType(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &cabinettypeschema.CabinetTypeInsert{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		cth.SessionService.GetSession,
		cth.UserService.GetUser,
		cth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCabinetType {
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
	err = cabinettypeschema.ValidateCabinetTypeInsert(in)
	if err != nil {
		message = msg.CabinetTypeWrong
		logger.Error.Print(err.Error())
		return
	}

	cabinetTypeDB, err := cth.CabinetTypeService.InsertCabinetType(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.CabinetTypeExists
			logger.Error.Print(err.Error())
			return
		}
	}

	util.RenderComponent(r, &out, cabinettypeview.CabinetTypeTableRow(cabinetTypeDB))
}

func (cth *CabinetTypeHandler) EditCabinetType(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &cabinettypeschema.CabinetTypeGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		cth.SessionService.GetSession,
		cth.UserService.GetUser,
		cth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCabinetType {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = cabinettypeschema.ValidateCabinetTypeGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	cabinetTypeDB, err := cth.CabinetTypeService.GetCabinetType(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, cabinettypeview.CabinetTypeTableRowEdit(cabinetTypeDB))
}

func (cth *CabinetTypeHandler) UpdateCabinetType(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &cabinettypeschema.CabinetTypeUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		cth.SessionService.GetSession,
		cth.UserService.GetUser,
		cth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCabinetType {
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
	err = cabinettypeschema.ValidateCabinetTypeUpdate(in)
	if err != nil {
		message = msg.CabinetTypeWrong
		logger.Error.Print(err.Error())
		return
	}

	cabinetTypeDB, err := cth.CabinetTypeService.UpdateCabinetType(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.CabinetTypeExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	util.RenderComponent(r, &out, cabinettypeview.CabinetTypeTableRow(cabinetTypeDB))
}

func (cth *CabinetTypeHandler) DeleteCabinetType(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &cabinettypeschema.CabinetTypeDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		cth.SessionService.GetSession,
		cth.UserService.GetUser,
		cth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCabinetType {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = cabinettypeschema.ValidateCabinetTypeDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	cabinetTypeDB, err := cth.CabinetTypeService.DeleteCabinetType(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, cabinettypeview.CabinetTypeTableRow(cabinetTypeDB))
}
