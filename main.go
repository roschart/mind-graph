package main

import (
	"fmt"
	"strings"
)

type trigger struct {
	name string
	node *node
}

type node struct {
	name     string
	triggers []trigger
}

func (n node) String() string {
	builder := strings.Builder{}
	builder.WriteString(n.name)

	for _, t := range n.triggers {
		builder.WriteString(fmt.Sprintf("\n-- %s --\n", t.name))
		builder.WriteString(t.node.name)

	}
	return builder.String()
}

func main() {
	nodes := create_nodes()
	for _, n := range nodes {
		fmt.Println(n)
		fmt.Println("##############")
	}
}

func create_nodes() []node {
	españa := node{
		name: "España",
	}

	galicia := node{
		name: "Galicia",
	}

	galicia.triggers = make([]trigger, 1)
	galicia.triggers[0] = trigger{name: "Páis", node: &españa}

	return []node{galicia, españa}

}
