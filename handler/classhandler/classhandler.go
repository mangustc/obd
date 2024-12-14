package classhandler

import (
	"fmt"
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/buildingschema"
	"github.com/mangustc/obd/schema/cabinetschema"
	"github.com/mangustc/obd/schema/classschema"
	"github.com/mangustc/obd/schema/classtypeschema"
	"github.com/mangustc/obd/schema/courseschema"
	"github.com/mangustc/obd/schema/groupschema"
	"github.com/mangustc/obd/schema/profschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/classview"
)

func NewClassHandler(
	ss handler.SessionService,
	us handler.UserService,
	js handler.JobService,
	cls handler.ClassService,
	cots handler.ClassTypeService,
	prs handler.ProfService,
	cs handler.CabinetService,
	cos handler.CourseService,
	grs handler.GroupService,
	bs handler.BuildingService,
) *ClassHandler {
	return &ClassHandler{
		SessionService:   ss,
		UserService:      us,
		JobService:       js,
		ClassService:     cls,
		ClassTypeService: cots,
		ProfService:      prs,
		CabinetService:   cs,
		CourseService:    cos,
		GroupService:     grs,
		BuildingService:  bs,
	}
}

type ClassHandler struct {
	SessionService   handler.SessionService
	UserService      handler.UserService
	JobService       handler.JobService
	ClassService     handler.ClassService
	ClassTypeService handler.ClassTypeService
	ProfService      handler.ProfService
	CabinetService   handler.CabinetService
	CourseService    handler.CourseService
	GroupService     handler.GroupService
	BuildingService  handler.BuildingService
}

func (clh *ClassHandler) Class(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	classTypesDB, _ := clh.ClassTypeService.GetClassTypes(&classtypeschema.ClassTypesGet{})
	classTypeInputOptions := classtypeschema.GetClassTypeInputOptionsFromClassTypesDB(classTypesDB)

	profsDB, _ := clh.ProfService.GetProfs(&profschema.ProfsGet{})
	profInputOptions := profschema.GetProfInputOptionsFromProfsDB(profsDB)

	buildingsDB, _ := clh.BuildingService.GetBuildings(&buildingschema.BuildingsGet{})
	cabinetsDB, _ := clh.CabinetService.GetCabinets(&cabinetschema.CabinetsGet{})
	cabinetInputOptions := cabinetschema.GetCabinetInputOptionsFromCabinetsDB(cabinetsDB, buildingsDB)

	coursesDB, _ := clh.CourseService.GetCourses(&courseschema.CoursesGet{})
	courseInputOptions := courseschema.GetCourseInputOptionsFromCoursesDB(coursesDB)

	groupsDB, _ := clh.GroupService.GetGroups(&groupschema.GroupsGet{})
	groupInputOptions := groupschema.GetGroupInputOptionsFromGroupsDB(groupsDB)

	util.RenderComponent(r, &out, classview.Class(
		classTypeInputOptions,
		profInputOptions,
		cabinetInputOptions,
		courseInputOptions,
		groupInputOptions,
	))
}

func (clh *ClassHandler) ClassPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, classview.ClassPage())
}

