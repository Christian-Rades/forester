package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type node struct {
	name     string
	label    string
	children map[string]*node
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	root := &node{
		name:     "/",
		children: map[string]*node{},
	}

	for scn.Scan() {
		line := scn.Text()
		parts := strings.Split(line, "/")
		parts = filterEmpty(parts)
		root.addPath(parts)
	}

	if scn.Err() != nil {
		panic(scn.Err())
	}

	nm := root.buildNodeMap(map[string][]*node{})

	resolveCollisons(nm)

	fmt.Println("graph tree {")
	fmt.Println("rankdir=LR")
	root.toDot(os.Stdout)
	fmt.Println("}")
}

func filterEmpty(parts []string) []string {
	filtered := make([]string, 0, len(parts))
	for _, p := range parts {
		if len(p) > 0 && p != " " {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func (n *node) addPath(parts []string) {
	if len(parts) == 0 {
		return
	}
	top := parts[0]
	rest := parts[1:]

	top = strings.ReplaceAll(top, "%s", ":string:")

	if child, ok := n.children[top]; ok {
		child.addPath(rest)
	} else {
		child := &node{
			name:     top,
			children: map[string]*node{},
		}
		n.children[top] = child
		child.addPath(rest)
	}
}

func (n *node) toDot(w io.Writer) {
	if len(n.label) > 0 {
		_, err := fmt.Fprintf(w, "%q [label=%q]\n", n.name, n.label)
		if err != nil {
			panic(err)
		}
	}

	for _, c := range n.children {
		_, err := fmt.Fprintf(w, "%q -- %q\n", n.name, c.name)
		if err != nil {
			panic(err)
		}
		c.toDot(w)
	}
}

func (n *node) buildNodeMap(nodeMap map[string][]*node) map[string][]*node {
	if nodeList, ok := nodeMap[n.name]; ok {
		nodeList = append(nodeList, n)
		nodeMap[n.name] = nodeList
	} else {
		nodeMap[n.name] = []*node{n}
	}
	for _, c := range n.children {
		nodeMap = c.buildNodeMap(nodeMap)
	}
	return nodeMap
}

func resolveCollisons(nodeMap map[string][]*node) {
	for _, nodeList := range nodeMap {
		if len(nodeList) > 1 {
			for i, node := range nodeList {
				node.label = node.name
				node.name = fmt.Sprintf("%s_%d", node.name, i)
			}
		}
	}
}
