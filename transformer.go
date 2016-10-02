package main
/*
import (
	"os"
)

type visitor map[string]func(n *node, p node)

func traverser(a tree, v visitor) {
	traverseNode(node(a), node{}, v)
}

func traverseArray(a []node, p node, v visitor) {
	for _, child := range a {
		traverseNode(child, p, v)
	}
}

func traverseNode(n, p node, v visitor) {
	for k, va := range v {
		if k == n.kind {
			va(&n, p)
		}
	}
	switch n.kind {
	case "program":
		traverseArray(n.body, n, v)
		break
	case "expression":
		traverseArray(n.params, n, v)
		break
	case "factor":
		break
	default :
//		fmt.Print(n.kind, ":(", n.x, ",", n.y, ")\n")
		os.Exit(1)
	}
}

func transformer(a tree) tree {
	ntree := tree{
		kind: "program",
		body: []node{},
	}
	a.context = &ntree.body
	traverser(a, map[string]func(n *node, p node){
		"factor": func(n *node, p node) {
			*p.context = append(*p.context, node{
				kind:  "factor",
				value: n.value,
			})
		},
		"expression": func(n *node, p node) {
			e := node{
				kind: "expression",
				callee: &node{
					kind: "identificador",
					name: n.name,
				},
				arguments: new([]node),
			}
			n.context = e.arguments
			if p.kind != "expression" {
				es := node{
					kind:       "expstat",
					expression: &e,
				}
				*p.context = append(*p.context, es)
			} else {
				*p.context = append(*p.context, e)
			}
		},
	})
	return ntree
}
*/