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
		"Успешно выполнено",
	)
	InternalServerError = NewMsg(
		http.StatusInternalServerError,
		schema.ErrorNotification,
		"Внутренняя ошибка сервера, попробуйте перезагрузить страницу",
	)
	AuthWrongPassword = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Неверный пароль, попробуйте еще раз",
	)
	AuthSuccessLogin = NewMsg(
		http.StatusOK,
		schema.SuccessNotification,
		"Успешный вход в систему",
	)
	AuthSuccessLogout = NewMsg(
		http.StatusOK,
		schema.SuccessNotification,
		"Успешный выход из системы",
	)
	Unauthorized = NewMsg(
		http.StatusUnauthorized,
		schema.ErrorNotification,
		"Несанкционированный доступ",
	)
	JobNameEmpty = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Имя должности не может быть пустым",
	)
	JobExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такая должность уже существует",
	)
	UserWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Фамилия, имя и пароль сотрудника не могут быть пустыми",
	)
	UserExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такой сотрудник уже существует",
	)
	GroupWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Номер, год поступления и название направления группы не могут быть пустыми",
	)
	GroupExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такая группа уже существует",
	)
	FinhelpCtgWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Описание и выплата категории не могут быть пустыми",
	)
	FinhelpCtgExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такая категория уже существует",
	)
	FinhelpStageWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Описание этапа не может быть пустым",
	)
	FinhelpStageExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такой этап уже существует",
	)
	StudentWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Фамилия, имя и номер телефона студента не могут быть пустыми",
	)
	StudentExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такой студент уже существует",
	)
	FinhelpProcWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Студент, категория и этап процесса материальной помощи не могут быть пустыми",
	)
	FinhelpProcExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такой процесс уже существует",
	)
	BuildingWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Название и адрес корпуса не могут быть пустыми",
	)
	BuildingExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такое здание уже существует",
	)
	CabinetTypeWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Название типа кабинета не может быть пустым",
	)
	CabinetTypeExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такой тип кабинета уже существует",
	)
	ClassTypeWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Название типа пары не может быть пустым",
	)
	ClassTypeExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такой тип пары уже существует",
	)
	CourseTypeWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Название типа дисциплины не может быть пустым",
	)
	CourseTypeExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такой тип дисциплины уже существует",
	)
	ProfWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Фамилия, имя, номер телефона и почта преподавателя не могут быть пустыми",
	)
	ProfExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такой преподаватель уже существует",
	)
	CabinetWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Тип кабинета и корпус не могут быть пустыми",
	)
	CabinetExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такой кабинет уже существует",
	)
	CourseWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Тип дисциплинны и название не могут быть пустыми",
	)
	CourseExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такая дисциплина уже существует",
	)
	ClassWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Дата пары должна быть формата ГГГГ-ММ-ДД и её номер в отрезке [1, 7]",
	)
	ClassExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такая пара уже существует",
	)
	PerfWrong = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Оценка успеваемости должна быть в отрезке [0, 5]",
	)
	PerfExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Успеваемость студента по этой дисциплине уже существует",
	)
	SkipWrong  = InternalServerError
	SkipExists = NewMsg(
		http.StatusUnprocessableEntity,
		schema.AlertNotification,
		"Такой пропуск уже стоит",
	)
)
