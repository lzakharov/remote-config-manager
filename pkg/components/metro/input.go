package metro

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// TextInput is a text input.
func TextInput() app.HTMLInput {
	return app.Input().Type("text").Attr("data-role", "input")
}
