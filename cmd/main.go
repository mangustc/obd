package main

import (
	"fmt"
	"net/http"

	"github.com/mangustc/obd/database"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/handler/authhandler"
	"github.com/mangustc/obd/handler/finhelpctghandler"
	"github.com/mangustc/obd/handler/finhelpstagehandler"
	"github.com/mangustc/obd/handler/grouphandler"
	"github.com/mangustc/obd/handler/jobhandler"
	"github.com/mangustc/obd/handler/studenthandler"
	"github.com/mangustc/obd/handler/userhandler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/middleware"
	"github.com/mangustc/obd/service/finhelpctgservice"
	"github.com/mangustc/obd/service/finhelpstageservice"
	"github.com/mangustc/obd/service/groupservice"
	"github.com/mangustc/obd/service/jobservice"
	"github.com/mangustc/obd/service/sessionservice"
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
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	%[2]sID INTEGER NOT NULL,
	%[3]sID INTEGER NOT NULL,
	%[4]sID INTEGER NOT NULL,
	%[5]sID INTEGER NOT NULL,
	%[6]sID INTEGER NOT NULL,
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE CASCADE,
	FOREIGN KEY (%[3]sID) REFERENCES %[3]s (%[3]sID) ON DELETE CASCADE,
	FOREIGN KEY (%[4]sID) REFERENCES %[4]s (%[4]sID) ON DELETE CASCADE,
	FOREIGN KEY (%[5]sID) REFERENCES %[5]s (%[5]sID) ON DELETE CASCADE,
	FOREIGN KEY (%[6]sID) REFERENCES %[6]s (%[6]sID) ON DELETE CASCADE
);`, classTN, classTypeTN, profTN, cabinetTN, courseTN, groupTN)
	skipCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[2]sID INTEGER NOT NULL,
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE CASCADE
);`, skipTN, classTN)
	perfCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sGrade INTEGER,
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

	js := jobservice.NewJobService(db, jobTN)
	us := userservice.NewUserService(db, userTN, jobTN)
	ss := sessionservice.NewSessionService(db, sessionTN, userTN)
	grs := groupservice.NewGroupService(db, groupTN)
	fctgs := finhelpctgservice.NewFinhelpCtgService(db, finhelpCtgTN)
	fsts := finhelpstageservice.NewFinhelpStageService(db, finhelpStageTN)
	sts := studentservice.NewStudentService(db, studentTN, groupTN)

	jh := jobhandler.NewJobHandler(ss, us, js)
	router.HandleFunc("GET /job", jh.JobPage)
	router.HandleFunc("POST /api/job", jh.Job)
	router.HandleFunc("POST /api/job/getjobs", jh.GetJobs)
	router.HandleFunc("POST /api/job/insertjob", jh.InsertJob)
	router.HandleFunc("POST /api/job/updatejob", jh.UpdateJob)
	router.HandleFunc("POST /api/job/deletejob", jh.DeleteJob)
	router.HandleFunc("POST /api/job/editjob", jh.EditJob)

	uh := userhandler.NewUserHandler(ss, us, js)
	router.HandleFunc("GET /user", uh.UserPage)
	router.HandleFunc("POST /api/user", uh.User)
	router.HandleFunc("POST /api/user/getusers", uh.GetUsers)
	router.HandleFunc("POST /api/user/insertuser", uh.InsertUser)
	router.HandleFunc("POST /api/user/updateuser", uh.UpdateUser)
	router.HandleFunc("POST /api/user/deleteuser", uh.DeleteUser)
	router.HandleFunc("POST /api/user/edituser", uh.EditUser)

	grh := grouphandler.NewGroupHandler(ss, us, js, grs)
	router.HandleFunc("GET /group", grh.GroupPage)
	router.HandleFunc("POST /api/group", grh.Group)
	router.HandleFunc("POST /api/group/getgroups", grh.GetGroups)
	router.HandleFunc("POST /api/group/insertgroup", grh.InsertGroup)
	router.HandleFunc("POST /api/group/updategroup", grh.UpdateGroup)
	router.HandleFunc("POST /api/group/deletegroup", grh.DeleteGroup)
	router.HandleFunc("POST /api/group/editgroup", grh.EditGroup)

	fctgh := finhelpctghandler.NewFinhelpCtgHandler(ss, us, js, fctgs)
	router.HandleFunc("GET /finhelpctg", fctgh.FinhelpCtgPage)
	router.HandleFunc("POST /api/finhelpctg", fctgh.FinhelpCtg)
	router.HandleFunc("POST /api/finhelpctg/getfinhelpctgs", fctgh.GetFinhelpCtgs)
	router.HandleFunc("POST /api/finhelpctg/insertfinhelpctg", fctgh.InsertFinhelpCtg)
	router.HandleFunc("POST /api/finhelpctg/updatefinhelpctg", fctgh.UpdateFinhelpCtg)
	router.HandleFunc("POST /api/finhelpctg/deletefinhelpctg", fctgh.DeleteFinhelpCtg)
	router.HandleFunc("POST /api/finhelpctg/editfinhelpctg", fctgh.EditFinhelpCtg)

	fsth := finhelpstagehandler.NewFinhelpStageHandler(ss, us, js, fsts)
	router.HandleFunc("GET /finhelpstage", fsth.FinhelpStagePage)
	router.HandleFunc("POST /api/finhelpstage", fsth.FinhelpStage)
	router.HandleFunc("POST /api/finhelpstage/getfinhelpstages", fsth.GetFinhelpStages)
	router.HandleFunc("POST /api/finhelpstage/insertfinhelpstage", fsth.InsertFinhelpStage)
	router.HandleFunc("POST /api/finhelpstage/updatefinhelpstage", fsth.UpdateFinhelpStage)
	router.HandleFunc("POST /api/finhelpstage/deletefinhelpstage", fsth.DeleteFinhelpStage)
	router.HandleFunc("POST /api/finhelpstage/editfinhelpstage", fsth.EditFinhelpStage)

	sth := studenthandler.NewStudentHandler(ss, us, js, grs, sts)
	router.HandleFunc("GET /student", sth.StudentPage)
	router.HandleFunc("POST /api/student", sth.Student)
	router.HandleFunc("POST /api/student/getstudents", sth.GetStudents)
	router.HandleFunc("POST /api/student/insertstudent", sth.InsertStudent)
	router.HandleFunc("POST /api/student/updatestudent", sth.UpdateStudent)
	router.HandleFunc("POST /api/student/deletestudent", sth.DeleteStudent)
	router.HandleFunc("POST /api/student/editstudent", sth.EditStudent)

	auh := authhandler.NewAuthHandler(ss, us)
	router.HandleFunc("GET /auth", auh.AuthPage)
	router.HandleFunc("POST /api/auth", auh.Auth)
	router.HandleFunc("POST /api/auth/userinput", auh.UserInput)
	router.HandleFunc("POST /api/auth/login", auh.AuthLogin)
	router.HandleFunc("POST /api/auth/logout", auh.AuthLogout)

	dh := handler.NewDefaultHandler(ss, us, js)
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
