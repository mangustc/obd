// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package classview

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/mangustc/obd/schema"
	"github.com/mangustc/obd/schema/classschema"
	"github.com/mangustc/obd/view"
)

const (
	pageTitle       = "Class"
	tableTitle      = pageTitle
	insertFormTitle = tableTitle
	getPOSTURL      = "/api/class/getclasss"
	insertPOSTURL   = "/api/class/insertclass"
	editPOSTURL     = "/api/class/editclass"
	updatePOSTURL   = "/api/class/updateclass"
	deletePOSTURL   = "/api/class/deleteclass"
	classTN         = "Class"
	classTypeTN     = "ClassType"
	profTN          = "Prof"
	cabinetTN       = "Cabinet"
	courseTN        = "Course"
	groupTN         = "Group"
	bodyVals        = `{"` + classTN + `ID": %d}`
	getclassPOSTURL = "/api/class"
)

var (
	taClassTypeID = schema.NewTA(classTypeTN+"ID", "Class Type", schema.OptionInput)
	taProfID      = schema.NewTA(profTN+"ID", "Professor", schema.OptionInput)
	taCabinetID   = schema.NewTA(cabinetTN+"ID", "Cabinet", schema.OptionInput)
	taCourseID    = schema.NewTA(courseTN+"ID", "Course", schema.OptionInput)
	taGroupID     = schema.NewTA(groupTN+"ID", "Group", schema.OptionInput)
	taClassStart  = schema.NewTA(classTN+"Start", "Class Date", schema.StringInput)
	taClassNumber = schema.NewTA(classTN+"Number", "Class Number", schema.NumberInput)
)

func getTableHeaders() []*schema.TableHeaderColumn {
	return []*schema.TableHeaderColumn{
		schema.NewTableHeaderColumn(taClassTypeID.TATitle, 15),
		schema.NewTableHeaderColumn(taProfID.TATitle, 15),
		schema.NewTableHeaderColumn(taCabinetID.TATitle, 15),
		schema.NewTableHeaderColumn(taCourseID.TATitle, 15),
		schema.NewTableHeaderColumn(taGroupID.TATitle, 15),
		schema.NewTableHeaderColumn(taClassStart.TATitle, 15),
		schema.NewTableHeaderColumn(taClassNumber.TATitle, 10),
	}
}

func getInsertFormInputs(
	classTypeInputOptions []*schema.InputOption,
	profInputOptions []*schema.InputOption,
	cabinetInputOptions []*schema.InputOption,
	courseInputOptions []*schema.InputOption,
	groupInputOptions []*schema.InputOption,
) []*schema.Input {
	return []*schema.Input{
		schema.NewInput(taClassTypeID.TATitle, taClassTypeID.TAName, taClassTypeID.TAInputType, "", nil, classTypeInputOptions, ""),
		schema.NewInput(taProfID.TATitle, taProfID.TAName, taProfID.TAInputType, "", nil, profInputOptions, ""),
		schema.NewInput(taCabinetID.TATitle, taCabinetID.TAName, taCabinetID.TAInputType, "", nil, cabinetInputOptions, ""),
		schema.NewInput(taCourseID.TATitle, taCourseID.TAName, taCourseID.TAInputType, "", nil, courseInputOptions, ""),
		schema.NewInput(taGroupID.TATitle, taGroupID.TAName, taGroupID.TAInputType, "", nil, groupInputOptions, ""),
		schema.NewInput(taClassStart.TATitle, taClassStart.TAName, taClassStart.TAInputType, "", nil, nil, ""),
		schema.NewInput(taClassNumber.TATitle, taClassNumber.TAName, taClassNumber.TAInputType, "", nil, nil, ""),
	}
}

