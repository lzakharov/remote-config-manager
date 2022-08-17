package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/klauspost/compress/gzhttp"
	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/lzakharov/remote-config-manager/pkg/components"
)

func main() {
	addr := flag.String("addr", ":8080", `Server address.`)
	flag.Parse()

	app.Route("/", &components.Main{})
	app.RunWhenOnBrowser()

	handler := &app.Handler{
		Name:        "remote-config-manager",
		Description: "Remote Configuration Manager UI",
		Author:      "Lev Zakharov",
		Styles: []string{
			"/web/node_modules/metro4/build/css/metro-all.min.css",
			"/web/node_modules/monaco-editor/min/vs/editor/editor.main.css",
		},
		Scripts: []string{
			"/web/node_modules/metro4/build/js/metro.min.js",
		},
		RawHeaders: []string{`
			<script>
				var require = { paths: { vs: '/web/node_modules/monaco-editor/min/vs' } };
			</script>
			<script src="/web/node_modules/monaco-editor/min/vs/loader.js"></script>
			<script src="/web/node_modules/monaco-editor/min/vs/editor/editor.main.nls.js"></script>
			<script src="/web/node_modules/monaco-editor/min/vs/editor/editor.main.js"></script>
		`,
		},
	}

	err := http.ListenAndServe(*addr, gzhttp.GzipHandler(handler))
	if err != nil {
		log.Fatal(err)
	}
}
