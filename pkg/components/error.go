package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func handleErr(err error) {
	app.Window().Get("Metro").Get("notify").
		Call("create", err.Error(), nil, map[string]interface{}{
			"cls":      "alert",
			"keepOpen": true,
		})
}
