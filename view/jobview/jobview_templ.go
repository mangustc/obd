// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package jobview

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/mangustc/obd/schema"
	"github.com/mangustc/obd/schema/jobschema"
	"github.com/mangustc/obd/view"
)

const (
	pageTitle       = "Job"
	tableTitle      = pageTitle
	insertFormTitle = tableTitle
	getPOSTURL      = "/api/job/getjobs"
	insertPOSTURL   = "/api/job/insertjob"
	editPOSTURL     = "/api/job/editjob"
	updatePOSTURL   = "/api/job/updatejob"
	deletePOSTURL   = "/api/job/deletejob"
	jobTN           = "Job"
	bodyVals        = `{"` + jobTN + `ID": %d}`
	getjobPOSTURL   = "/api/job"
)

var (
	taJobName               = schema.NewTA(jobTN+"Name", "Job Name", schema.StringInput)
	taJobAccessUser         = schema.NewTA(jobTN+"AccessUser", "User", schema.BooleanInput)
	taJobAccessJob          = schema.NewTA(jobTN+"AccessJob", "Job", schema.BooleanInput)
	taJobAccessStudent      = schema.NewTA(jobTN+"AccessStudent", "Student", schema.BooleanInput)
	taJobAccessGroup        = schema.NewTA(jobTN+"AccessGroup", "Group", schema.BooleanInput)
	taJobAccessFinhelpCtg   = schema.NewTA(jobTN+"AccessFinhelpCtg", "FinhelpCtg", schema.BooleanInput)
	taJobAccessFinhelpStage = schema.NewTA(jobTN+"AccessFinhelpStage", "FinhelpStage", schema.BooleanInput)
	taJobAccessFinhelpProc  = schema.NewTA(jobTN+"AccessFinhelpProc", "FinhelpProc", schema.BooleanInput)
	taJobAccessBuilding     = schema.NewTA(jobTN+"AccessBuilding", "Building", schema.BooleanInput)
	taJobAccessCabinetType  = schema.NewTA(jobTN+"AccessCabinetType", "CabinetType", schema.BooleanInput)
	taJobAccessCabinet      = schema.NewTA(jobTN+"AccessCabinet", "Cabinet", schema.BooleanInput)
	taJobAccessClassType    = schema.NewTA(jobTN+"AccessClassType", "ClassType", schema.BooleanInput)
	taJobAccessProf         = schema.NewTA(jobTN+"AccessProf", "Prof", schema.BooleanInput)
	taJobAccessCourseType   = schema.NewTA(jobTN+"AccessCourseType", "CourseType", schema.BooleanInput)
	taJobAccessCourse       = schema.NewTA(jobTN+"AccessCourse", "Course", schema.BooleanInput)
	taJobAccessPerf         = schema.NewTA(jobTN+"AccessPerf", "Perf", schema.BooleanInput)
	taJobAccessSkip         = schema.NewTA(jobTN+"AccessSkip", "Skip", schema.BooleanInput)
	taJobAccessClass        = schema.NewTA(jobTN+"AccessClass", "Class", schema.BooleanInput)
)

func getTableHeaders() []*schema.TableHeaderColumn {
	return []*schema.TableHeaderColumn{
		schema.NewTableHeaderColumn(taJobName.TATitle, 30),
		schema.NewTableHeaderColumn(taJobAccessUser.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessJob.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessStudent.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessGroup.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessFinhelpCtg.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessFinhelpStage.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessFinhelpProc.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessBuilding.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessCabinetType.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessCabinet.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessClassType.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessProf.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessCourseType.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessCourse.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessPerf.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessSkip.TATitle, 4),
		schema.NewTableHeaderColumn(taJobAccessClass.TATitle, 4),
	}
}

func getInsertFormInputs() []*schema.Input {
	return []*schema.Input{
		schema.NewInput(taJobName.TATitle, taJobName.TAName, taJobName.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessUser.TATitle, taJobAccessUser.TAName, taJobAccessUser.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessJob.TATitle, taJobAccessJob.TAName, taJobAccessJob.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessStudent.TATitle, taJobAccessStudent.TAName, taJobAccessStudent.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessGroup.TATitle, taJobAccessGroup.TAName, taJobAccessGroup.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessFinhelpCtg.TATitle, taJobAccessFinhelpCtg.TAName, taJobAccessFinhelpCtg.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessFinhelpStage.TATitle, taJobAccessFinhelpStage.TAName, taJobAccessFinhelpStage.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessFinhelpProc.TATitle, taJobAccessFinhelpProc.TAName, taJobAccessFinhelpProc.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessBuilding.TATitle, taJobAccessBuilding.TAName, taJobAccessBuilding.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessCabinetType.TATitle, taJobAccessCabinetType.TAName, taJobAccessCabinetType.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessCabinet.TATitle, taJobAccessCabinet.TAName, taJobAccessCabinet.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessClassType.TATitle, taJobAccessClassType.TAName, taJobAccessClassType.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessProf.TATitle, taJobAccessProf.TAName, taJobAccessProf.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessCourseType.TATitle, taJobAccessCourseType.TAName, taJobAccessCourseType.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessCourse.TATitle, taJobAccessCourse.TAName, taJobAccessCourse.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessPerf.TATitle, taJobAccessPerf.TAName, taJobAccessPerf.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessSkip.TATitle, taJobAccessSkip.TAName, taJobAccessSkip.TAInputType, "", nil, nil, ""),
		schema.NewInput(taJobAccessClass.TATitle, taJobAccessClass.TAName, taJobAccessClass.TAInputType, "", nil, nil, ""),
	}
}

