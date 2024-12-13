package finhelpprochandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/finhelpctgschema"
	"github.com/mangustc/obd/schema/finhelpprocschema"
	"github.com/mangustc/obd/schema/finhelpstageschema"
	"github.com/mangustc/obd/schema/groupschema"
	"github.com/mangustc/obd/schema/studentschema"
	"github.com/mangustc/obd/schema/userschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/finhelpprocview"
)

func NewFinhelpProcHandler(
	ss handler.SessionService,
	us handler.UserService,
	js handler.JobService,
	sts handler.StudentService,
	fctgs handler.FinhelpCtgService,
	fsts handler.FinhelpStageService,
	fprs handler.FinhelpProcService,
	grs handler.GroupService,
) *FinhelpProcHandler {
	return &FinhelpProcHandler{
		SessionService:      ss,
		UserService:         us,
		JobService:          js,
		StudentService:      sts,
		FinhelpCtgService:   fctgs,
		FinhelpStageService: fsts,
		FinhelpProcService:  fprs,
		GroupService:        grs,
	}
}

type FinhelpProcHandler struct {
	SessionService      handler.SessionService
	UserService         handler.UserService
	JobService          handler.JobService
	StudentService      handler.StudentService
	FinhelpCtgService   handler.FinhelpCtgService
	FinhelpStageService handler.FinhelpStageService
	FinhelpProcService  handler.FinhelpProcService
	GroupService        handler.GroupService
}

func (fprh *FinhelpProcHandler) FinhelpProc(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	groupsDB, _ := fprh.GroupService.GetGroups(&groupschema.GroupsGet{})
	studentsDB, _ := fprh.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	finhelpCtgsDB, _ := fprh.FinhelpCtgService.GetFinhelpCtgs(&finhelpctgschema.FinhelpCtgsGet{})
	finhelpCtgInputOptions := finhelpctgschema.GetFinhelpCtgInputOptionsFromFinhelpCtgsDB(finhelpCtgsDB)

	finhelpStagesDB, _ := fprh.FinhelpStageService.GetFinhelpStages(&finhelpstageschema.FinhelpStagesGet{})
	finhelpStageInputOptions := finhelpstageschema.GetFinhelpStageInputOptionsFromFinhelpStagesDB(finhelpStagesDB)

	util.RenderComponent(r, &out, finhelpprocview.FinhelpProc(
		studentInputOptions,
		finhelpCtgInputOptions,
		finhelpStageInputOptions,
	))
}

func (fprh *FinhelpProcHandler) FinhelpProcPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, finhelpprocview.FinhelpProcPage())
}

