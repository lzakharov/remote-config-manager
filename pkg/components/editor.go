package components

import (
	"fmt"

	"github.com/lzakharov/remote-config-manager/pkg/components/metro"
	"github.com/lzakharov/remote-config-manager/pkg/service"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Editor is a config editor.
type Editor struct {
	app.Compo

	key string
}

func (e *Editor) Render() app.UI {
	return metro.ContainerFluid(
		metro.Row(
			metro.Cell12(
				app.Div().ID("editor").
					Style("height", "600px"),
			),
		),
		metro.Row(
			metro.Cell(
				metro.PrependedTextInput("Name:").
					Attr("data-cls-input", "text-bold fg-black").
					ReadOnly(true).
					Value(e.key),
			),
			metro.Cell(
				metro.ContainerFluid(
					metro.Button().
						Disabled(e.key == "").
						Text("Close").
						OnClick(e.onClose),
					metro.Button().
						Disabled(e.key == "").
						Text("Refresh").
						OnClick(e.onRefresh),
					metro.SuccessButton().
						Disabled(e.key == "").
						Text("Save").
						OnClick(e.onSave),
				),
			),
		),
	)
}

func (e *Editor) OnMount(ctx app.Context) {
	app.Window().Set("editor",
		app.Window().Get("monaco").Get("editor").
			Call("create",
				app.Window().GetElementByID("editor"),
				map[string]interface{}{
					"language": "yaml",
					"minimap": map[string]interface{}{
						"enabled": false,
					},
				}))

	ctx.GetState(editorKeyState, &e.key)
	e.refresh()
	e.Update()

	ctx.ObserveState(editorKeyState).
		OnChange(func() {
			e.refresh()
			e.Update()
		}).
		Value(&e.key)
}

func (e *Editor) OnResize(_ app.Context) {
	app.Window().Get("editor").Call("layout")
	e.Update()
}

func (e *Editor) onRefresh(_ app.Context, _ app.Event) {
	e.refresh()
}

func (e *Editor) refresh() {
	if e.key == "" {
		app.Window().Get("editor").Call("setValue", "")
		return
	}

	value, err := service.Get(e.key)
	if err != nil {
		handleErr(err)
	}

	app.Window().Get("editor").Call("setValue", value)
}

func (e *Editor) onSave(ctx app.Context, _ app.Event) {
	value := app.Window().Get("editor").Call("getValue").String()

	ctx.Async(func() {
		err := service.Put(e.key, value)
		if err != nil {
			handleErr(err)
		}

		metro.Notify(fmt.Sprintf("%q saved", e.key), "success", false)
	})
}

func (e *Editor) onClose(ctx app.Context, _ app.Event) {
	ctx.SetState(editorKeyState, "", app.Persist)
}
