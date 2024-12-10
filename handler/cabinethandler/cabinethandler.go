package cabinethandler

import (
	"net/http"

	"github.com/mangustc/obd/errs"
	"github.com/mangustc/obd/handler"
	"github.com/mangustc/obd/logger"
	"github.com/mangustc/obd/msg"
	"github.com/mangustc/obd/schema/buildingschema"
	"github.com/mangustc/obd/schema/cabinetschema"
	"github.com/mangustc/obd/schema/cabinettypeschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view/cabinetview"
)

func NewCabinetHandler(
	ss handler.SessionService,
	us handler.UserService,
	js handler.JobService,
	cs handler.CabinetService,
	bs handler.BuildingService,
	cts handler.CabinetTypeService,
) *CabinetHandler {
	return &CabinetHandler{
		SessionService:     ss,
		UserService:        us,
		JobService:         js,
		CabinetService:     cs,
		BuildingService:    bs,
		CabinetTypeService: cts,
	}
}

type CabinetHandler struct {
	SessionService     handler.SessionService
	UserService        handler.UserService
	JobService         handler.JobService
	BuildingService    handler.BuildingService
	CabinetTypeService handler.CabinetTypeService
	CabinetService     handler.CabinetService
}

func (ch *CabinetHandler) Cabinet(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	buildingsDB, _ := ch.BuildingService.GetBuildings(&buildingschema.BuildingsGet{})
	buildingInputOptions := buildingschema.GetBuildingInputOptionsFromBuildingsDB(buildingsDB)

	cabinetTypesDB, _ := ch.CabinetTypeService.GetCabinetTypes(&cabinettypeschema.CabinetTypesGet{})
	cabinetTypeInputOptions := cabinettypeschema.GetCabinetTypeInputOptionsFromCabinetTypesDB(cabinetTypesDB)

	util.RenderComponent(r, &out, cabinetview.Cabinet(buildingInputOptions, cabinetTypeInputOptions))
}

func (ch *CabinetHandler) CabinetPage(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)

	util.RenderComponent(r, &out, cabinetview.CabinetPage())
}

