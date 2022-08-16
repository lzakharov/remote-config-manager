package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/lzakharov/remote-config-manager/pkg/components/metro"
)

// Navbar is a navigation menu.
type Navbar struct {
	app.Compo
}

func (*Navbar) Render() app.UI {
	return metro.AppBar(
		metro.AppBarMenu(
			metro.Brand("Remote Config Manager"),
			app.Li().Body(
				app.A().
					Href("https://github.com/lzakharov/remote-config-manager/blob/main/docs/README.md").
					Text("Docs"),
			),
			app.Li().Body(
				app.A().
					Href("https://github.com/lzakharov/remote-config-manager").
					Text("GitHub"),
			),
		),
	)
}
