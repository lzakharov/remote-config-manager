package metro

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Button is a default button.
func Button() app.HTMLButton {
	return app.Button().Class("button")
}

// SuccessButton is a success button.
func SuccessButton() app.HTMLButton {
	return app.Button().Class("button", "success")
}