func getInputsFromJobDB(jobDB *jobschema.JobDB) []*schema.Input {
	return []*schema.Input{
		schema.NewInput("", taJobName.TAName, taJobName.TAInputType, jobDB.JobName, nil, nil, ""),
		schema.NewInput("", taJobAccessUser.TAName, taJobAccessUser.TAInputType, jobDB.JobAccessUser, nil, nil, ""),
		schema.NewInput("", taJobAccessJob.TAName, taJobAccessJob.TAInputType, jobDB.JobAccessJob, nil, nil, ""),
		schema.NewInput("", taJobAccessStudent.TAName, taJobAccessStudent.TAInputType, jobDB.JobAccessStudent, nil, nil, ""),
		schema.NewInput("", taJobAccessGroup.TAName, taJobAccessGroup.TAInputType, jobDB.JobAccessGroup, nil, nil, ""),
		schema.NewInput("", taJobAccessFinhelpCtg.TAName, taJobAccessFinhelpCtg.TAInputType, jobDB.JobAccessFinhelpCtg, nil, nil, ""),
		schema.NewInput("", taJobAccessFinhelpStage.TAName, taJobAccessFinhelpStage.TAInputType, jobDB.JobAccessFinhelpStage, nil, nil, ""),
		schema.NewInput("", taJobAccessFinhelpProc.TAName, taJobAccessFinhelpProc.TAInputType, jobDB.JobAccessFinhelpProc, nil, nil, ""),
		schema.NewInput("", taJobAccessBuilding.TAName, taJobAccessBuilding.TAInputType, jobDB.JobAccessBuilding, nil, nil, ""),
		schema.NewInput("", taJobAccessCabinetType.TAName, taJobAccessCabinetType.TAInputType, jobDB.JobAccessCabinetType, nil, nil, ""),
		schema.NewInput("", taJobAccessCabinet.TAName, taJobAccessCabinet.TAInputType, jobDB.JobAccessCabinet, nil, nil, ""),
		schema.NewInput("", taJobAccessClassType.TAName, taJobAccessClassType.TAInputType, jobDB.JobAccessClassType, nil, nil, ""),
		schema.NewInput("", taJobAccessProf.TAName, taJobAccessProf.TAInputType, jobDB.JobAccessProf, nil, nil, ""),
		schema.NewInput("", taJobAccessCourseType.TAName, taJobAccessCourseType.TAInputType, jobDB.JobAccessCourseType, nil, nil, ""),
		schema.NewInput("", taJobAccessCourse.TAName, taJobAccessCourse.TAInputType, jobDB.JobAccessCourse, nil, nil, ""),
		schema.NewInput("", taJobAccessPerf.TAName, taJobAccessPerf.TAInputType, jobDB.JobAccessPerf, nil, nil, ""),
		schema.NewInput("", taJobAccessSkip.TAName, taJobAccessSkip.TAInputType, jobDB.JobAccessSkip, nil, nil, ""),
		schema.NewInput("", taJobAccessClass.TAName, taJobAccessClass.TAInputType, jobDB.JobAccessClass, nil, nil, ""),
	}
}

func JobTableRowEdit(jobDB *jobschema.JobDB) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = view.TableRowEdit(getInputsFromJobDB(jobDB), fmt.Sprintf(bodyVals, jobDB.JobID), updatePOSTURL, deletePOSTURL).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func JobTableRow(jobDB *jobschema.JobDB) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = view.TableRow(getInputsFromJobDB(jobDB), fmt.Sprintf(bodyVals, jobDB.JobID), editPOSTURL, deletePOSTURL).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func JobTableRows(jobsDB []*jobschema.JobDB) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		for _, jobDB := range jobsDB {
			templ_7745c5c3_Err = JobTableRow(jobDB).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

func JobAddForm() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = view.InsertForm(insertFormTitle, insertPOSTURL, getInsertFormInputs()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func JobTable() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = view.Table(tableTitle, getPOSTURL, getTableHeaders()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Job() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = JobAddForm().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = JobTable().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func JobPage() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var8 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div hx-post=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(getjobPOSTURL)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/jobview/jobview.templ`, Line: 144, Col: 26}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"this\" hx-swap=\"outerHTML\" hx-trigger=\"load\"></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = view.Layout(pageTitle).Render(templ.WithChildren(ctx, templ_7745c5c3_Var8), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