func getInputsFromClassDB(classDB *classschema.ClassDB,
	classTypeInputOptions []*schema.InputOption,
	profInputOptions []*schema.InputOption,
	cabinetInputOptions []*schema.InputOption,
	courseInputOptions []*schema.InputOption,
	groupInputOptions []*schema.InputOption,
) []*schema.Input {
	return []*schema.Input{
		schema.NewInput("", taClassTypeID.TAName, taClassTypeID.TAInputType, classDB.ClassTypeID, nil, classTypeInputOptions, fmt.Sprint(classDB.ClassTypeID)),
		schema.NewInput("", taProfID.TAName, taProfID.TAInputType, classDB.ProfID, nil, profInputOptions, fmt.Sprint(classDB.ProfID)),
		schema.NewInput("", taCabinetID.TAName, taCabinetID.TAInputType, classDB.CabinetID, nil, cabinetInputOptions, fmt.Sprint(classDB.CabinetID)),
		schema.NewInput("", taCourseID.TAName, taCourseID.TAInputType, classDB.CourseID, nil, courseInputOptions, fmt.Sprint(classDB.CourseID)),
		schema.NewInput("", taGroupID.TAName, taGroupID.TAInputType, classDB.GroupID, nil, groupInputOptions, fmt.Sprint(classDB.GroupID)),
		schema.NewInput("", taClassStart.TAName, taClassStart.TAInputType, classDB.ClassStart, nil, nil, ""),
		schema.NewInput("", taClassNumber.TAName, taClassNumber.TAInputType, classDB.ClassNumber, nil, nil, ""),
	}
}

func ClassTableRowEdit(classDB *classschema.ClassDB,
	classTypeInputOptions []*schema.InputOption,
	profInputOptions []*schema.InputOption,
	cabinetInputOptions []*schema.InputOption,
	courseInputOptions []*schema.InputOption,
	groupInputOptions []*schema.InputOption,
) templ.Component {
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
		templ_7745c5c3_Err = view.TableRowEdit(getInputsFromClassDB(classDB,
			classTypeInputOptions,
			profInputOptions,
			cabinetInputOptions,
			courseInputOptions,
			groupInputOptions,
		), fmt.Sprintf(bodyVals, classDB.ClassID), updatePOSTURL, deletePOSTURL).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ClassTableRow(classDB *classschema.ClassDB,
	classTypeInputOptions []*schema.InputOption,
	profInputOptions []*schema.InputOption,
	cabinetInputOptions []*schema.InputOption,
	courseInputOptions []*schema.InputOption,
	groupInputOptions []*schema.InputOption,
) templ.Component {
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
		templ_7745c5c3_Err = view.TableRow(getInputsFromClassDB(classDB,
			classTypeInputOptions,
			profInputOptions,
			cabinetInputOptions,
			courseInputOptions,
			groupInputOptions,
		), fmt.Sprintf(bodyVals, classDB.ClassID), editPOSTURL, deletePOSTURL).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ClassTableRows(classsDB []*classschema.ClassDB,
	classTypeInputOptions []*schema.InputOption,
	profInputOptions []*schema.InputOption,
	cabinetInputOptions []*schema.InputOption,
	courseInputOptions []*schema.InputOption,
	groupInputOptions []*schema.InputOption,
) templ.Component {
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
		for _, classDB := range classsDB {
			templ_7745c5c3_Err = ClassTableRow(classDB,
				classTypeInputOptions,
				profInputOptions,
				cabinetInputOptions,
				courseInputOptions,
				groupInputOptions,
			).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

func Class(
	classTypeInputOptions []*schema.InputOption,
	profInputOptions []*schema.InputOption,
	cabinetInputOptions []*schema.InputOption,
	courseInputOptions []*schema.InputOption,
	groupInputOptions []*schema.InputOption,
) templ.Component {
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
		templ_7745c5c3_Err = view.InsertForm(insertFormTitle, insertPOSTURL, getInsertFormInputs(
			classTypeInputOptions,
			profInputOptions,
			cabinetInputOptions,
			courseInputOptions,
			groupInputOptions,
		)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = view.Table(tableTitle, getPOSTURL, getTableHeaders()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ClassPage() templ.Component {
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
		templ_7745c5c3_Var6 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(getclassPOSTURL)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/classview/classview.templ`, Line: 157, Col: 28}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"this\" hx-swap=\"outerHTML\" hx-trigger=\"load\"></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = view.Layout(pageTitle).Render(templ.WithChildren(ctx, templ_7745c5c3_Var6), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
