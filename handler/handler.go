package handler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/buildingschema"
	"github.com/mangustc/obd/schema/cabinettypeschema"
	"github.com/mangustc/obd/schema/classtypeschema"
	"github.com/mangustc/obd/schema/coursetypeschema"
	"github.com/mangustc/obd/schema/finhelpctgschema"
	"github.com/mangustc/obd/schema/finhelpprocschema"
	"github.com/mangustc/obd/schema/finhelpstageschema"
	"github.com/mangustc/obd/schema/groupschema"
	"github.com/mangustc/obd/schema/jobschema"
	"github.com/mangustc/obd/schema/sessionschema"
	"github.com/mangustc/obd/schema/studentschema"
	"github.com/mangustc/obd/schema/userschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view"
)

type (
	JobService interface {
		InsertJob(data *jobschema.JobInsert) (jobDB *jobschema.JobDB, err error)
		UpdateJob(data *jobschema.JobUpdate) (jobDB *jobschema.JobDB, err error)
		DeleteJob(data *jobschema.JobDelete) (jobDB *jobschema.JobDB, err error)
		GetJob(data *jobschema.JobGet) (jobDB *jobschema.JobDB, err error)
		GetJobs(data *jobschema.JobsGet) (jobsDB []*jobschema.JobDB, err error)
	}
	UserService interface {
		InsertUser(data *userschema.UserInsert) (userDB *userschema.UserDB, err error)
		UpdateUser(data *userschema.UserUpdate) (userDB *userschema.UserDB, err error)
		DeleteUser(data *userschema.UserDelete) (userDB *userschema.UserDB, err error)
		GetUser(data *userschema.UserGet) (userDB *userschema.UserDB, err error)
		GetUsers(data *userschema.UsersGet) (usersDB []*userschema.UserDB, err error)
	}
	SessionService interface {
		InsertSession(data *sessionschema.SessionInsert) (sessionDB *sessionschema.SessionDB, err error)
		DeleteSession(data *sessionschema.SessionDelete) (sessionDB *sessionschema.SessionDB, err error)
		GetSession(data *sessionschema.SessionGet) (sessionDB *sessionschema.SessionDB, err error)
		GetSessions(data *sessionschema.SessionsGet) (sessionsDB []*sessionschema.SessionDB, err error)
	}
	GroupService interface {
		InsertGroup(data *groupschema.GroupInsert) (groupDB *groupschema.GroupDB, err error)
		UpdateGroup(data *groupschema.GroupUpdate) (groupDB *groupschema.GroupDB, err error)
		DeleteGroup(data *groupschema.GroupDelete) (groupDB *groupschema.GroupDB, err error)
		GetGroup(data *groupschema.GroupGet) (groupDB *groupschema.GroupDB, err error)
		GetGroups(data *groupschema.GroupsGet) (groupsDB []*groupschema.GroupDB, err error)
	}
	FinhelpCtgService interface {
		InsertFinhelpCtg(data *finhelpctgschema.FinhelpCtgInsert) (finhelpCtgDB *finhelpctgschema.FinhelpCtgDB, err error)
		UpdateFinhelpCtg(data *finhelpctgschema.FinhelpCtgUpdate) (finhelpCtgDB *finhelpctgschema.FinhelpCtgDB, err error)
		DeleteFinhelpCtg(data *finhelpctgschema.FinhelpCtgDelete) (finhelpCtgDB *finhelpctgschema.FinhelpCtgDB, err error)
		GetFinhelpCtg(data *finhelpctgschema.FinhelpCtgGet) (finhelpCtgDB *finhelpctgschema.FinhelpCtgDB, err error)
		GetFinhelpCtgs(data *finhelpctgschema.FinhelpCtgsGet) (finhelpCtgsDB []*finhelpctgschema.FinhelpCtgDB, err error)
	}
	FinhelpStageService interface {
		InsertFinhelpStage(data *finhelpstageschema.FinhelpStageInsert) (finhelpStageDB *finhelpstageschema.FinhelpStageDB, err error)
		UpdateFinhelpStage(data *finhelpstageschema.FinhelpStageUpdate) (finhelpStageDB *finhelpstageschema.FinhelpStageDB, err error)
		DeleteFinhelpStage(data *finhelpstageschema.FinhelpStageDelete) (finhelpStageDB *finhelpstageschema.FinhelpStageDB, err error)
		GetFinhelpStage(data *finhelpstageschema.FinhelpStageGet) (finhelpStageDB *finhelpstageschema.FinhelpStageDB, err error)
		GetFinhelpStages(data *finhelpstageschema.FinhelpStagesGet) (finhelpStagesDB []*finhelpstageschema.FinhelpStageDB, err error)
	}
	FinhelpProcService interface {
		InsertFinhelpProc(data *finhelpprocschema.FinhelpProcInsert) (finhelpProcDB *finhelpprocschema.FinhelpProcDB, err error)
		UpdateFinhelpProc(data *finhelpprocschema.FinhelpProcUpdate) (finhelpProcDB *finhelpprocschema.FinhelpProcDB, err error)
		DeleteFinhelpProc(data *finhelpprocschema.FinhelpProcDelete) (finhelpProcDB *finhelpprocschema.FinhelpProcDB, err error)
		GetFinhelpProc(data *finhelpprocschema.FinhelpProcGet) (finhelpProcDB *finhelpprocschema.FinhelpProcDB, err error)
		GetFinhelpProcs(data *finhelpprocschema.FinhelpProcsGet) (finhelpProcsDB []*finhelpprocschema.FinhelpProcDB, err error)
	}
	StudentService interface {
		InsertStudent(data *studentschema.StudentInsert) (studentDB *studentschema.StudentDB, err error)
		UpdateStudent(data *studentschema.StudentUpdate) (studentDB *studentschema.StudentDB, err error)
		DeleteStudent(data *studentschema.StudentDelete) (studentDB *studentschema.StudentDB, err error)
		GetStudent(data *studentschema.StudentGet) (studentDB *studentschema.StudentDB, err error)
		GetStudents(data *studentschema.StudentsGet) (studentsDB []*studentschema.StudentDB, err error)
	}
	BuildingService interface {
		InsertBuilding(data *buildingschema.BuildingInsert) (buildingDB *buildingschema.BuildingDB, err error)
		UpdateBuilding(data *buildingschema.BuildingUpdate) (buildingDB *buildingschema.BuildingDB, err error)
		DeleteBuilding(data *buildingschema.BuildingDelete) (buildingDB *buildingschema.BuildingDB, err error)
		GetBuilding(data *buildingschema.BuildingGet) (buildingDB *buildingschema.BuildingDB, err error)
		GetBuildings(data *buildingschema.BuildingsGet) (buildingsDB []*buildingschema.BuildingDB, err error)
	}
	CabinetTypeService interface {
		InsertCabinetType(data *cabinettypeschema.CabinetTypeInsert) (cabinetTypeDB *cabinettypeschema.CabinetTypeDB, err error)
		UpdateCabinetType(data *cabinettypeschema.CabinetTypeUpdate) (cabinetTypeDB *cabinettypeschema.CabinetTypeDB, err error)
		DeleteCabinetType(data *cabinettypeschema.CabinetTypeDelete) (cabinetTypeDB *cabinettypeschema.CabinetTypeDB, err error)
		GetCabinetType(data *cabinettypeschema.CabinetTypeGet) (cabinetTypeDB *cabinettypeschema.CabinetTypeDB, err error)
		GetCabinetTypes(data *cabinettypeschema.CabinetTypesGet) (cabinetTypesDB []*cabinettypeschema.CabinetTypeDB, err error)
	}
	ClassTypeService interface {
		InsertClassType(data *classtypeschema.ClassTypeInsert) (classTypeDB *classtypeschema.ClassTypeDB, err error)
		UpdateClassType(data *classtypeschema.ClassTypeUpdate) (classTypeDB *classtypeschema.ClassTypeDB, err error)
		DeleteClassType(data *classtypeschema.ClassTypeDelete) (classTypeDB *classtypeschema.ClassTypeDB, err error)
		GetClassType(data *classtypeschema.ClassTypeGet) (classTypeDB *classtypeschema.ClassTypeDB, err error)
		GetClassTypes(data *classtypeschema.ClassTypesGet) (classTypesDB []*classtypeschema.ClassTypeDB, err error)
	}
	CourseTypeService interface {
		InsertCourseType(data *coursetypeschema.CourseTypeInsert) (courseTypeDB *coursetypeschema.CourseTypeDB, err error)
		UpdateCourseType(data *coursetypeschema.CourseTypeUpdate) (courseTypeDB *coursetypeschema.CourseTypeDB, err error)
		DeleteCourseType(data *coursetypeschema.CourseTypeDelete) (courseTypeDB *coursetypeschema.CourseTypeDB, err error)
		GetCourseType(data *coursetypeschema.CourseTypeGet) (courseTypeDB *coursetypeschema.CourseTypeDB, err error)
		GetCourseTypes(data *coursetypeschema.CourseTypesGet) (courseTypesDB []*coursetypeschema.CourseTypeDB, err error)
	}
)

func NewDefaultHandler(ss SessionService, us UserService, js JobService) *DefaultHandler {
	return &DefaultHandler{
		SessionService: ss,
		UserService:    us,
		JobService:     js,
	}
}

type DefaultHandler struct {
	SessionService SessionService
	UserService    UserService
	JobService     JobService
}

func (dh *DefaultHandler) Default(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, view.Layout("OBD"))
}

func (dh *DefaultHandler) Navigation(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	jobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		dh.SessionService.GetSession,
		dh.UserService.GetUser,
		dh.JobService.GetJob,
	)
	if err != nil {
		if err == errs.ErrNotFound {
			jobDB = nil
		} else {
			err := errs.ErrUnauthorized
			message = msg.Unauthorized
			logger.Error.Print(err.Error())
			return
		}
	}

	util.RenderComponent(r, &out, view.NavigationByJobDB(jobDB))
}
