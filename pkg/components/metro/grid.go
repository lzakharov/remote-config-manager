package metro

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Grid is a grid layout container.
func Grid(rows ...app.UI) app.HTMLDiv {
	return app.Div().Class("grid").Body(rows...)
}

// Row is a grid layout row.
func Row(cells ...app.UI) app.HTMLDiv {
	return app.Div().Class("row").Body(cells...)
}

// Cell is a grid layout cell.
func Cell(elems ...app.UI) app.HTMLDiv {
	return app.Div().Class("cell").Body(elems...)
}

// Cell3 is a grid layout cell of three columns.
func Cell3(elems ...app.UI) app.HTMLDiv {
	return app.Div().Class("cell-3").Body(elems...)
}

// Cell9 is a grid layout cell of nine columns.
func Cell9(elems ...app.UI) app.HTMLDiv {
	return app.Div().Class("cell-9").Body(elems...)
}

// Cell12 is a grid layout cell of twelve columns.
func Cell12(elems ...app.UI) app.HTMLDiv {
	return app.Div().Class("cell-12").Body(elems...)
}
