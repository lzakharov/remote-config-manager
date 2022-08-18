package components

import (
	"sort"

	"github.com/lzakharov/remote-config-manager/pkg/filetree"
	"github.com/lzakharov/remote-config-manager/pkg/transport"
	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/lzakharov/remote-config-manager/pkg/components/metro"
)

// Tree is a navigation tree.
type Tree struct {
	app.Compo
}

func (t *Tree) Render() app.UI {
	keys, err := transport.ListKeys()
	if err != nil {
		handleErr(err)
	}

	sort.Strings(keys)

	return metro.ContainerFluid(
		metro.Row(
			metro.Cell(
				app.P().Class("text-bold").Body(
					app.Text("Keys"),
				),
			),
		),
		app.Hr(),
		metro.Row(
			metro.TreeView(
				t.build(filetree.Build(keys))...,
			),
		),
	)
}

func (t *Tree) build(tree filetree.Tree) []app.UI {
	result := make([]app.UI, len(tree))

	for _, node := range tree {
		if node.IsLeaf() {
			result = append(result, t.leaf(node.Value, node.Name))
		} else {
			result = append(result, t.nonLeaf(node.Value, t.build(node.Tree)...))
		}
	}

	return result
}

func (*Tree) nonLeaf(text string, elems ...app.UI) app.HTMLLi {
	return metro.TreeViewNode(text).Body(
		app.Ul().Body(elems...),
	)
}

func (*Tree) leaf(text, key string) app.HTMLLi {
	return metro.TreeViewNode(text).
		OnClick(func(ctx app.Context, _ app.Event) {
			ctx.SetState(editorKeyState, key, app.Persist)
		})
}
