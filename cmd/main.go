package main

import (
	"fmt"
	"net/http"

	"github.com/mangustc/obd/database"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/handler/authhandler"
	"github.com/mangustc/obd/handler/buildinghandler"
	"github.com/mangustc/obd/handler/cabinethandler"
	"github.com/mangustc/obd/handler/cabinettypehandler"
	"github.com/mangustc/obd/handler/classhandler"
	"github.com/mangustc/obd/handler/classtypehandler"
	"github.com/mangustc/obd/handler/coursehandler"
	"github.com/mangustc/obd/handler/coursetypehandler"
	"github.com/mangustc/obd/handler/finhelpctghandler"
	"github.com/mangustc/obd/handler/finhelpprochandler"
	"github.com/mangustc/obd/handler/finhelpstagehandler"
	"github.com/mangustc/obd/handler/grouphandler"
	"github.com/mangustc/obd/handler/jobhandler"
	"github.com/mangustc/obd/handler/perfhandler"
	"github.com/mangustc/obd/handler/profhandler"
	"github.com/mangustc/obd/handler/skiphandler"
	"github.com/mangustc/obd/handler/studenthandler"
	"github.com/mangustc/obd/handler/userhandler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/middleware"
	"github.com/mangustc/obd/service/buildingservice"
	"github.com/mangustc/obd/service/cabinetservice"
	"github.com/mangustc/obd/service/cabinettypeservice"
	"github.com/mangustc/obd/service/classservice"
	"github.com/mangustc/obd/service/classtypeservice"
	"github.com/mangustc/obd/service/courseservice"
	"github.com/mangustc/obd/service/coursetypeservice"
	"github.com/mangustc/obd/service/finhelpctgservice"
	"github.com/mangustc/obd/service/finhelpprocservice"
	"github.com/mangustc/obd/service/finhelpstageservice"
	"github.com/mangustc/obd/service/groupservice"
	"github.com/mangustc/obd/service/jobservice"
	"github.com/mangustc/obd/service/perfservice"
	"github.com/mangustc/obd/service/profservice"
	"github.com/mangustc/obd/service/sessionservice"
	"github.com/mangustc/obd/service/skipservice"
	"github.com/mangustc/obd/service/studentservice"
	"github.com/mangustc/obd/service/userservice"
)

var databaseFile = "database.db"

var (
	userTN         = "User"
	jobTN          = "Job"
	studentTN      = "Student"
	groupTN        = "UniGroup"   // University Group
	finhelpCtgTN   = "FinhelpCtg" // Financial Help Category
	finhelpStageTN = "FinhelpStage"
	finhelpProcTN  = "FinhelpProc" // Financial Help Process
	buildingTN     = "Building"
	cabinetTypeTN  = "CabinetType"
	cabinetTN      = "Cabinet"
	classTypeTN    = "ClassType"
	profTN         = "Prof" // Professor
	courseTypeTN   = "CourseType"
	courseTN       = "Course"
	perfTN         = "Perf" // Performance
	skipTN         = "Skip"
	classTN        = "Class"
	sessionTN      = "Session"
)