func (clh *ClassHandler) GetClasss(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &classschema.ClasssGet{}

	err = classschema.ValidateClasssGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	classsDB, err := clh.ClassService.GetClasss(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	for _, classDB := range classsDB {
		classDB.ClassStart = fmt.Sprintf("%.10s", classDB.ClassStart)
	}

	classTypesDB, _ := clh.ClassTypeService.GetClassTypes(&classtypeschema.ClassTypesGet{})
	classTypeInputOptions := classtypeschema.GetClassTypeInputOptionsFromClassTypesDB(classTypesDB)

	profsDB, _ := clh.ProfService.GetProfs(&profschema.ProfsGet{})
	profInputOptions := profschema.GetProfInputOptionsFromProfsDB(profsDB)

	buildingsDB, _ := clh.BuildingService.GetBuildings(&buildingschema.BuildingsGet{})
	cabinetsDB, _ := clh.CabinetService.GetCabinets(&cabinetschema.CabinetsGet{})
	cabinetInputOptions := cabinetschema.GetCabinetInputOptionsFromCabinetsDB(cabinetsDB, buildingsDB)

	coursesDB, _ := clh.CourseService.GetCourses(&courseschema.CoursesGet{})
	courseInputOptions := courseschema.GetCourseInputOptionsFromCoursesDB(coursesDB)

	groupsDB, _ := clh.GroupService.GetGroups(&groupschema.GroupsGet{})
	groupInputOptions := groupschema.GetGroupInputOptionsFromGroupsDB(groupsDB)

	util.RenderComponent(r, &out, classview.ClassTableRows(classsDB,
		classTypeInputOptions,
		profInputOptions,
		cabinetInputOptions,
		courseInputOptions,
		groupInputOptions,
	))
}

func (clh *ClassHandler) InsertClass(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &classschema.ClassInsert{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		clh.SessionService.GetSession,
		clh.UserService.GetUser,
		clh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessClass {
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
	err = classschema.ValidateClassInsert(in)
	if err != nil {
		message = msg.ClassWrong
		logger.Error.Print(err.Error())
		return
	}

	classDB, err := clh.ClassService.InsertClass(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.ClassExists
			logger.Error.Print(err.Error())
			return
		}
	}
	classDB.ClassStart = fmt.Sprintf("%.10s", classDB.ClassStart)

	classTypesDB, _ := clh.ClassTypeService.GetClassTypes(&classtypeschema.ClassTypesGet{})
	classTypeInputOptions := classtypeschema.GetClassTypeInputOptionsFromClassTypesDB(classTypesDB)

	profsDB, _ := clh.ProfService.GetProfs(&profschema.ProfsGet{})
	profInputOptions := profschema.GetProfInputOptionsFromProfsDB(profsDB)

	buildingsDB, _ := clh.BuildingService.GetBuildings(&buildingschema.BuildingsGet{})
	cabinetsDB, _ := clh.CabinetService.GetCabinets(&cabinetschema.CabinetsGet{})
	cabinetInputOptions := cabinetschema.GetCabinetInputOptionsFromCabinetsDB(cabinetsDB, buildingsDB)

	coursesDB, _ := clh.CourseService.GetCourses(&courseschema.CoursesGet{})
	courseInputOptions := courseschema.GetCourseInputOptionsFromCoursesDB(coursesDB)

	groupsDB, _ := clh.GroupService.GetGroups(&groupschema.GroupsGet{})
	groupInputOptions := groupschema.GetGroupInputOptionsFromGroupsDB(groupsDB)

	util.RenderComponent(r, &out, classview.ClassTableRow(classDB,
		classTypeInputOptions,
		profInputOptions,
		cabinetInputOptions,
		courseInputOptions,
		groupInputOptions,
	))
}

func (clh *ClassHandler) EditClass(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &classschema.ClassGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		clh.SessionService.GetSession,
		clh.UserService.GetUser,
		clh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessClass {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = classschema.ValidateClassGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	classDB, err := clh.ClassService.GetClass(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	classDB.ClassStart = fmt.Sprintf("%.10s", classDB.ClassStart)

	classTypesDB, _ := clh.ClassTypeService.GetClassTypes(&classtypeschema.ClassTypesGet{})
	classTypeInputOptions := classtypeschema.GetClassTypeInputOptionsFromClassTypesDB(classTypesDB)

	profsDB, _ := clh.ProfService.GetProfs(&profschema.ProfsGet{})
	profInputOptions := profschema.GetProfInputOptionsFromProfsDB(profsDB)

	buildingsDB, _ := clh.BuildingService.GetBuildings(&buildingschema.BuildingsGet{})
	cabinetsDB, _ := clh.CabinetService.GetCabinets(&cabinetschema.CabinetsGet{})
	cabinetInputOptions := cabinetschema.GetCabinetInputOptionsFromCabinetsDB(cabinetsDB, buildingsDB)

	coursesDB, _ := clh.CourseService.GetCourses(&courseschema.CoursesGet{})
	courseInputOptions := courseschema.GetCourseInputOptionsFromCoursesDB(coursesDB)

	groupsDB, _ := clh.GroupService.GetGroups(&groupschema.GroupsGet{})
	groupInputOptions := groupschema.GetGroupInputOptionsFromGroupsDB(groupsDB)

	util.RenderComponent(r, &out, classview.ClassTableRowEdit(classDB,
		classTypeInputOptions,
		profInputOptions,
		cabinetInputOptions,
		courseInputOptions,
		groupInputOptions,
	))
}

func (clh *ClassHandler) UpdateClass(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &classschema.ClassUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		clh.SessionService.GetSession,
		clh.UserService.GetUser,
		clh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessClass {
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
	err = classschema.ValidateClassUpdate(in)
	if err != nil {
		message = msg.ClassWrong
		logger.Error.Print(err.Error())
		return
	}

	classDB, err := clh.ClassService.UpdateClass(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.ClassExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}
	classDB.ClassStart = fmt.Sprintf("%.10s", classDB.ClassStart)

	classTypesDB, _ := clh.ClassTypeService.GetClassTypes(&classtypeschema.ClassTypesGet{})
	classTypeInputOptions := classtypeschema.GetClassTypeInputOptionsFromClassTypesDB(classTypesDB)

	profsDB, _ := clh.ProfService.GetProfs(&profschema.ProfsGet{})
	profInputOptions := profschema.GetProfInputOptionsFromProfsDB(profsDB)

	buildingsDB, _ := clh.BuildingService.GetBuildings(&buildingschema.BuildingsGet{})
	cabinetsDB, _ := clh.CabinetService.GetCabinets(&cabinetschema.CabinetsGet{})
	cabinetInputOptions := cabinetschema.GetCabinetInputOptionsFromCabinetsDB(cabinetsDB, buildingsDB)

	coursesDB, _ := clh.CourseService.GetCourses(&courseschema.CoursesGet{})
	courseInputOptions := courseschema.GetCourseInputOptionsFromCoursesDB(coursesDB)

	groupsDB, _ := clh.GroupService.GetGroups(&groupschema.GroupsGet{})
	groupInputOptions := groupschema.GetGroupInputOptionsFromGroupsDB(groupsDB)

	util.RenderComponent(r, &out, classview.ClassTableRow(classDB,
		classTypeInputOptions,
		profInputOptions,
		cabinetInputOptions,
		courseInputOptions,
		groupInputOptions,
	))
}

func (clh *ClassHandler) DeleteClass(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &classschema.ClassDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		clh.SessionService.GetSession,
		clh.UserService.GetUser,
		clh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessClass {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = classschema.ValidateClassDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	classDB, err := clh.ClassService.DeleteClass(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	classDB.ClassStart = fmt.Sprintf("%.10s", classDB.ClassStart)

	classTypesDB, _ := clh.ClassTypeService.GetClassTypes(&classtypeschema.ClassTypesGet{})
	classTypeInputOptions := classtypeschema.GetClassTypeInputOptionsFromClassTypesDB(classTypesDB)

	profsDB, _ := clh.ProfService.GetProfs(&profschema.ProfsGet{})
	profInputOptions := profschema.GetProfInputOptionsFromProfsDB(profsDB)

	buildingsDB, _ := clh.BuildingService.GetBuildings(&buildingschema.BuildingsGet{})
	cabinetsDB, _ := clh.CabinetService.GetCabinets(&cabinetschema.CabinetsGet{})
	cabinetInputOptions := cabinetschema.GetCabinetInputOptionsFromCabinetsDB(cabinetsDB, buildingsDB)

	coursesDB, _ := clh.CourseService.GetCourses(&courseschema.CoursesGet{})
	courseInputOptions := courseschema.GetCourseInputOptionsFromCoursesDB(coursesDB)

	groupsDB, _ := clh.GroupService.GetGroups(&groupschema.GroupsGet{})
	groupInputOptions := groupschema.GetGroupInputOptionsFromGroupsDB(groupsDB)

	util.RenderComponent(r, &out, classview.ClassTableRow(classDB,
		classTypeInputOptions,
		profInputOptions,
		cabinetInputOptions,
		courseInputOptions,
		groupInputOptions,
	))
}
