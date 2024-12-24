// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package profview

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/mangustc/obd/schema"
	"github.com/mangustc/obd/schema/profschema"
	"github.com/mangustc/obd/view"
)

const (
	pageTitle       = "Преподаватели"
	tableTitle      = pageTitle
	insertFormTitle = "Преподавателя"
	getPOSTURL      = "/api/prof/getprofs"
	insertPOSTURL   = "/api/prof/insertprof"
	editPOSTURL     = "/api/prof/editprof"
	updatePOSTURL   = "/api/prof/updateprof"
	deletePOSTURL   = "/api/prof/deleteprof"
	profTN          = "Prof"
	bodyVals        = `{"` + profTN + `ID": %d}`
	getprofPOSTURL  = "/api/prof"
)

var (
	taProfLastname    = schema.NewTA(profTN+"Lastname", "Фамилия", schema.StringInput)
	taProfFirstname   = schema.NewTA(profTN+"Firstname", "Имя", schema.StringInput)
	taProfMiddlename  = schema.NewTA(profTN+"Middlename", "Отчество", schema.StringInput)
	taProfPhoneNumber = schema.NewTA(profTN+"PhoneNumber", "Номер телефона", schema.StringInput)
	taProfEmail       = schema.NewTA(profTN+"Email", "Почта", schema.StringInput)
	taProfIsHidden    = schema.NewTA(profTN+"IsHidden", "Скрыт", schema.BooleanInput)
)

func getTableHeaders() []*schema.TableHeaderColumn {
	return []*schema.TableHeaderColumn{
		schema.NewTableHeaderColumn(taProfLastname.TATitle, 15),
		schema.NewTableHeaderColumn(taProfFirstname.TATitle, 15),
		schema.NewTableHeaderColumn(taProfMiddlename.TATitle, 15),
		schema.NewTableHeaderColumn(taProfPhoneNumber.TATitle, 10),
		schema.NewTableHeaderColumn(taProfEmail.TATitle, 10),
		schema.NewTableHeaderColumn(taProfIsHidden.TATitle, 5),
	}
}

func getInsertFormInputs() []*schema.Input {
	return []*schema.Input{
		schema.NewInput(taProfLastname.TATitle, taProfLastname.TAName, taProfLastname.TAInputType, "", nil, nil, ""),
		schema.NewInput(taProfFirstname.TATitle, taProfFirstname.TAName, taProfFirstname.TAInputType, "", nil, nil, ""),
		schema.NewInput(taProfMiddlename.TATitle, taProfMiddlename.TAName, taProfMiddlename.TAInputType, "", nil, nil, ""),
		schema.NewInput(taProfPhoneNumber.TATitle, taProfPhoneNumber.TAName, taProfPhoneNumber.TAInputType, "", nil, nil, ""),
		schema.NewInput(taProfEmail.TATitle, taProfEmail.TAName, taProfEmail.TAInputType, "", nil, nil, ""),
	}
}

func getInputsFromProfDB(profDB *profschema.ProfDB) []*schema.Input {
	return []*schema.Input{
		schema.NewInput("", taProfLastname.TAName, taProfLastname.TAInputType, profDB.ProfLastname, nil, nil, ""),
		schema.NewInput("", taProfFirstname.TAName, taProfFirstname.TAInputType, profDB.ProfFirstname, nil, nil, ""),
		schema.NewInput("", taProfMiddlename.TAName, taProfMiddlename.TAInputType, profDB.ProfMiddlename, nil, nil, ""),
		schema.NewInput("", taProfPhoneNumber.TAName, taProfPhoneNumber.TAInputType, profDB.ProfPhoneNumber, nil, nil, ""),
		schema.NewInput("", taProfEmail.TAName, taProfEmail.TAInputType, profDB.ProfEmail, nil, nil, ""),
		schema.NewInput("", taProfIsHidden.TAName, taProfIsHidden.TAInputType, profDB.ProfIsHidden, nil, nil, ""),
	}
}

func ProfTableRowEdit(profDB *profschema.ProfDB) templ.Component {
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
		templ_7745c5c3_Err = view.TableRowEdit(getInputsFromProfDB(profDB), fmt.Sprintf(bodyVals, profDB.ProfID), updatePOSTURL, deletePOSTURL).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ProfTableRow(profDB *profschema.ProfDB) templ.Component {
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
		templ_7745c5c3_Err = view.TableRow(getInputsFromProfDB(profDB), fmt.Sprintf(bodyVals, profDB.ProfID), editPOSTURL, deletePOSTURL).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ProfTableRows(profsDB []*profschema.ProfDB) templ.Component {
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
		for _, profDB := range profsDB {
			templ_7745c5c3_Err = ProfTableRow(profDB).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

func Prof() templ.Component {
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
		templ_7745c5c3_Err = view.Table(tableTitle, getPOSTURL, getTableHeaders()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ProfPage() templ.Component {
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
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(getprofPOSTURL)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/profview/profview.templ`, Line: 87, Col: 27}
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
