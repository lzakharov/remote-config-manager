package metro

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// TreeView is a tree view.
func TreeView(nodes ...app.UI) app.HTMLUl {
	return app.Ul().Attr("data-role", "treeview").Body(nodes...)
}

// TreeViewNode is a tree view node.
func TreeViewNode(text string) app.HTMLLi {
	return app.Li().
		Attr("data-caption", text)
}
