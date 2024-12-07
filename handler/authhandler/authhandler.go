package authhandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/sessionschema"
	"github.com/mangustc/obd/schema/userschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view"
	"github.com/mangustc/obd/view/authview"
)

func NewAuthHandler(sh handler.SessionService, uh handler.UserService) *AuthHandler {
	return &AuthHandler{
		SessionService: sh,
		UserService:    uh,
	}
}

type AuthHandler struct {
	SessionService handler.SessionService
	UserService    handler.UserService
}

func (auh *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, authview.Auth())
}

func (auh *AuthHandler) AuthPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, authview.AuthPage())
}

func (auh *AuthHandler) UserInput(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	usersGet := &userschema.UsersGet{}
	usersDB, _ := auh.UserService.GetUsers(usersGet)
	userInputOptions := userschema.GetUserInputOptionsFromUsersDB(usersDB)

	util.RenderComponent(r, &out, authview.UserInput(userInputOptions))
}

func (auh *AuthHandler) AuthLogin(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &userschema.UserGet{}

	in.UserID, err = util.GetIntFromForm(r, "UserID")
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	err = userschema.ValidateUserGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	userDB, err := auh.UserService.GetUser(in)
	if err != nil {
		if err == errs.ErrNotFound {
			err = errs.ErrInternalServer
		}
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	userPassword := util.GetStringFromForm(r, "UserPassword")
	if userDB.UserPassword != userPassword {
		err := errs.ErrUnprocessableEntity
		message = msg.AuthWrongPassword
		logger.Error.Print(err.Error())
		return
	}

	sessionInsert := &sessionschema.SessionInsert{
		UserID: in.UserID,
	}
	sessionDB, err := auh.SessionService.InsertSession(sessionInsert)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	util.SetUserSessionCookie(w, sessionDB.SessionUUID)

	util.RenderComponent(r, &out, view.ClearMain())
	util.RenderComponent(r, &out, view.Navigation())
	message = msg.AuthSuccessLogin
}

func (auh *AuthHandler) AuthLogout(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.DeleteUserSessionCookie(w)

	util.RenderComponent(r, &out, view.ClearMain())
	util.RenderComponent(r, &out, view.Navigation())
	message = msg.AuthSuccessLogout
}
