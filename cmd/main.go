package main

import (
	"fmt"
	"net/http"

	"github.com/mangustc/obd/database"
	"github.com/mangustc/obd/handler/jobhandler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/middleware"
	"github.com/mangustc/obd/service/jobservice"
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
	%[1]sCreatedAt DATETIME DEFAULT (datetime('now')),
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
	%[1]sAccessUniGroup INTEGER NOT NULL DEFAULT FALSE,
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
	%[1]sAccessClass INTEGER NOT NULL DEFAULT FALSE,
	%[1]sAccessSession
);`, jobTN)
	studentCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sLastname VARCHAR(35) NOT NULL,
	%[1]sFirstname VARCHAR(35) NOT NULL,
	%[1]sMiddlename VARCHAR(35) NOT NULL,
	%[1]sPhoneNumber VARCHAR(15) NOT NULL,
	%[2]sID INTEGER NOT NULL,
	%[1]sCreatedAt DATETIME DEFAULT (datetime('now')),
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE,
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE RESTRICT
);`, studentTN, groupTN)
	groupCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sNumber VARCHAR(10) NOT NULL,
	%[1]sYear INTEGER NOT NULL,
	%[1]sCourseName VARCHAR(50) NOT NULL,
	UNIQUE(%[1]sNumber, %[1]sYear)
);`, groupTN)
	finhelpCtgCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sDescription VARCHAR(250) NOT NULL,
	%[1]sPayment INTEGER NOT NULL
);`, finhelpCtgTN)
	finhelpStageCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
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
	%[1]sAddress VARCHAR(100) NOT NULL
);`, buildingTN)
	cabinetTypeCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sName VARCHAR(50) NOT NULL
);`, cabinetTypeTN)
	cabinetCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sNumber VARCHAR(10) NOT NULL,
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
	%[1]sCreatedAt DATETIME DEFAULT (datetime('now')),
	%[1]sIsHidden INTEGER NOT NULL DEFAULT FALSE
);`, profTN)
	courseTypeCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sName VARCHAR(50) NOT NULL
);`, courseTypeTN)
	courseCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sName VARCHAR(50) NOT NULL,
	%[2]sID INTEGER NOT NULL,
	FOREIGN KEY (%[2]sID) REFERENCES %[2]s (%[2]sID) ON DELETE RESTRICT
);`, courseTN, courseTypeTN)
	classTypeCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sName VARCHAR(50) NOT NULL
);`, classTypeTN)
	classCT = fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %[1]s (
	%[1]sID INTEGER PRIMARY KEY AUTOINCREMENT,
	%[1]sStart DATETIME NOT NULL,
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
	jh := jobhandler.NewJobHandler(js)

	router.HandleFunc("GET /job", jh.Job)
	router.HandleFunc("POST /api/job/getjobs", jh.GetJobs)
	router.HandleFunc("POST /api/job/insertjob", jh.InsertJob)
	router.HandleFunc("POST /api/job/updatejob", jh.UpdateJob)
	router.HandleFunc("POST /api/job/deletejob", jh.DeleteJob)
	router.HandleFunc("POST /api/job/editjob", jh.EditJob)

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
