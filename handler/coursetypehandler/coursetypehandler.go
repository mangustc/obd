package coursetypehandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/coursetypeschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/coursetypeview"
)

func NewCourseTypeHandler(ss handler.SessionService, us handler.UserService, js handler.JobService, fstages handler.CourseTypeService) *CourseTypeHandler {
	return &CourseTypeHandler{
		SessionService:    ss,
		UserService:       us,
		JobService:        js,
		CourseTypeService: fstages,
	}
}

type CourseTypeHandler struct {
	SessionService    handler.SessionService
	UserService       handler.UserService
	JobService        handler.JobService
	CourseTypeService handler.CourseTypeService
}

func (coth *CourseTypeHandler) CourseType(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, coursetypeview.CourseType())
}

func (coth *CourseTypeHandler) CourseTypePage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, coursetypeview.CourseTypePage())
}

func (coth *CourseTypeHandler) GetCourseTypes(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &coursetypeschema.CourseTypesGet{}

	err = coursetypeschema.ValidateCourseTypesGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	courseTypesDB, err := coth.CourseTypeService.GetCourseTypes(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, coursetypeview.CourseTypeTableRows(courseTypesDB))
}

func (coth *CourseTypeHandler) InsertCourseType(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &coursetypeschema.CourseTypeInsert{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		coth.SessionService.GetSession,
		coth.UserService.GetUser,
		coth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCourseType {
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
	err = coursetypeschema.ValidateCourseTypeInsert(in)
	if err != nil {
		message = msg.CourseTypeWrong
		logger.Error.Print(err.Error())
		return
	}

	courseTypeDB, err := coth.CourseTypeService.InsertCourseType(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.CourseTypeExists
			logger.Error.Print(err.Error())
			return
		}
	}

	util.RenderComponent(r, &out, coursetypeview.CourseTypeTableRow(courseTypeDB))
}

func (coth *CourseTypeHandler) EditCourseType(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &coursetypeschema.CourseTypeGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		coth.SessionService.GetSession,
		coth.UserService.GetUser,
		coth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCourseType {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = coursetypeschema.ValidateCourseTypeGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	courseTypeDB, err := coth.CourseTypeService.GetCourseType(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, coursetypeview.CourseTypeTableRowEdit(courseTypeDB))
}

func (coth *CourseTypeHandler) UpdateCourseType(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &coursetypeschema.CourseTypeUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		coth.SessionService.GetSession,
		coth.UserService.GetUser,
		coth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCourseType {
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
	err = coursetypeschema.ValidateCourseTypeUpdate(in)
	if err != nil {
		message = msg.CourseTypeWrong
		logger.Error.Print(err.Error())
		return
	}

	courseTypeDB, err := coth.CourseTypeService.UpdateCourseType(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.CourseTypeExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	util.RenderComponent(r, &out, coursetypeview.CourseTypeTableRow(courseTypeDB))
}

func (coth *CourseTypeHandler) DeleteCourseType(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &coursetypeschema.CourseTypeDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		coth.SessionService.GetSession,
		coth.UserService.GetUser,
		coth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCourseType {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = coursetypeschema.ValidateCourseTypeDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	courseTypeDB, err := coth.CourseTypeService.DeleteCourseType(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, coursetypeview.CourseTypeTableRow(courseTypeDB))
}
