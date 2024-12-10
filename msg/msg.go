package msg

import (
	"net/http"

	"github.com/mangustc/obd/schema"
)

type Msg struct {
	MsgCode             int
	MsgNotificationType schema.NotificationType
	MsgStr              string
}

func NewMsg(msgCode int, msgNotificationType schema.NotificationType, msgStr string) *Msg {
	if msgCode == 0 || msgNotificationType == 0 {
		panic("One or more arguments are zero")
	}
	return &Msg{
		MsgCode:             msgCode,
		MsgNotificationType: msgNotificationType,
		MsgStr:              msgStr,
	}
}

var (
	Nothing = NewMsg(
		http.StatusOK,
		schema.NoNotification,
		"",
	)
	OK = NewMsg(
		http.StatusOK,
		schema.SuccessNotification,
		"Successful",
	)
	InternalServerError = NewMsg(
		http.StatusInternalServerError,
		schema.ErrorNotification,
		"Internal server error, try refreshing the page",
	)
	AuthWrongPassword = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Wrong password, please try again",
	)
	AuthSuccessLogin = NewMsg(
		http.StatusOK,
		schema.SuccessNotification,
		"Successfully logged in",
	)
	AuthSuccessLogout = NewMsg(
		http.StatusOK,
		schema.SuccessNotification,
		"Successfully logged out",
	)
	Unauthorized = NewMsg(
		http.StatusUnauthorized,
		schema.ErrorNotification,
		"Unauthorized access",
	)
	JobNameEmpty = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Job name can't be empty",
	)
	JobExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"This job already exists",
	)
	UserWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"User first name, last name and password can't be empty",
	)
	UserExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"This user already exists",
	)
	GroupWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Group number, course name and year can't be empty",
	)
	GroupExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"This group already exists",
	)
	FinhelpCtgWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Financial Help Category descrpition and payment can't be empty",
	)
	FinhelpCtgExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"This Financial Help Category already exists",
	)
	FinhelpStageWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Financial Help Stage descrpition can't be empty",
	)
	FinhelpStageExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"This Financial Help Stage already exists",
	)
	StudentWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Student first name, last name and phone number can't be empty",
	)
	StudentExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"This student already exists",
	)
	FinhelpProcWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Financial Help Process student, category and stage can't be empty",
	)
	FinhelpProcExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"This Financial Help Process already exists",
	)
	BuildingWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Building name and address can't be empty",
	)
	BuildingExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"This Building already exists",
	)
	CabinetTypeWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Cabinet Type name can't be empty",
	)
	CabinetTypeExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"This Cabinet Type already exists",
	)
	ClassTypeWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Class Type name can't be empty",
	)
	ClassTypeExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"This Class Type already exists",
	)
	CourseTypeWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Course Type name can't be empty",
	)
	CourseTypeExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"This Course Type already exists",
	)
	ProfWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Prof first name, last name, phone number and email can't be empty",
	)
	ProfExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"This prof already exists",
	)
)
