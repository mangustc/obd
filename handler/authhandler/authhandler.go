package authhandler

import (
	"net/http"

	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/authview"
)

func NewAuthHandler(js handler.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: js,
	}
}

type AuthHandler struct {
	AuthService handler.AuthService
}

func (jh *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	defer util.RespondHTTP(w, &code, &out)

	util.RenderComponent(r, &out, authview.Auth())
}

func (jh *AuthHandler) AuthPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	defer util.RespondHTTP(w, &code, &out)

	util.RenderComponent(r, &out, authview.AuthPage())
}
