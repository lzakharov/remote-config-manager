package metro

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// ContainerFluid is a fluid container.
func ContainerFluid(elems ...app.UI) app.HTMLDiv {
	return app.Div().Class("container-fluid").Body(elems...)
}
