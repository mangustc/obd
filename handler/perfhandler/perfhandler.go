package perfhandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/courseschema"
	"github.com/mangustc/obd/schema/groupschema"
	"github.com/mangustc/obd/schema/perfschema"
	"github.com/mangustc/obd/schema/studentschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/perfview"
)

func NewPerfHandler(
	ss handler.SessionService,
	us handler.UserService,
	js handler.JobService,
	ps handler.PerfService,
	cos handler.CourseService,
	sts handler.StudentService,
	grs handler.GroupService,
) *PerfHandler {
	return &PerfHandler{
		SessionService: ss,
		UserService:    us,
		JobService:     js,
		PerfService:    ps,
		CourseService:  cos,
		StudentService: sts,
		GroupService:   grs,
	}
}

type PerfHandler struct {
	SessionService handler.SessionService
	UserService    handler.UserService
	JobService     handler.JobService
	PerfService    handler.PerfService
	CourseService  handler.CourseService
	StudentService handler.StudentService
	GroupService   handler.GroupService
}

func (ph *PerfHandler) Perf(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	coursesDB, _ := ph.CourseService.GetCourses(&courseschema.CoursesGet{})
	courseInputOptions := courseschema.GetCourseInputOptionsFromCoursesDB(coursesDB)

	groupsDB, _ := ph.GroupService.GetGroups(&groupschema.GroupsGet{})
	studentsDB, _ := ph.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	util.RenderComponent(r, &out, perfview.Perf(
		courseInputOptions,
		studentInputOptions,
	))
}

func (ph *PerfHandler) PerfPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, perfview.PerfPage())
}

func (ph *PerfHandler) GetPerfs(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &perfschema.PerfsGet{}

	err = perfschema.ValidatePerfsGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	perfsDB, err := ph.PerfService.GetPerfs(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	coursesDB, _ := ph.CourseService.GetCourses(&courseschema.CoursesGet{})
	courseInputOptions := courseschema.GetCourseInputOptionsFromCoursesDB(coursesDB)

	groupsDB, _ := ph.GroupService.GetGroups(&groupschema.GroupsGet{})
	studentsDB, _ := ph.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	util.RenderComponent(r, &out, perfview.PerfTableRows(perfsDB,
		courseInputOptions,
		studentInputOptions,
	))
}

func (ph *PerfHandler) InsertPerf(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &perfschema.PerfInsert{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		ph.SessionService.GetSession,
		ph.UserService.GetUser,
		ph.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessPerf {
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
	err = perfschema.ValidatePerfInsert(in)
	if err != nil {
		message = msg.PerfWrong
		logger.Error.Print(err.Error())
		return
	}

	perfDB, err := ph.PerfService.InsertPerf(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.PerfExists
			logger.Error.Print(err.Error())
			return
		}
	}

	coursesDB, _ := ph.CourseService.GetCourses(&courseschema.CoursesGet{})
	courseInputOptions := courseschema.GetCourseInputOptionsFromCoursesDB(coursesDB)

	groupsDB, _ := ph.GroupService.GetGroups(&groupschema.GroupsGet{})
	studentsDB, _ := ph.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	util.RenderComponent(r, &out, perfview.PerfTableRow(perfDB,
		courseInputOptions,
		studentInputOptions,
	))
}

func (ph *PerfHandler) EditPerf(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &perfschema.PerfGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		ph.SessionService.GetSession,
		ph.UserService.GetUser,
		ph.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessPerf {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = perfschema.ValidatePerfGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	perfDB, err := ph.PerfService.GetPerf(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	coursesDB, _ := ph.CourseService.GetCourses(&courseschema.CoursesGet{})
	courseInputOptions := courseschema.GetCourseInputOptionsFromCoursesDB(coursesDB)

	groupsDB, _ := ph.GroupService.GetGroups(&groupschema.GroupsGet{})
	studentsDB, _ := ph.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	util.RenderComponent(r, &out, perfview.PerfTableRowEdit(perfDB,
		courseInputOptions,
		studentInputOptions,
	))
}

func (ph *PerfHandler) UpdatePerf(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &perfschema.PerfUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		ph.SessionService.GetSession,
		ph.UserService.GetUser,
		ph.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessPerf {
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
	err = perfschema.ValidatePerfUpdate(in)
	if err != nil {
		message = msg.PerfWrong
		logger.Error.Print(err.Error())
		return
	}

	perfDB, err := ph.PerfService.UpdatePerf(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.PerfExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	coursesDB, _ := ph.CourseService.GetCourses(&courseschema.CoursesGet{})
	courseInputOptions := courseschema.GetCourseInputOptionsFromCoursesDB(coursesDB)

	groupsDB, _ := ph.GroupService.GetGroups(&groupschema.GroupsGet{})
	studentsDB, _ := ph.StudentService.GetStudents(&studentschema.StudentsGet{})
	studentInputOptions := studentschema.GetStudentInputOptionsFromStudentsDB(studentsDB, groupsDB)

	util.RenderComponent(r, &out, perfview.PerfTableRow(perfDB,
		courseInputOptions,
		studentInputOptions,
	))
}

func (ph *PerfHandler) DeletePerf(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &perfschema.PerfDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		ph.SessionService.GetSession,
		ph.UserService.GetUser,
		ph.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessPerf {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = perfschema.ValidatePerfDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	_, err = ph.PerfService.DeletePerf(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
}