func (ch *CabinetHandler) GetCabinets(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &cabinetschema.CabinetsGet{}

	err = cabinetschema.ValidateCabinetsGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	cabinetsDB, err := ch.CabinetService.GetCabinets(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	buildingsDB, _ := ch.BuildingService.GetBuildings(&buildingschema.BuildingsGet{})
	buildingInputOptions := buildingschema.GetBuildingInputOptionsFromBuildingsDB(buildingsDB)

	cabinetTypesDB, _ := ch.CabinetTypeService.GetCabinetTypes(&cabinettypeschema.CabinetTypesGet{})
	cabinetTypeInputOptions := cabinettypeschema.GetCabinetTypeInputOptionsFromCabinetTypesDB(cabinetTypesDB)

	util.RenderComponent(r, &out, cabinetview.CabinetTableRows(cabinetsDB, buildingInputOptions, cabinetTypeInputOptions))
}

func (ch *CabinetHandler) InsertCabinet(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &cabinetschema.CabinetInsert{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		ch.SessionService.GetSession,
		ch.UserService.GetUser,
		ch.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCabinet {
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
	err = cabinetschema.ValidateCabinetInsert(in)
	if err != nil {
		message = msg.CabinetWrong
		logger.Error.Print(err.Error())
		return
	}

	cabinetDB, err := ch.CabinetService.InsertCabinet(in)
	if err != nil {
		if err != errs.ErrUnprocessableEntity {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.CabinetExists
			logger.Error.Print(err.Error())
			return
		}
	}

	buildingsDB, _ := ch.BuildingService.GetBuildings(&buildingschema.BuildingsGet{})
	buildingInputOptions := buildingschema.GetBuildingInputOptionsFromBuildingsDB(buildingsDB)

	cabinetTypesDB, _ := ch.CabinetTypeService.GetCabinetTypes(&cabinettypeschema.CabinetTypesGet{})
	cabinetTypeInputOptions := cabinettypeschema.GetCabinetTypeInputOptionsFromCabinetTypesDB(cabinetTypesDB)

	util.RenderComponent(r, &out, cabinetview.CabinetTableRow(cabinetDB, buildingInputOptions, cabinetTypeInputOptions))
}

func (ch *CabinetHandler) EditCabinet(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &cabinetschema.CabinetGet{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		ch.SessionService.GetSession,
		ch.UserService.GetUser,
		ch.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCabinet {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = cabinetschema.ValidateCabinetGet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	cabinetDB, err := ch.CabinetService.GetCabinet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	buildingsDB, _ := ch.BuildingService.GetBuildings(&buildingschema.BuildingsGet{})
	buildingInputOptions := buildingschema.GetBuildingInputOptionsFromBuildingsDB(buildingsDB)

	cabinetTypesDB, _ := ch.CabinetTypeService.GetCabinetTypes(&cabinettypeschema.CabinetTypesGet{})
	cabinetTypeInputOptions := cabinettypeschema.GetCabinetTypeInputOptionsFromCabinetTypesDB(cabinetTypesDB)

	util.RenderComponent(r, &out, cabinetview.CabinetTableRowEdit(cabinetDB, buildingInputOptions, cabinetTypeInputOptions))
}

func (ch *CabinetHandler) UpdateCabinet(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &cabinetschema.CabinetUpdate{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		ch.SessionService.GetSession,
		ch.UserService.GetUser,
		ch.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCabinet {
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
	err = cabinetschema.ValidateCabinetUpdate(in)
	if err != nil {
		message = msg.CabinetWrong
		logger.Error.Print(err.Error())
		return
	}

	cabinetDB, err := ch.CabinetService.UpdateCabinet(in)
	if err != nil {
		if err == errs.ErrUnprocessableEntity {
			message = msg.CabinetExists
			logger.Error.Print(err.Error())
			return
		} else {
			message = msg.InternalServerError
			logger.Error.Print(err.Error())
			return

		}
	}

	buildingsDB, _ := ch.BuildingService.GetBuildings(&buildingschema.BuildingsGet{})
	buildingInputOptions := buildingschema.GetBuildingInputOptionsFromBuildingsDB(buildingsDB)

	cabinetTypesDB, _ := ch.CabinetTypeService.GetCabinetTypes(&cabinettypeschema.CabinetTypesGet{})
	cabinetTypeInputOptions := cabinettypeschema.GetCabinetTypeInputOptionsFromCabinetTypesDB(cabinetTypesDB)

	util.RenderComponent(r, &out, cabinetview.CabinetTableRow(cabinetDB, buildingInputOptions, cabinetTypeInputOptions))
}

func (ch *CabinetHandler) DeleteCabinet(w http.ResponseWriter, r *http.Request) {
	var err error

	util.InitHTMLHandler(w, r)
	var message *msg.Msg = msg.Nothing
	var out []byte
	defer util.RespondHTTP(w, r, &message, &out)
	in := &cabinetschema.CabinetDelete{}

	sessionJobDB, _, err := util.GetJobBySessionCookie(
		w, r,
		ch.SessionService.GetSession,
		ch.UserService.GetUser,
		ch.JobService.GetJob,
	)
	if !sessionJobDB.JobAccessCabinet {
		err := errs.ErrUnauthorized
		message = msg.Unauthorized
		logger.Error.Print(err.Error())
		return
	}

	err = util.ParseStructFromForm(r, in)
	err = cabinetschema.ValidateCabinetDelete(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	cabinetDB, err := ch.CabinetService.DeleteCabinet(in)
	if err != nil {
		message = msg.InternalServerError
		logger.Error.Print(err.Error())
		return
	}

	buildingsDB, _ := ch.BuildingService.GetBuildings(&buildingschema.BuildingsGet{})
	buildingInputOptions := buildingschema.GetBuildingInputOptionsFromBuildingsDB(buildingsDB)

	cabinetTypesDB, _ := ch.CabinetTypeService.GetCabinetTypes(&cabinettypeschema.CabinetTypesGet{})
	cabinetTypeInputOptions := cabinettypeschema.GetCabinetTypeInputOptionsFromCabinetTypesDB(cabinetTypesDB)

	util.RenderComponent(r, &out, cabinetview.CabinetTableRow(cabinetDB, buildingInputOptions, cabinetTypeInputOptions))
}
