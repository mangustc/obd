package handler

import (
	"net/http"

	"github.com/mangustc/obd/schema/jobschema"
	"github.com/mangustc/obd/schema/userschema"
	"github.com/mangustc/obd/util"
	"github.com/mangustc/obd/view"
)

type (
	JobService interface {
		InsertJob(data *jobschema.JobInsert) (jobDB *jobschema.JobDB, err error)
		UpdateJob(data *jobschema.JobUpdate) (jobDB *jobschema.JobDB, err error)
		DeleteJob(data *jobschema.JobDelete) (jobDB *jobschema.JobDB, err error)
		GetJob(data *jobschema.JobGet) (jobDB *jobschema.JobDB, err error)
		GetJobs(data *jobschema.JobsGet) (jobsDB []*jobschema.JobDB, err error)
	}
	AuthService interface{}
	UserService interface {
		InsertUser(data *userschema.UserInsert) (userDB *userschema.UserDB, err error)
		UpdateUser(data *userschema.UserUpdate) (userDB *userschema.UserDB, err error)
		DeleteUser(data *userschema.UserDelete) (userDB *userschema.UserDB, err error)
		GetUser(data *userschema.UserGet) (userDB *userschema.UserDB, err error)
		GetUsers(data *userschema.UsersGet) (usersDB []*userschema.UserDB, err error)
	}
)

func Default(w http.ResponseWriter, r *http.Request) {
	// var err error

	util.InitHTMLHandler(w, r)
	var code int = http.StatusOK
	var out []byte
	defer util.RespondHTTP(w, &code, &out)

	util.RenderComponent(r, &out, view.Layout("OBD"))
}
