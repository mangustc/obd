package classtypehandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/classtypeschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/classtypeview"
)

func NewClassTypeHandler(ss handler.SessionService, us handler.UserService, js handler.JobService, fstages handler.ClassTypeService) *ClassTypeHandler {
	return &ClassTypeHandler{
		SessionService:   ss,
		UserService:      us,
		JobService:       js,
		ClassTypeService: fstages,
	}
}

type ClassTypeHandler struct {
	SessionService   handler.SessionService
	UserService      handler.UserService
	JobService       handler.JobService
	ClassTypeService handler.ClassTypeService
}

func (clth *ClassTypeHandler) ClassType(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, classtypeview.ClassType())
}

func (clth *ClassTypeHandler) ClassTypePage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, classtypeview.ClassTypePage())
}

func (clth *ClassTypeHandler) GetClassTypes(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &classtypeschema.ClassTypesGet{}

	err = classtypeschema.ValidateClassTypesGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	classTypesDB, err := clth.ClassTypeService.GetClassTypes(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, classtypeview.ClassTypeTableRows(classTypesDB))
}

func (clth *ClassTypeHandler) InsertClassType(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &classtypeschema.ClassTypeInsert{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		clth.SessionService.GetSession,
		clth.UserService.GetUser,
		clth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessClassType {
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
	err = classtypeschema.ValidateClassTypeInsert(in)
	if err != nil {
		message = msg.ClassTypeWrong
		logger.Error.Print(err.Error())
		return
	}

	classTypeDB, err := clth.ClassTypeService.InsertClassType(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.ClassTypeExists
			logger.Error.Print(err.Error())
			return
		}
	}

	util.RenderComponent(r, &out, classtypeview.ClassTypeTableRow(classTypeDB))
}

func (clth *ClassTypeHandler) EditClassType(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &classtypeschema.ClassTypeGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		clth.SessionService.GetSession,
		clth.UserService.GetUser,
		clth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessClassType {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = classtypeschema.ValidateClassTypeGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	classTypeDB, err := clth.ClassTypeService.GetClassType(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, classtypeview.ClassTypeTableRowEdit(classTypeDB))
}

func (clth *ClassTypeHandler) UpdateClassType(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &classtypeschema.ClassTypeUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		clth.SessionService.GetSession,
		clth.UserService.GetUser,
		clth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessClassType {
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
	err = classtypeschema.ValidateClassTypeUpdate(in)
	if err != nil {
		message = msg.ClassTypeWrong
		logger.Error.Print(err.Error())
		return
	}

	classTypeDB, err := clth.ClassTypeService.UpdateClassType(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.ClassTypeExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	util.RenderComponent(r, &out, classtypeview.ClassTypeTableRow(classTypeDB))
}

func (clth *ClassTypeHandler) DeleteClassType(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &classtypeschema.ClassTypeDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		clth.SessionService.GetSession,
		clth.UserService.GetUser,
		clth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessClassType {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = classtypeschema.ValidateClassTypeDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	classTypeDB, err := clth.ClassTypeService.DeleteClassType(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, classtypeview.ClassTypeTableRow(classTypeDB))
}
