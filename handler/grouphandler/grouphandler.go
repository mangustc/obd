package grouphandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/groupschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/groupview"
)

func NewGroupHandler(ss handler.SessionService, us handler.UserService, js handler.JobService, grs handler.GroupService) *GroupHandler {
	return &GroupHandler{
		SessionService: ss,
		UserService:    us,
		JobService:     js,
		GroupService:   grs,
	}
}

type GroupHandler struct {
	SessionService handler.SessionService
	UserService    handler.UserService
	JobService     handler.JobService
	GroupService   handler.GroupService
}

func (grh *GroupHandler) Group(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, groupview.Group())
}

func (grh *GroupHandler) GroupPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, groupview.GroupPage())
}

func (grh *GroupHandler) GetGroups(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &groupschema.GroupsGet{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		grh.SessionService.GetSession,
		grh.UserService.GetUser,
		grh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessGroup {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = groupschema.ValidateGroupsGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	groupsDB, err := grh.GroupService.GetGroups(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, groupview.GroupTableRows(groupsDB))
}

func (grh *GroupHandler) InsertGroup(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &groupschema.GroupInsert{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		grh.SessionService.GetSession,
		grh.UserService.GetUser,
		grh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessGroup {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	in.GroupYear, err = util.GetIntFromForm(r, "GroupYear")
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	in.GroupNumber = util.GetStringFromForm(r, "GroupNumber")
	in.GroupCourseName = util.GetStringFromForm(r, "GroupCourseName")
	err = groupschema.ValidateGroupInsert(in)
	if err != nil {
		message = msg.GroupWrong
		logger.Error.Print(err.Error())
		return
	}

	groupDB, err := grh.GroupService.InsertGroup(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.GroupExists
			logger.Error.Print(err.Error())
			return
		}
	}

	util.RenderComponent(r, &out, groupview.GroupTableRow(groupDB))
}

func (grh *GroupHandler) EditGroup(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &groupschema.GroupGet{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		grh.SessionService.GetSession,
		grh.UserService.GetUser,
		grh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessGroup {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	in.GroupID, err = util.GetIntFromForm(r, "GroupID")
	err = groupschema.ValidateGroupGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	groupDB, err := grh.GroupService.GetGroup(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, groupview.GroupTableRowEdit(groupDB))
}

func (grh *GroupHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &groupschema.GroupUpdate{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		grh.SessionService.GetSession,
		grh.UserService.GetUser,
		grh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessGroup {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	in.GroupID, err = util.GetIntFromForm(r, "GroupID")
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	in.GroupYear, err = util.GetIntFromForm(r, "GroupYear")
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	in.GroupNumber = util.GetStringFromForm(r, "GroupNumber")
	in.GroupCourseName = util.GetStringFromForm(r, "GroupCourseName")
	err = groupschema.ValidateGroupUpdate(in)
	if err != nil {
		message = msg.GroupWrong
		logger.Error.Print(err.Error())
		return
	}

	groupDB, err := grh.GroupService.UpdateGroup(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.GroupExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	util.RenderComponent(r, &out, groupview.GroupTableRow(groupDB))
}

func (grh *GroupHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &groupschema.GroupDelete{}

	sessionJobDB, err := util.GetJobBySessionCookie(
		w, r,
		grh.SessionService.GetSession,
		grh.UserService.GetUser,
		grh.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessGroup {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	in.GroupID, err = util.GetIntFromForm(r, "GroupID")
	err = groupschema.ValidateGroupDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	groupDB, err := grh.GroupService.DeleteGroup(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, groupview.GroupTableRow(groupDB))
}