var (
	sessionCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sUUID VARCHAR(32) NOT NULL,
	%[2]sID INTEGER NOT NULL,
	%[1]sCreatedAt DATETIME DEFAULT (datetime('now')),
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE CASCADE
);`, sessionTN, userTN)
	userCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sLastname VARCHAR(35) NOT NULL,
	%[1]sFirstname VARCHAR(35) NOT NULL,
	%[1]sMiddlename VARCHAR(35) NOT NULL,
	%[1]sPassword VARCHAR(30) NOT NULL,
	%[2]sID INTEGER NOT NULL,
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE RESTRICT
);`, userTN, jobTN)
	jobCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sName VARCHAR(50) NOT NULL UNIQUE,
	%[1]sAccessUser INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessJob INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessStudent INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessGroup INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessFinhelpCtg INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessFinhelpStage INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessFinhelpProc INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessBuilding INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessCabinetType INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessCabinet INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessClassType INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessProf INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessCourseType INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessCourse INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessPerf INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessSkip INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessClass INTEGER NOT NULL DEFAULT FALSE
);`, jobTN)
	studentCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sLastname VARCHAR(35) NOT NULL,
	%[1]sFirstname VARCHAR(35) NOT NULL,
	%[1]sMiddlename VARCHAR(35) NOT NULL,
	%[1]sPhoneNumber VARCHAR(15) NOT NULL,
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	%[2]sID INTEGER NOT NULL,
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE RESTRICT
);`, studentTN, groupTN)
	groupCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sNumber VARCHAR(10) NOT NULL,
	%[1]sYear INTEGER NOT NULL,
	%[1]sCourseName VARCHAR(50) NOT NULL,
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	UNIQUE(%[1]sNumber, %[1]sYear)
);`, groupTN)
	finhelpCtgCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sDescription VARCHAR(250) NOT NULL,
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	%[1]sPayment INTEGER NOT NULL
);`, finhelpCtgTN)
	finhelpStageCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sName VARCHAR(50) NOT NULL,
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	%[1]sDescription VARCHAR(150) NOT NULL
);`, finhelpStageTN)
	finhelpProcCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[2]sID INTEGER NOT NULL,
	%[3]sID INTEGER NOT NULL,
	%[4]sID INTEGER NOT NULL,
	%[5]sID INTEGER NOT NULL,
	%[1]sCreatedAt DATETIME DEFAULT (datetime('now')),
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE RESTRICT,
	FOREIGN KEY (%[3]sID) REFERENCES %[3]s (%[3]sID) ON DELETE RESTRICT,
	FOREIGN KEY (%[4]sID) REFERENCES %[4]s (%[4]sID) ON DELETE RESTRICT,
	FOREIGN KEY (%[5]sID) REFERENCES %[5]s (%[5]sID) ON DELETE RESTRICT
);`, finhelpProcTN, studentTN, userTN, finhelpCtgTN, finhelpStageTN)
	buildingCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sName VARCHAR(50) NOT NULL,
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAddress VARCHAR(100) NOT NULL
);`, buildingTN)
	cabinetTypeCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	%[1]sName VARCHAR(50) NOT NULL
);`, cabinetTypeTN)
	cabinetCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sNumber VARCHAR(10) NOT NULL,
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	%[2]sID INTEGER NOT NULL,
	%[3]sID INTEGER NOT NULL,
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE RESTRICT,
	FOREIGN KEY (%[3]sID) REFERENCES %[3]s (%[3]sID) ON DELETE RESTRICT
);`, cabinetTN, buildingTN, cabinetTypeTN)
	profCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sLastname VARCHAR(35) NOT NULL,
	%[1]sFirstname VARCHAR(35) NOT NULL,
	%[1]sMiddlename VARCHAR(35) NOT NULL,
	%[1]sPhoneNumber VARCHAR(15) NOT NULL,
	%[1]sEmail VARCHAR(100) NOT NULL,
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE
);`, profTN)
	courseTypeCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	%[1]sName VARCHAR(50) NOT NULL
);`, courseTypeTN)
	courseCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sYear INTEGER NOT NULL,
	%[1]sName VARCHAR(50) NOT NULL,
	%[2]sID INTEGER NOT NULL,
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE RESTRICT
);`, courseTN, courseTypeTN)
	classTypeCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	%[1]sName VARCHAR(50) NOT NULL
);`, classTypeTN)
	classCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sStart DATE NOT NULL,
	%[1]sNumber INTEGER NOT NULL,
	%[2]sID INTEGER NOT NULL,
	%[3]sID INTEGER NOT NULL,
	%[4]sID INTEGER NOT NULL,
	%[5]sID INTEGER NOT NULL,
	%[6]sID INTEGER NOT NULL,
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE CASCADE,
	FOREIGN KEY (%[3]sID) REFERENCES %[3]s (%[3]sID) ON DELETE CASCADE,
	FOREIGN KEY (%[4]sID) REFERENCES %[4]s (%[4]sID) ON DELETE CASCADE,
	FOREIGN KEY (%[5]sID) REFERENCES %[5]s (%[5]sID) ON DELETE CASCADE,
	FOREIGN KEY (%[6]sID) REFERENCES %[6]s (%[6]sID) ON DELETE CASCADE,
	UNIQUE(%[1]sStart, %[1]sNumber, %[6]sID)
);`, classTN, classTypeTN, profTN, cabinetTN, courseTN, groupTN)
	skipCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[2]sID INTEGER NOT NULL,
	%[3]sID INTEGER NOT NULL,
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE CASCADE,
	FOREIGN KEY (%[3]sID) REFERENCES %[3]s (%[3]sID) ON DELETE CASCADE
);`, skipTN, classTN, studentTN)
	perfCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sGrade INTEGER NOT NULL DEFAULT 0,
	%[2]sID INTEGER NOT NULL,
	%[3]sID INTEGER NOT NULL,
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE CASCADE,
	FOREIGN KEY (%[3]sID) REFERENCES %[3]s (%[3]sID) ON DELETE CASCADE
);`, perfTN, courseTN, studentTN)
)

