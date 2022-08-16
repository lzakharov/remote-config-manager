package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/lzakharov/remote-config-manager/pkg/components/metro"
)

// App is an application component.
type App struct {
	app.Compo
}

func (a *App) Render() app.UI {
	return metro.ContainerFluid(
		metro.Grid(
			metro.Row(
				metro.Cell3(
					metro.ContainerFluid(
						&Tree{},
					),
				),
				metro.Cell9(
					&Editor{},
				),
			),
		),
	)
}
