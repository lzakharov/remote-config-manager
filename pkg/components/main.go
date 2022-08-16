package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Main is a root application component.
type Main struct {
	app.Compo
}

func (m *Main) Render() app.UI {
	return app.Div().Body(
		&Navbar{},
		&App{},
		&Footer{},
	)
}
