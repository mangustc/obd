package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	E "github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema"
	"github.com/mangustc/obd/schema/jobschema"
	"github.com/mangustc/obd/schema/sessionschema"
	"github.com/mangustc/obd/schema/userschema"
	"github.com/mangustc/obd/view"
	"github.com/mattn/go-sqlite3"
)

func GetStringFromForm(r *http.Request, name string) string {
	return r.Form.Get(name)
}

func GetIntFromForm(r *http.Request, name string) (int, error) {
	i, err := strconv.Atoi(r.Form.Get(name))
	return i, err
}

func GetBoolFromForm(r *http.Request, name string) (bool, error) {
	str := r.Form.Get(name)
	if str == "on" {
		return true, nil
	} else if str == "off" {
		return false, nil
	}
	return false, fmt.Errorf("Value is not bool. Value: %s.", str)
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func ComponentBytes(component templ.Component, r *http.Request) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := component.Render(r.Context(), buf)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func RenderComponent(r *http.Request, writeBytes *[]byte, component templ.Component) error {
	if writeBytes == nil {
		writeBytes = &[]byte{}
	}
	out, err := ComponentBytes(component, r)
	if err != nil {
		return err
	}
	*writeBytes = append(*writeBytes, out...)
	return nil
}

func RespondHTTP(w http.ResponseWriter, r *http.Request, msg **msg.Msg, out *[]byte) {
	if msg == nil {
		panic("Code should not be nil")
	}
	// http.StatusOK is written to header by default
	if (*msg).MsgCode != http.StatusOK {
		w.WriteHeader((*msg).MsgCode)
	}
	RenderMsg(w, r, out, *msg)
	w.Write(*out)
}

func InitHTMLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	r.ParseForm()
}

func GetJobBySessionCookie(
	w http.ResponseWriter,
	r *http.Request,
	getSession func(*sessionschema.SessionGet) (*sessionschema.SessionDB, error),
	getUser func(*userschema.UserGet) (*userschema.UserDB, error),
	getJob func(*jobschema.JobGet) (*jobschema.JobDB, error),
) (*jobschema.JobDB, error) {
	sessionUUID, err := GetUserSessionCookieValue(w, r)
	if err != nil {
		return &jobschema.JobDB{}, err
	}

	sessionDB, err := getSession(&sessionschema.SessionGet{
		SessionUUID: sessionUUID,
	})
	if err != nil {
		DeleteUserSessionCookie(w)
		return &jobschema.JobDB{}, E.ErrNotFound
	}

	userDB, err := getUser(&userschema.UserGet{
		UserID: sessionDB.UserID,
	})
	if err != nil {
		DeleteUserSessionCookie(w)
		return &jobschema.JobDB{}, E.ErrNotFound
	}

	jobDB, err := getJob(&jobschema.JobGet{
		JobID: userDB.JobID,
	})
	if err != nil {
		DeleteUserSessionCookie(w)
		return &jobschema.JobDB{}, E.ErrNotFound
	}

	return jobDB, nil
}

func SetUserSessionCookie(w http.ResponseWriter, sessionUUID string) {
	cookie := http.Cookie{
		Name:     "session",
		Value:    sessionUUID,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
}

func GetUserSessionCookieValue(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return "", E.ErrNotFound
		}
		DeleteUserSessionCookie(w)
		return "", E.ErrInternalServer
	}
	sessionUUID := cookie.Value
	err = uuid.Validate(sessionUUID)
	if err != nil {
		DeleteUserSessionCookie(w)
		return "", E.ErrNotFound
	}
	return sessionUUID, nil
}

func DeleteUserSessionCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	}
	http.SetCookie(w, &cookie)
}

func IsErrorSQL(err error, sqlErr error) bool {
	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.Code == sqlErr {
			return true
		}
	}
	return false
}

func GetUintFromString(str string) (uint, error) {
	ui, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		return 0, E.ErrUnprocessableEntity
	}
	return uint(ui), nil
}

func GetFloatFromString(str string) (float32, error) {
	f64, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, E.ErrUnprocessableEntity
	}
	return float32(f64), nil
}

func IsZero(v any) bool {
	vr := reflect.ValueOf(v)
	return vr.IsZero()
}

func GetCodeByErr(err error) (int, string) {
	switch err {
	case E.ErrInternalServer:
		return http.StatusInternalServerError, fmt.Sprintf("Internal server error (%s)", err.Error())
	case E.ErrUnprocessableEntity:
		return http.StatusUnprocessableEntity, fmt.Sprintf("Unprocessable Entity (%s)", err.Error())
	case E.ErrNotFound:
		return http.StatusNotFound, fmt.Sprintf("Not Found (%s)", err.Error())
	case E.ErrUnauthorized:
		return http.StatusUnauthorized, fmt.Sprintf("Unauthorized (%s)", err.Error())
	default:
		panic("Random error?")
	}
}

func RenderMsg(w http.ResponseWriter, r *http.Request, writeBytes *[]byte, msg *msg.Msg) {
	if msg.MsgNotificationType == schema.SuccessNotification {
		RenderComponent(r, writeBytes, view.ErrorIndexOOB(msg.MsgNotificationType, msg.MsgStr))
	} else if msg.MsgNotificationType != schema.NoNotification {
		w.Header().Add("HX-Retarget", "#notifications")
		w.Header().Add("HX-Reswap", "innerHTML")
		RenderComponent(r, writeBytes, view.ErrorIndex(msg.MsgNotificationType, msg.MsgStr))
	}
}
