package skiphandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/classschema"
	"github.com/mangustc/obd/schema/groupschema"
	"github.com/mangustc/obd/schema/skipschema"
	"github.com/mangustc/obd/schema/studentschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/skipview"
)

func NewSkipHandler(
	ss handler.SessionService,
	us handler.UserService,
	js handler.JobService,
	sks handler.SkipService,
	cls handler.ClassService,
	sts handler.StudentService,
	grs handler.GroupService,
) *SkipHandler {
	return &SkipHandler{
		SessionService: ss,
		UserService:    us,
		JobService:     js,
		SkipService:    sks,
		ClassService:   cls,
		StudentService: sts,
		GroupService:   grs,
	}
}

type SkipHandler struct {
	SessionService handler.SessionService
	UserService    handler.UserService
	JobService     handler.JobService
	SkipService    handler.SkipService
	ClassService   handler.ClassService
	StudentService handler.StudentService
	GroupService   handler.GroupService
}

func (skh *SkipHandler) Skip(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	groupsDB, _ := skh.GroupService.GetGroups(&groupschema.GroupsGet{})
	classsDB, _ := skh.ClassService.GetClasss(&classschema.ClasssGet{})
	classInputOptions := classschema.GetClassInputOptionsFromClasssDB(classsDB, groupsDB)

	studentsDB, _ := skh.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	util.RenderComponent(r, &out, skipview.Skip(
		classInputOptions,
		studentInputOptions,
	))
}

func (skh *SkipHandler) SkipPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, skipview.SkipPage())
}

func (skh *SkipHandler) GetSkips(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &skipschema.SkipsGet{}

	err = skipschema.ValidateSkipsGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	skipsDB, err := skh.SkipService.GetSkips(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	groupsDB, _ := skh.GroupService.GetGroups(&groupschema.GroupsGet{})
	classsDB, _ := skh.ClassService.GetClasss(&classschema.ClasssGet{})
	classInputOptions := classschema.GetClassInputOptionsFromClasssDB(classsDB, groupsDB)

	studentsDB, _ := skh.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	util.RenderComponent(r, &out, skipview.SkipTableRows(skipsDB,
		classInputOptions,
		studentInputOptions,
	))
}

func (skh *SkipHandler) InsertSkip(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &skipschema.SkipInsert{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		skh.SessionService.GetSession,
		skh.UserService.GetUser,
		skh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessSkip {
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
	err = skipschema.ValidateSkipInsert(in)
	if err != nil {
		message = msg.SkipWrong
		logger.Error.Print(err.Error())
		return
	}

	skipDB, err := skh.SkipService.InsertSkip(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.SkipExists
			logger.Error.Print(err.Error())
			return
		}
	}

	groupsDB, _ := skh.GroupService.GetGroups(&groupschema.GroupsGet{})
	classsDB, _ := skh.ClassService.GetClasss(&classschema.ClasssGet{})
	classInputOptions := classschema.GetClassInputOptionsFromClasssDB(classsDB, groupsDB)

	studentsDB, _ := skh.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	util.RenderComponent(r, &out, skipview.SkipTableRow(skipDB,
		classInputOptions,
		studentInputOptions,
	))
}

func (skh *SkipHandler) EditSkip(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &skipschema.SkipGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		skh.SessionService.GetSession,
		skh.UserService.GetUser,
		skh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessSkip {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = skipschema.ValidateSkipGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	skipDB, err := skh.SkipService.GetSkip(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	groupsDB, _ := skh.GroupService.GetGroups(&groupschema.GroupsGet{})
	classsDB, _ := skh.ClassService.GetClasss(&classschema.ClasssGet{})
	classInputOptions := classschema.GetClassInputOptionsFromClasssDB(classsDB, groupsDB)

	studentsDB, _ := skh.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	util.RenderComponent(r, &out, skipview.SkipTableRowEdit(skipDB,
		classInputOptions,
		studentInputOptions,
	))
}

func (skh *SkipHandler) UpdateSkip(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &skipschema.SkipUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		skh.SessionService.GetSession,
		skh.UserService.GetUser,
		skh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessSkip {
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
	err = skipschema.ValidateSkipUpdate(in)
	if err != nil {
		message = msg.SkipWrong
		logger.Error.Print(err.Error())
		return
	}

	skipDB, err := skh.SkipService.UpdateSkip(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.SkipExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	groupsDB, _ := skh.GroupService.GetGroups(&groupschema.GroupsGet{})
	classsDB, _ := skh.ClassService.GetClasss(&classschema.ClasssGet{})
	classInputOptions := classschema.GetClassInputOptionsFromClasssDB(classsDB, groupsDB)

	studentsDB, _ := skh.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	util.RenderComponent(r, &out, skipview.SkipTableRow(skipDB,
		classInputOptions,
		studentInputOptions,
	))
}

func (skh *SkipHandler) DeleteSkip(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &skipschema.SkipDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		skh.SessionService.GetSession,
		skh.UserService.GetUser,
		skh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessSkip {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = skipschema.ValidateSkipDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	skipDB, err := skh.SkipService.DeleteSkip(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	groupsDB, _ := skh.GroupService.GetGroups(&groupschema.GroupsGet{})
	classsDB, _ := skh.ClassService.GetClasss(&classschema.ClasssGet{})
	classInputOptions := classschema.GetClassInputOptionsFromClasssDB(classsDB, groupsDB)

	studentsDB, _ := skh.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	util.RenderComponent(r, &out, skipview.SkipTableRow(skipDB,
		classInputOptions,
		studentInputOptions,
	))
}
