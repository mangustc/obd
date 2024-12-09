package studenthandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/groupschema"
	"github.com/mangustc/obd/schema/studentschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/studentview"
)

func NewStudentHandler(ss handler.SessionService, us handler.UserService, js handler.JobService, grs handler.GroupService, sts handler.StudentService) *StudentHandler {
	return &StudentHandler{
		SessionService: ss,
		UserService:    us,
		JobService:     js,
		GroupService:   grs,
		StudentService: sts,
	}
}

type StudentHandler struct {
	SessionService handler.SessionService
	UserService    handler.UserService
	JobService     handler.JobService
	GroupService   handler.GroupService
	StudentService handler.StudentService
}

func (uh *StudentHandler) Student(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	groupsGet := &groupschema.GroupsGet{}
	groupsDB, _ := uh.GroupService.GetGroups(groupsGet)
	groupInputOptions := groupschema.GetGroupInputOptionsFromGroupsDB(groupsDB)

	util.RenderComponent(r, &out, studentview.Student(groupInputOptions))
}

func (uh *StudentHandler) StudentPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, studentview.StudentPage())
}

func (uh *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &studentschema.StudentsGet{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessStudent {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = studentschema.ValidateStudentsGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	studentsDB, err := uh.StudentService.GetStudents(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	groupsGet := &groupschema.GroupsGet{}
	groupsDB, _ := uh.GroupService.GetGroups(groupsGet)
	groupInputOptions := groupschema.GetGroupInputOptionsFromGroupsDB(groupsDB)

	util.RenderComponent(r, &out, studentview.StudentTableRows(studentsDB, groupInputOptions))
}

func (uh *StudentHandler) InsertStudent(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &studentschema.StudentInsert{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessStudent {
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
	err = studentschema.ValidateStudentInsert(in)
	if err != nil {
		message = msg.StudentWrong
		logger.Error.Print(err.Error())
		return
	}

	studentDB, err := uh.StudentService.InsertStudent(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.StudentExists
			logger.Error.Print(err.Error())
			return
		}
	}

	groupsGet := &groupschema.GroupsGet{}
	groupsDB, _ := uh.GroupService.GetGroups(groupsGet)
	groupInputOptions := groupschema.GetGroupInputOptionsFromGroupsDB(groupsDB)

	util.RenderComponent(r, &out, studentview.StudentTableRow(studentDB, groupInputOptions))
}

func (uh *StudentHandler) EditStudent(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &studentschema.StudentGet{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessStudent {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = studentschema.ValidateStudentGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	studentDB, err := uh.StudentService.GetStudent(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	groupsGet := &groupschema.GroupsGet{}
	groupsDB, _ := uh.GroupService.GetGroups(groupsGet)
	groupInputOptions := groupschema.GetGroupInputOptionsFromGroupsDB(groupsDB)

	util.RenderComponent(r, &out, studentview.StudentTableRowEdit(studentDB, groupInputOptions))
}

func (uh *StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &studentschema.StudentUpdate{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessStudent {
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
	err = studentschema.ValidateStudentUpdate(in)
	if err != nil {
		message = msg.StudentWrong
		logger.Error.Print(err.Error())
		return
	}

	studentDB, err := uh.StudentService.UpdateStudent(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.StudentExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	groupsGet := &groupschema.GroupsGet{}
	groupsDB, _ := uh.GroupService.GetGroups(groupsGet)
	groupInputOptions := groupschema.GetGroupInputOptionsFromGroupsDB(groupsDB)

	util.RenderComponent(r, &out, studentview.StudentTableRow(studentDB, groupInputOptions))
}

func (uh *StudentHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &studentschema.StudentDelete{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		uh.SessionService.GetSession,
		uh.UserService.GetUser,
		uh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessStudent {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = studentschema.ValidateStudentDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	studentDB, err := uh.StudentService.DeleteStudent(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	groupsGet := &groupschema.GroupsGet{}
	groupsDB, _ := uh.GroupService.GetGroups(groupsGet)
	groupInputOptions := groupschema.GetGroupInputOptionsFromGroupsDB(groupsDB)

	util.RenderComponent(r, &out, studentview.StudentTableRow(studentDB, groupInputOptions))
}
