package buildinghandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/buildingschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/buildingview"
)

func NewBuildingHandler(ss handler.SessionService, us handler.UserService, js handler.JobService, fstages handler.BuildingService) *BuildingHandler {
	return &BuildingHandler{
		SessionService:  ss,
		UserService:     us,
		JobService:      js,
		BuildingService: fstages,
	}
}

type BuildingHandler struct {
	SessionService  handler.SessionService
	UserService     handler.UserService
	JobService      handler.JobService
	BuildingService handler.BuildingService
}

func (fsth *BuildingHandler) Building(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, buildingview.Building())
}

func (fsth *BuildingHandler) BuildingPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, buildingview.BuildingPage())
}

func (fsth *BuildingHandler) GetBuildings(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &buildingschema.BuildingsGet{}

	err = buildingschema.ValidateBuildingsGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	buildingsDB, err := fsth.BuildingService.GetBuildings(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, buildingview.BuildingTableRows(buildingsDB))
}

func (fsth *BuildingHandler) InsertBuilding(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &buildingschema.BuildingInsert{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		fsth.SessionService.GetSession,
		fsth.UserService.GetUser,
		fsth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessBuilding {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	err = buildingschema.ValidateBuildingInsert(in)
	if err != nil {
		message = msg.BuildingWrong
		logger.Error.Print(err.Error())
		return
	}

	buildingDB, err := fsth.BuildingService.InsertBuilding(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.BuildingExists
			logger.Error.Print(err.Error())
			return
		}
	}

	util.RenderComponent(r, &out, buildingview.BuildingTableRow(buildingDB))
}

func (fsth *BuildingHandler) EditBuilding(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &buildingschema.BuildingGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		fsth.SessionService.GetSession,
		fsth.UserService.GetUser,
		fsth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessBuilding {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = buildingschema.ValidateBuildingGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	buildingDB, err := fsth.BuildingService.GetBuilding(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, buildingview.BuildingTableRowEdit(buildingDB))
}

func (fsth *BuildingHandler) UpdateBuilding(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.OK
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &buildingschema.BuildingUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		fsth.SessionService.GetSession,
		fsth.UserService.GetUser,
		fsth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessBuilding {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}
	err = buildingschema.ValidateBuildingUpdate(in)
	if err != nil {
		message = msg.BuildingWrong
		logger.Error.Print(err.Error())
		return
	}

	buildingDB, err := fsth.BuildingService.UpdateBuilding(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.BuildingExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	util.RenderComponent(r, &out, buildingview.BuildingTableRow(buildingDB))
}

func (fsth *BuildingHandler) DeleteBuilding(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &buildingschema.BuildingDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		fsth.SessionService.GetSession,
		fsth.UserService.GetUser,
		fsth.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessBuilding {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = buildingschema.ValidateBuildingDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	buildingDB, err := fsth.BuildingService.DeleteBuilding(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	util.RenderComponent(r, &out, buildingview.BuildingTableRow(buildingDB))
}
