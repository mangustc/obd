package coursehandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/courseschema"
	"github.com/mangustc/obd/schema/coursetypeschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/courseview"
)

func NewCourseHandler(
	ss handler.SessionService,
	us handler.UserService,
	js handler.JobService,
	cos handler.CourseService,
	cots handler.CourseTypeService,
) *CourseHandler {
	return &CourseHandler{
		SessionService:    ss,
		UserService:       us,
		JobService:        js,
		CourseService:     cos,
		CourseTypeService: cots,
	}
}

type CourseHandler struct {
	SessionService    handler.SessionService
	UserService       handler.UserService
	JobService        handler.JobService
	CourseService     handler.CourseService
	CourseTypeService handler.CourseTypeService
}

func (coh *CourseHandler) Course(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	courseTypesDB, _ := coh.CourseTypeService.GetCourseTypes(&coursetypeschema.CourseTypesGet{})
	courseTypeInputOptions := coursetypeschema.GetCourseTypeInputOptionsFromCourseTypesDB(courseTypesDB)

	util.RenderComponent(r, &out, courseview.Course(courseTypeInputOptions))
}

func (coh *CourseHandler) CoursePage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, courseview.CoursePage())
}

func (coh *CourseHandler) GetCourses(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &courseschema.CoursesGet{}

	err = courseschema.ValidateCoursesGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	coursesDB, err := coh.CourseService.GetCourses(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	courseTypesDB, _ := coh.CourseTypeService.GetCourseTypes(&coursetypeschema.CourseTypesGet{})
	courseTypeInputOptions := coursetypeschema.GetCourseTypeInputOptionsFromCourseTypesDB(courseTypesDB)

	util.RenderComponent(r, &out, courseview.CourseTableRows(coursesDB, courseTypeInputOptions))
}

func (coh *CourseHandler) InsertCourse(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &courseschema.CourseInsert{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		coh.SessionService.GetSession,
		coh.UserService.GetUser,
		coh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCourse {
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
	err = courseschema.ValidateCourseInsert(in)
	if err != nil {
		message = msg.CourseWrong
		logger.Error.Print(err.Error())
		return
	}

	courseDB, err := coh.CourseService.InsertCourse(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.CourseExists
			logger.Error.Print(err.Error())
			return
		}
	}

	courseTypesDB, _ := coh.CourseTypeService.GetCourseTypes(&coursetypeschema.CourseTypesGet{})
	courseTypeInputOptions := coursetypeschema.GetCourseTypeInputOptionsFromCourseTypesDB(courseTypesDB)

	util.RenderComponent(r, &out, courseview.CourseTableRow(courseDB, courseTypeInputOptions))
}

func (coh *CourseHandler) EditCourse(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &courseschema.CourseGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		coh.SessionService.GetSession,
		coh.UserService.GetUser,
		coh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCourse {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = courseschema.ValidateCourseGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	courseDB, err := coh.CourseService.GetCourse(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	courseTypesDB, _ := coh.CourseTypeService.GetCourseTypes(&coursetypeschema.CourseTypesGet{})
	courseTypeInputOptions := coursetypeschema.GetCourseTypeInputOptionsFromCourseTypesDB(courseTypesDB)

	util.RenderComponent(r, &out, courseview.CourseTableRowEdit(courseDB, courseTypeInputOptions))
}

func (coh *CourseHandler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &courseschema.CourseUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		coh.SessionService.GetSession,
		coh.UserService.GetUser,
		coh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCourse {
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
	err = courseschema.ValidateCourseUpdate(in)
	if err != nil {
		message = msg.CourseWrong
		logger.Error.Print(err.Error())
		return
	}

	courseDB, err := coh.CourseService.UpdateCourse(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.CourseExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	courseTypesDB, _ := coh.CourseTypeService.GetCourseTypes(&coursetypeschema.CourseTypesGet{})
	courseTypeInputOptions := coursetypeschema.GetCourseTypeInputOptionsFromCourseTypesDB(courseTypesDB)

	util.RenderComponent(r, &out, courseview.CourseTableRow(courseDB, courseTypeInputOptions))
}

func (coh *CourseHandler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &courseschema.CourseDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		coh.SessionService.GetSession,
		coh.UserService.GetUser,
		coh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCourse {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = courseschema.ValidateCourseDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	courseDB, err := coh.CourseService.DeleteCourse(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	courseTypesDB, _ := coh.CourseTypeService.GetCourseTypes(&coursetypeschema.CourseTypesGet{})
	courseTypeInputOptions := coursetypeschema.GetCourseTypeInputOptionsFromCourseTypesDB(courseTypesDB)

	util.RenderComponent(r, &out, courseview.CourseTableRow(courseDB, courseTypeInputOptions))
}