func (fprh *FinhelpProcHandler) GetFinhelpProcs(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpprocschema.FinhelpProcsGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		fprh.SessionService.GetSession,
		fprh.UserService.GetUser,
		fprh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpProc {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = finhelpprocschema.ValidateFinhelpProcsGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	finhelpProcsDB, err := fprh.FinhelpProcService.GetFinhelpProcs(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	usersDB, _ := fprh.UserService.GetUsers(&userschema.UsersGet{})
	userInputOptions := userschema.GetUserInputOptionsFromUsersDB(usersDB)

	groupsDB, _ := fprh.GroupService.GetGroups(&groupschema.GroupsGet{})
	studentsDB, _ := fprh.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	finhelpCtgsDB, _ := fprh.FinhelpCtgService.GetFinhelpCtgs(&finhelpctgschema.FinhelpCtgsGet{})
	finhelpCtgInputOptions := finhelpctgschema.GetFinhelpCtgInputOptionsFromFinhelpCtgsDB(finhelpCtgsDB)

	finhelpStagesDB, _ := fprh.FinhelpStageService.GetFinhelpStages(&finhelpstageschema.FinhelpStagesGet{})
	finhelpStageInputOptions := finhelpstageschema.GetFinhelpStageInputOptionsFromFinhelpStagesDB(finhelpStagesDB)

	util.RenderComponent(r, &out, finhelpprocview.FinhelpProcTableRows(finhelpProcsDB,
		userInputOptions,
		studentInputOptions,
		finhelpCtgInputOptions,
		finhelpStageInputOptions,
	))
}

func (fprh *FinhelpProcHandler) InsertFinhelpProc(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpprocschema.FinhelpProcInsert{}

	sessionJobDB, sessionUserDB, err := util.GetJobBySessionCookie(
		w, r,
		fprh.SessionService.GetSession,
		fprh.UserService.GetUser,
		fprh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpProc {
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
	in.UserID = sessionUserDB.UserID
	err = finhelpprocschema.ValidateFinhelpProcInsert(in)
	if err != nil {
		message = msg.FinhelpProcWrong
		logger.Error.Print(err.Error())
		return
	}

	finhelpProcDB, err := fprh.FinhelpProcService.InsertFinhelpProc(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.FinhelpProcExists
			logger.Error.Print(err.Error())
			return
		}
	}

	usersDB, _ := fprh.UserService.GetUsers(&userschema.UsersGet{})
	userInputOptions := userschema.GetUserInputOptionsFromUsersDB(usersDB)

	groupsDB, _ := fprh.GroupService.GetGroups(&groupschema.GroupsGet{})
	studentsDB, _ := fprh.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	finhelpCtgsDB, _ := fprh.FinhelpCtgService.GetFinhelpCtgs(&finhelpctgschema.FinhelpCtgsGet{})
	finhelpCtgInputOptions := finhelpctgschema.GetFinhelpCtgInputOptionsFromFinhelpCtgsDB(finhelpCtgsDB)

	finhelpStagesDB, _ := fprh.FinhelpStageService.GetFinhelpStages(&finhelpstageschema.FinhelpStagesGet{})
	finhelpStageInputOptions := finhelpstageschema.GetFinhelpStageInputOptionsFromFinhelpStagesDB(finhelpStagesDB)

	util.RenderComponent(r, &out, finhelpprocview.FinhelpProcTableRow(finhelpProcDB,
		userInputOptions,
		studentInputOptions,
		finhelpCtgInputOptions,
		finhelpStageInputOptions,
	))
}

func (fprh *FinhelpProcHandler) EditFinhelpProc(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpprocschema.FinhelpProcGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		fprh.SessionService.GetSession,
		fprh.UserService.GetUser,
		fprh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpProc {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = finhelpprocschema.ValidateFinhelpProcGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	finhelpProcDB, err := fprh.FinhelpProcService.GetFinhelpProc(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	usersDB, _ := fprh.UserService.GetUsers(&userschema.UsersGet{})
	userInputOptions := userschema.GetUserInputOptionsFromUsersDB(usersDB)

	groupsDB, _ := fprh.GroupService.GetGroups(&groupschema.GroupsGet{})
	studentsDB, _ := fprh.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	finhelpCtgsDB, _ := fprh.FinhelpCtgService.GetFinhelpCtgs(&finhelpctgschema.FinhelpCtgsGet{})
	finhelpCtgInputOptions := finhelpctgschema.GetFinhelpCtgInputOptionsFromFinhelpCtgsDB(finhelpCtgsDB)

	finhelpStagesDB, _ := fprh.FinhelpStageService.GetFinhelpStages(&finhelpstageschema.FinhelpStagesGet{})
	finhelpStageInputOptions := finhelpstageschema.GetFinhelpStageInputOptionsFromFinhelpStagesDB(finhelpStagesDB)

	util.RenderComponent(r, &out, finhelpprocview.FinhelpProcTableRowEdit(finhelpProcDB,
		userInputOptions,
		studentInputOptions,
		finhelpCtgInputOptions,
		finhelpStageInputOptions,
	))
}

func (fprh *FinhelpProcHandler) UpdateFinhelpProc(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpprocschema.FinhelpProcUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		fprh.SessionService.GetSession,
		fprh.UserService.GetUser,
		fprh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpProc {
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
	err = finhelpprocschema.ValidateFinhelpProcUpdate(in)
	if err != nil {
		message = msg.FinhelpProcWrong
		logger.Error.Print(err.Error())
		return
	}

	finhelpProcDB, err := fprh.FinhelpProcService.UpdateFinhelpProc(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.FinhelpProcExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	usersDB, _ := fprh.UserService.GetUsers(&userschema.UsersGet{})
	userInputOptions := userschema.GetUserInputOptionsFromUsersDB(usersDB)

	groupsDB, _ := fprh.GroupService.GetGroups(&groupschema.GroupsGet{})
	studentsDB, _ := fprh.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	finhelpCtgsDB, _ := fprh.FinhelpCtgService.GetFinhelpCtgs(&finhelpctgschema.FinhelpCtgsGet{})
	finhelpCtgInputOptions := finhelpctgschema.GetFinhelpCtgInputOptionsFromFinhelpCtgsDB(finhelpCtgsDB)

	finhelpStagesDB, _ := fprh.FinhelpStageService.GetFinhelpStages(&finhelpstageschema.FinhelpStagesGet{})
	finhelpStageInputOptions := finhelpstageschema.GetFinhelpStageInputOptionsFromFinhelpStagesDB(finhelpStagesDB)

	util.RenderComponent(r, &out, finhelpprocview.FinhelpProcTableRow(finhelpProcDB,
		userInputOptions,
		studentInputOptions,
		finhelpCtgInputOptions,
		finhelpStageInputOptions,
	))
}

func (fprh *FinhelpProcHandler) DeleteFinhelpProc(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &finhelpprocschema.FinhelpProcDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		fprh.SessionService.GetSession,
		fprh.UserService.GetUser,
		fprh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessFinhelpProc {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = finhelpprocschema.ValidateFinhelpProcDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	finhelpProcDB, err := fprh.FinhelpProcService.DeleteFinhelpProc(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	usersDB, _ := fprh.UserService.GetUsers(&userschema.UsersGet{})
	userInputOptions := userschema.GetUserInputOptionsFromUsersDB(usersDB)

	groupsDB, _ := fprh.GroupService.GetGroups(&groupschema.GroupsGet{})
	studentsDB, _ := fprh.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	finhelpCtgsDB, _ := fprh.FinhelpCtgService.GetFinhelpCtgs(&finhelpctgschema.FinhelpCtgsGet{})
	finhelpCtgInputOptions := finhelpctgschema.GetFinhelpCtgInputOptionsFromFinhelpCtgsDB(finhelpCtgsDB)

	finhelpStagesDB, _ := fprh.FinhelpStageService.GetFinhelpStages(&finhelpstageschema.FinhelpStagesGet{})
	finhelpStageInputOptions := finhelpstageschema.GetFinhelpStageInputOptionsFromFinhelpStagesDB(finhelpStagesDB)

	util.RenderComponent(r, &out, finhelpprocview.FinhelpProcTableRow(finhelpProcDB,
		userInputOptions,
		studentInputOptions,
		finhelpCtgInputOptions,
		finhelpStageInputOptions,
	))
}
