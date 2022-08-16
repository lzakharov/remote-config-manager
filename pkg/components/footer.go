package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Footer is an application footer.
type Footer struct {
	app.Compo
}

func (*Footer) Render() app.UI {
	return app.Footer().Class("p-2").Body(
		app.P().Body(
			app.Text("Created with "),
			app.Span().Class("fg-red").Text("♥"),
			app.Text(" by "),
			app.A().
				Href("https://lzakharov.github.io").
				Text("Lev Zakharov"),
			app.Text(". Created with "),
			app.A().
				Href("https://github.com/maxence-charriere/go-app").
				Text("go-app"),
			app.Text(". Licensed under the "),
			app.A().
				Href("https://github.com/lzakharov/remote-config-manager/blob/main/LICENSE").
				Text("MIT License"),
			app.Text("."),
		),
	)
}
