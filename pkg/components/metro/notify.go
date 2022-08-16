package metro

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// Notify creates a new notification.
func Notify(msg, cls string, keepOpen bool) {
	app.Window().Get("Metro").Get("notify").
		Call("create", msg, nil, map[string]interface{}{
			"cls":      cls,
			"keepOpen": keepOpen,
		})
}
