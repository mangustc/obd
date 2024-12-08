package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB(dbName string) *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		panic("Failed to create DB: " + err.Error())
	}

	err = execQuery(db, "PRAGMA foreign_keys = ON")
	if err != nil {
		panic("Failed to create DB: " + err.Error())
	}

	return db
}

func execQuery(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

type UserService struct {
	DB     *sql.DB
	UserTN string
	JobTN  string
}

func NewUserService(db *sql.DB, userTN string, jobTN string) (us *UserService) {
	if db == nil || userTN == "" || jobTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &UserService{
		DB:     db,
		UserTN: userTN,
		JobTN:  jobTN,
	}
}

type UserDB struct {
	UserID         int    `json:"UserID"`
	UserLastname   string `json:"UserLastname"`
	UserFirstname  string `json:"UserFirstname"`
	UserMiddlename string `json:"UserMiddlename"`
	UserPassword   string `json:"UserPassword"`
	UserIsHidden   bool   `json:"UserIsHidden"`
	JobID          int    `json:"JobID"`
}

type JobDB struct {
	JobID                 int    `json:"JobID"`
	JobName               string `json:"JobName"`
	JobAccessUser         bool   `json:"JobAccessUser"`
	JobAccessJob          bool   `json:"JobAccessJob"`
	JobAccessStudent      bool   `json:"JobAccessStudent"`
	JobAccessGroup        bool   `json:"JobAccessGroup"`
	JobAccessFinhelpCtg   bool   `json:"JobAccessFinhelpCtg"`
	JobAccessFinhelpStage bool   `json:"JobAccessFinhelpStage"`
	JobAccessFinhelpProc  bool   `json:"JobAccessProc"`
	JobAccessBuilding     bool   `json:"JobAccessBuilding"`
	JobAccessCabinetType  bool   `json:"JobAccessCabinetType"`
	JobAccessCabinet      bool   `json:"JobAccessCabinet"`
	JobAccessClassType    bool   `json:"JobAccessClassType"`
	JobAccessProf         bool   `json:"JobAccessProf"`
	JobAccessCourseType   bool   `json:"JobAccessCourseType"`
	JobAccessCourse       bool   `json:"JobAccessCourse"`
	JobAccessPerf         bool   `json:"JobAccessPerf"`
	JobAccessSkip         bool   `json:"JobAccessSkip"`
	JobAccessClass        bool   `json:"JobAccessClass"`
}

type JobService struct {
	DB    *sql.DB
	JobTN string
}

func NewJobService(db *sql.DB, jobTN string) (us *JobService) {
	if db == nil || jobTN == "" {
		panic("Error creating service, one of the args is zero")
	}

	return &JobService{
		DB:    db,
		JobTN: jobTN,
	}
}

const (
	databaseFile = "database.db"
	jobTN        = "Job"
	userTN       = "User"
)

func main() {
	db := GetDB(databaseFile)
	js := NewJobService(db, jobTN)
	query := fmt.Sprintf(`
INSERT INTO %[1]s (
		%[1]sName,
		%[1]sAccessUser,
		%[1]sAccessJob,
		%[1]sAccessStudent,
		%[1]sAccessGroup,
		%[1]sAccessFinhelpCtg,
		%[1]sAccessFinhelpStage,
		%[1]sAccessFinhelpProc,
		%[1]sAccessBuilding,
		%[1]sAccessCabinetType,
		%[1]sAccessCabinet,
		%[1]sAccessClassType,
		%[1]sAccessProf,
		%[1]sAccessCourseType,
		%[1]sAccessCourse,
		%[1]sAccessPerf,
		%[1]sAccessSkip,
		%[1]sAccessClass
	)
	VALUES (
		"%[2]s",
		"%[3]t",
		"%[4]t",
		"%[5]t",
		"%[6]t",
		"%[7]t",
		"%[8]t",
		"%[9]t",
		"%[10]t",
		"%[11]t",
		"%[12]t",
		"%[13]t",
		"%[14]t",
		"%[15]t",
		"%[16]t",
		"%[17]t",
		"%[18]t",
		"%[19]t"
	)
RETURNING *`,
		js.JobTN,
		"Admin",
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
	)

	stmt, err := js.DB.Prepare(query)
	if err != nil {
		fmt.Println("Internal server error")
		return
	}

	jobDB := &JobDB{}
	err = stmt.QueryRow().Scan(
		&jobDB.JobID,
		&jobDB.JobName,
		&jobDB.JobAccessUser,
		&jobDB.JobAccessJob,
		&jobDB.JobAccessStudent,
		&jobDB.JobAccessGroup,
		&jobDB.JobAccessFinhelpCtg,
		&jobDB.JobAccessFinhelpStage,
		&jobDB.JobAccessFinhelpProc,
		&jobDB.JobAccessBuilding,
		&jobDB.JobAccessCabinetType,
		&jobDB.JobAccessCabinet,
		&jobDB.JobAccessClassType,
		&jobDB.JobAccessProf,
		&jobDB.JobAccessCourseType,
		&jobDB.JobAccessCourse,
		&jobDB.JobAccessPerf,
		&jobDB.JobAccessSkip,
		&jobDB.JobAccessClass,
	)
	if err != nil {
		fmt.Println("Internal server error")
		return
	}
	stmt.Close()

	us := NewUserService(db, userTN, jobTN)
	query = fmt.Sprintf(`
INSERT INTO %[1]s (
		%[1]sLastname,
		%[1]sFirstname,
		%[1]sMiddlename,
		%[1]sPassword,
		%[2]sID
	)
	VALUES (
		"%[3]s",
		"%[4]s",
		"%[5]s",
		"%[6]s",
		"%[7]d"
	)
RETURNING
	%[1]sID,
	%[1]sLastname,
	%[1]sFirstname,
	%[1]sMiddlename,
	%[1]sPassword,
	%[1]sIsHidden,
	%[2]sID`,
		us.UserTN,
		us.JobTN,
		"Admin",
		"Admin",
		"Admin",
		"Admin",
		jobDB.JobID,
	)

	stmt, err = us.DB.Prepare(query)
	if err != nil {
		fmt.Println("Internal server error")
		return
	}

	userDB := &UserDB{}
	err = stmt.QueryRow().Scan(
		&userDB.UserID,
		&userDB.UserLastname,
		&userDB.UserFirstname,
		&userDB.UserMiddlename,
		&userDB.UserPassword,
		&userDB.UserIsHidden,
		&userDB.JobID,
	)
	if err != nil {
		fmt.Println("Error")
		return
	}
	stmt.Close()
}
