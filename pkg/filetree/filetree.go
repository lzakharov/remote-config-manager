package filetree

import (
	"strings"
)

const sep = "/"

// Tree represents a file tree.
type Tree []Node

// Node is a file tree node.
type Node struct {
	Value string
	Name  string
	Tree  Tree
}

// IsLeaf reports whether node is a leaf.
func (n Node) IsLeaf() bool {
	return len(n.Tree) == 0
}

// Build builds a new file tree from file names.
func Build(names []string) Tree {
	result := make(Tree, 0)

	for _, name := range names {
		result = add(result, name, parse(name))
	}

	return result
}

func parse(name string) []string {
	var path []string

	if strings.HasPrefix(name, sep) {
		path = append(path, sep)
		name = name[1:]
	}
	path = append(path, strings.Split(name, sep)...)

	return path
}

func add(t Tree, name string, path []string) Tree {
	if len(path) == 0 {
		return t
	}

	var i int

	for i = 0; i < len(t); i++ {
		if t[i].Value == path[0] {
			break
		}
	}

	if i == len(t) {
		t = append(t, Node{Value: path[0], Name: name})
	}

	t[i].Tree = add(t[i].Tree, name, path[1:])

	return t
}