func main() {
	router := http.NewServeMux()
	db := database.GetDB(databaseFile)

	database.NewTable(db, jobCT)
	database.NewTable(db, cabinetTypeCT)
	database.NewTable(db, buildingCT)
	database.NewTable(db, classTypeCT)
	database.NewTable(db, profCT)
	database.NewTable(db, courseTypeCT)
	database.NewTable(db, groupCT)
	database.NewTable(db, finhelpStageCT)
	database.NewTable(db, finhelpCtgCT)
	database.NewTable(db, userCT)
	database.NewTable(db, sessionCT)
	database.NewTable(db, studentCT)
	database.NewTable(db, finhelpProcCT)
	database.NewTable(db, cabinetCT)
	database.NewTable(db, courseCT)
	database.NewTable(db, perfCT)
	database.NewTable(db, classCT)
	database.NewTable(db, skipCT)
	logger.Info.Println("Database created")

	js := jobservice.NewJobService(db,
		jobTN,
	)
	us := userservice.NewUserService(db,
		userTN,
		jobTN,
	)
	ss := sessionservice.NewSessionService(db,
		sessionTN,
		userTN,
	)
	grs := groupservice.NewGroupService(db,
		groupTN,
	)
	fctgs := finhelpctgservice.NewFinhelpCtgService(db,
		finhelpCtgTN,
	)
	fsts := finhelpstageservice.NewFinhelpStageService(db,
		finhelpStageTN,
	)
	fprs := finhelpprocservice.NewFinhelpProcService(db,
		finhelpProcTN,
		userTN,
		studentTN,
		finhelpCtgTN,
		finhelpStageTN,
	)
	sts := studentservice.NewStudentService(db,
		studentTN,
		groupTN,
	)
	bs := buildingservice.NewBuildingService(db,
		buildingTN,
	)
	cts := cabinettypeservice.NewCabinetTypeService(db,
		cabinetTypeTN,
	)
	clts := classtypeservice.NewClassTypeService(db,
		classTypeTN,
	)
	cots := coursetypeservice.NewCourseTypeService(db,
		courseTypeTN,
	)
	prs := profservice.NewProfService(db,
		profTN,
	)
	cs := cabinetservice.NewCabinetService(db,
		cabinetTN,
		buildingTN,
		cabinetTypeTN,
	)
	cos := courseservice.NewCourseService(db,
		courseTN,
		courseTypeTN,
	)
	cls := classservice.NewClassService(db,
		classTN,
		classTypeTN,
		profTN,
		cabinetTN,
		courseTN,
		groupTN,
	)
	ps := perfservice.NewPerfService(db,
		perfTN,
		courseTN,
		studentTN,
	)
	sks := skipservice.NewSkipService(db,
		skipTN,
		classTN,
		studentTN,
	)

	jh := jobhandler.NewJobHandler(
		ss,
		us,
		js,
	)
	router.HandleFunc("GET /job", jh.JobPage)
	router.HandleFunc("POST /api/job", jh.Job)
	router.HandleFunc("POST /api/job/getjobs", jh.GetJobs)
	router.HandleFunc("POST /api/job/insertjob", jh.InsertJob)
	router.HandleFunc("POST /api/job/updatejob", jh.UpdateJob)
	router.HandleFunc("POST /api/job/deletejob", jh.DeleteJob)
	router.HandleFunc("POST /api/job/editjob", jh.EditJob)

	uh := userhandler.NewUserHandler(ss,
		us,
		js,
	)
	router.HandleFunc("GET /user", uh.UserPage)
	router.HandleFunc("POST /api/user", uh.User)
	router.HandleFunc("POST /api/user/getusers", uh.GetUsers)
	router.HandleFunc("POST /api/user/insertuser", uh.InsertUser)
	router.HandleFunc("POST /api/user/updateuser", uh.UpdateUser)
	router.HandleFunc("POST /api/user/deleteuser", uh.DeleteUser)
	router.HandleFunc("POST /api/user/edituser", uh.EditUser)

	grh := grouphandler.NewGroupHandler(
		ss,
		us,
		js,
		grs,
	)
	router.HandleFunc("GET /group", grh.GroupPage)
	router.HandleFunc("POST /api/group", grh.Group)
	router.HandleFunc("POST /api/group/getgroups", grh.GetGroups)
	router.HandleFunc("POST /api/group/insertgroup", grh.InsertGroup)
	router.HandleFunc("POST /api/group/updategroup", grh.UpdateGroup)
	router.HandleFunc("POST /api/group/deletegroup", grh.DeleteGroup)
	router.HandleFunc("POST /api/group/editgroup", grh.EditGroup)

	fctgh := finhelpctghandler.NewFinhelpCtgHandler(
		ss,
		us,
		js,
		fctgs,
	)
	router.HandleFunc("GET /finhelpctg", fctgh.FinhelpCtgPage)
	router.HandleFunc("POST /api/finhelpctg", fctgh.FinhelpCtg)
	router.HandleFunc("POST /api/finhelpctg/getfinhelpctgs", fctgh.GetFinhelpCtgs)
	router.HandleFunc("POST /api/finhelpctg/insertfinhelpctg", fctgh.InsertFinhelpCtg)
	router.HandleFunc("POST /api/finhelpctg/updatefinhelpctg", fctgh.UpdateFinhelpCtg)
	router.HandleFunc("POST /api/finhelpctg/deletefinhelpctg", fctgh.DeleteFinhelpCtg)
	router.HandleFunc("POST /api/finhelpctg/editfinhelpctg", fctgh.EditFinhelpCtg)

	fsth := finhelpstagehandler.NewFinhelpStageHandler(
		ss,
		us,
		js,
		fsts,
	)
	router.HandleFunc("GET /finhelpstage", fsth.FinhelpStagePage)
	router.HandleFunc("POST /api/finhelpstage", fsth.FinhelpStage)
	router.HandleFunc("POST /api/finhelpstage/getfinhelpstages", fsth.GetFinhelpStages)
	router.HandleFunc("POST /api/finhelpstage/insertfinhelpstage", fsth.InsertFinhelpStage)
	router.HandleFunc("POST /api/finhelpstage/updatefinhelpstage", fsth.UpdateFinhelpStage)
	router.HandleFunc("POST /api/finhelpstage/deletefinhelpstage", fsth.DeleteFinhelpStage)
	router.HandleFunc("POST /api/finhelpstage/editfinhelpstage", fsth.EditFinhelpStage)

	fprh := finhelpprochandler.NewFinhelpProcHandler(
		ss,
		us,
		js,
		sts,
		fctgs,
		fsts,
		fprs,
		grs,
	)
	router.HandleFunc("GET /finhelpproc", fprh.FinhelpProcPage)
	router.HandleFunc("POST /api/finhelpproc", fprh.FinhelpProc)
	router.HandleFunc("POST /api/finhelpproc/getfinhelpprocs", fprh.GetFinhelpProcs)
	router.HandleFunc("POST /api/finhelpproc/insertfinhelpproc", fprh.InsertFinhelpProc)
	router.HandleFunc("POST /api/finhelpproc/updatefinhelpproc", fprh.UpdateFinhelpProc)
	router.HandleFunc("POST /api/finhelpproc/deletefinhelpproc", fprh.DeleteFinhelpProc)
	router.HandleFunc("POST /api/finhelpproc/editfinhelpproc", fprh.EditFinhelpProc)

	sth := studenthandler.NewStudentHandler(
		ss,
		us,
		js,
		grs,
		sts,
	)
	router.HandleFunc("GET /student", sth.StudentPage)
	router.HandleFunc("POST /api/student", sth.Student)
	router.HandleFunc("POST /api/student/getstudents", sth.GetStudents)
	router.HandleFunc("POST /api/student/insertstudent", sth.InsertStudent)
	router.HandleFunc("POST /api/student/updatestudent", sth.UpdateStudent)
	router.HandleFunc("POST /api/student/deletestudent", sth.DeleteStudent)
	router.HandleFunc("POST /api/student/editstudent", sth.EditStudent)

	bh := buildinghandler.NewBuildingHandler(
		ss,
		us,
		js,
		bs,
	)
	router.HandleFunc("GET /building", bh.BuildingPage)
	router.HandleFunc("POST /api/building", bh.Building)
	router.HandleFunc("POST /api/building/getbuildings", bh.GetBuildings)
	router.HandleFunc("POST /api/building/insertbuilding", bh.InsertBuilding)
	router.HandleFunc("POST /api/building/updatebuilding", bh.UpdateBuilding)
	router.HandleFunc("POST /api/building/deletebuilding", bh.DeleteBuilding)
	router.HandleFunc("POST /api/building/editbuilding", bh.EditBuilding)

	cth := cabinettypehandler.NewCabinetTypeHandler(
		ss,
		us,
		js,
		cts,
	)
	router.HandleFunc("GET /cabinettype", cth.CabinetTypePage)
	router.HandleFunc("POST /api/cabinettype", cth.CabinetType)
	router.HandleFunc("POST /api/cabinettype/getcabinettypes", cth.GetCabinetTypes)
	router.HandleFunc("POST /api/cabinettype/insertcabinettype", cth.InsertCabinetType)
	router.HandleFunc("POST /api/cabinettype/updatecabinettype", cth.UpdateCabinetType)
	router.HandleFunc("POST /api/cabinettype/deletecabinettype", cth.DeleteCabinetType)
	router.HandleFunc("POST /api/cabinettype/editcabinettype", cth.EditCabinetType)

	clth := classtypehandler.NewClassTypeHandler(
		ss,
		us,
		js,
		clts,
	)
	router.HandleFunc("GET /classtype", clth.ClassTypePage)
	router.HandleFunc("POST /api/classtype", clth.ClassType)
	router.HandleFunc("POST /api/classtype/getclasstypes", clth.GetClassTypes)
	router.HandleFunc("POST /api/classtype/insertclasstype", clth.InsertClassType)
	router.HandleFunc("POST /api/classtype/updateclasstype", clth.UpdateClassType)
	router.HandleFunc("POST /api/classtype/deleteclasstype", clth.DeleteClassType)
	router.HandleFunc("POST /api/classtype/editclasstype", clth.EditClassType)

	coth := coursetypehandler.NewCourseTypeHandler(
		ss,
		us,
		js,
		cots,
	)
	router.HandleFunc("GET /coursetype", coth.CourseTypePage)
	router.HandleFunc("POST /api/coursetype", coth.CourseType)
	router.HandleFunc("POST /api/coursetype/getcoursetypes", coth.GetCourseTypes)
	router.HandleFunc("POST /api/coursetype/insertcoursetype", coth.InsertCourseType)
	router.HandleFunc("POST /api/coursetype/updatecoursetype", coth.UpdateCourseType)
	router.HandleFunc("POST /api/coursetype/deletecoursetype", coth.DeleteCourseType)
	router.HandleFunc("POST /api/coursetype/editcoursetype", coth.EditCourseType)

	prh := profhandler.NewProfHandler(
		ss,
		us,
		js,
		prs,
	)
	router.HandleFunc("GET /prof", prh.ProfPage)
	router.HandleFunc("POST /api/prof", prh.Prof)
	router.HandleFunc("POST /api/prof/getprofs", prh.GetProfs)
	router.HandleFunc("POST /api/prof/insertprof", prh.InsertProf)
	router.HandleFunc("POST /api/prof/updateprof", prh.UpdateProf)
	router.HandleFunc("POST /api/prof/deleteprof", prh.DeleteProf)
	router.HandleFunc("POST /api/prof/editprof", prh.EditProf)

	ch := cabinethandler.NewCabinetHandler(
		ss,
		us,
		js,
		cs,
		bs,
		cts,
	)
	router.HandleFunc("GET /cabinet", ch.CabinetPage)
	router.HandleFunc("POST /api/cabinet", ch.Cabinet)
	router.HandleFunc("POST /api/cabinet/getcabinets", ch.GetCabinets)
	router.HandleFunc("POST /api/cabinet/insertcabinet", ch.InsertCabinet)
	router.HandleFunc("POST /api/cabinet/updatecabinet", ch.UpdateCabinet)
	router.HandleFunc("POST /api/cabinet/deletecabinet", ch.DeleteCabinet)
	router.HandleFunc("POST /api/cabinet/editcabinet", ch.EditCabinet)

	coh := coursehandler.NewCourseHandler(
		ss,
		us,
		js,
		cos,
		cots,
	)
	router.HandleFunc("GET /course", coh.CoursePage)
	router.HandleFunc("POST /api/course", coh.Course)
	router.HandleFunc("POST /api/course/getcourses", coh.GetCourses)
	router.HandleFunc("POST /api/course/insertcourse", coh.InsertCourse)
	router.HandleFunc("POST /api/course/updatecourse", coh.UpdateCourse)
	router.HandleFunc("POST /api/course/deletecourse", coh.DeleteCourse)
	router.HandleFunc("POST /api/course/editcourse", coh.EditCourse)

	clh := classhandler.NewClassHandler(
		ss,
		us,
		js,
		cls,
		clts,
		prs,
		cs,
		cos,
		grs,
		bs,
	)
	router.HandleFunc("GET /class", clh.ClassPage)
	router.HandleFunc("POST /api/class", clh.Class)
	router.HandleFunc("POST /api/class/getclasss", clh.GetClasss)
	router.HandleFunc("POST /api/class/insertclass", clh.InsertClass)
	router.HandleFunc("POST /api/class/updateclass", clh.UpdateClass)
	router.HandleFunc("POST /api/class/deleteclass", clh.DeleteClass)
	router.HandleFunc("POST /api/class/editclass", clh.EditClass)

	ph := perfhandler.NewPerfHandler(
		ss,
		us,
		js,
		ps,
		cos,
		sts,
		grs,
	)
	router.HandleFunc("GET /perf", ph.PerfPage)
	router.HandleFunc("POST /api/perf", ph.Perf)
	router.HandleFunc("POST /api/perf/getperfs", ph.GetPerfs)
	router.HandleFunc("POST /api/perf/insertperf", ph.InsertPerf)
	router.HandleFunc("POST /api/perf/updateperf", ph.UpdatePerf)
	router.HandleFunc("POST /api/perf/deleteperf", ph.DeletePerf)
	router.HandleFunc("POST /api/perf/editperf", ph.EditPerf)

	skh := skiphandler.NewSkipHandler(
		ss,
		us,
		js,
		sks,
		cls,
		sts,
		grs,
	)
	router.HandleFunc("GET /skip", skh.SkipPage)
	router.HandleFunc("POST /api/skip", skh.Skip)
	router.HandleFunc("POST /api/skip/getskips", skh.GetSkips)
	router.HandleFunc("POST /api/skip/insertskip", skh.InsertSkip)
	router.HandleFunc("POST /api/skip/updateskip", skh.UpdateSkip)
	router.HandleFunc("POST /api/skip/deleteskip", skh.DeleteSkip)
	router.HandleFunc("POST /api/skip/editskip", skh.EditSkip)

	auh := authhandler.NewAuthHandler(
		ss,
		us,
	)
	router.HandleFunc("GET /auth", auh.AuthPage)
	router.HandleFunc("POST /api/auth", auh.Auth)
	router.HandleFunc("POST /api/auth/userinput", auh.UserInput)
	router.HandleFunc("POST /api/auth/login", auh.AuthLogin)
	router.HandleFunc("POST /api/auth/logout", auh.AuthLogout)

	dh := handler.NewDefaultHandler(
		ss,
		us,
		js,
	)
	router.HandleFunc("GET /", dh.Default)
	router.HandleFunc("POST /api/navigation", dh.Navigation)
	cssDir := http.FileServer(http.Dir("./css"))
	jsDir := http.FileServer(http.Dir("./js"))
	router.Handle("GET /css/", http.StripPrefix("/css", cssDir))
	router.Handle("GET /js/", http.StripPrefix("/js", jsDir))
	port := ":1323"
	middlewareStack := middleware.CreateStack(
		middleware.Logging,
		middleware.StripSlash,
	)
	server := http.Server{
		Addr:    port,
		Handler: middlewareStack(router),
	}

	logger.Info.Println("Server is listening on port " + port)
	server.ListenAndServe()
}
