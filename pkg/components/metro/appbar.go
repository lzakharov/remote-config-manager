package metro

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// AppBar is an app bar.
func AppBar(elems ...app.UI) app.HTMLHeader {
	return app.Header().
		Class("container-fluid").
		Attr("data-role", "appbar").
		Attr("data-expand-point", "fs").
		Body(elems...)
}

// AppBarMenu is an app bar menu.
func AppBarMenu(elems ...app.UI) app.HTMLHeader {
	return app.Header().
		Class("app-bar-menu").
		Body(elems...)
}

// Brand is a brand component.
func Brand(name string) app.HTMLA {
	return app.A().Href("#").Class("brand", "no-hover").Text(name)
}
