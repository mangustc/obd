package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	E "github.com/mangustc/obd/errs"
	"github.com/mattn/go-sqlite3"
)

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

func RespondHTTP(w http.ResponseWriter, code *int, out *[]byte) {
	if code == nil {
		panic("Code should not be nil")
	}
	// http.StatusOK is written to header by default
	if *code != http.StatusOK {
		w.WriteHeader(*code)
	}
	w.Write(*out)
}

func InitHTMLHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	r.ParseForm()
}

func SetUserSessionCookie(w http.ResponseWriter, sessionUUID uuid.UUID) {
	cookie := http.Cookie{
		Name:     "session",
		Value:    sessionUUID.String(),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
}

func GetUserSessionCookieValue(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	sessionUUID, err := getUserSessionCookieValue(r)
	if err != nil {
		switch err {
		case E.ErrNotFound:
			return uuid.UUID{}, err
		case E.ErrUnprocessableEntity:
			DeleteUserSessionCookie(w)
			return uuid.UUID{}, E.ErrNotFound
		default:
			return uuid.UUID{}, E.ErrInternalServer
		}
	}
	return sessionUUID, nil
}

func getUserSessionCookieValue(r *http.Request) (uuid.UUID, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return uuid.UUID{}, E.ErrNotFound
		}
		return uuid.UUID{}, E.ErrInternalServer
	}
	sessionUUID, err := uuid.Parse(cookie.Value)
	if err != nil {
		return uuid.UUID{}, E.ErrUnprocessableEntity
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
