package cubery

import (
	"fmt"
	"strings"
)

type node struct {
	route          string  // route to be matched, example: /p/:lang
	part           string  // part of a route, example: :lang
	children       []*node // child nodes
	isPreciseMatch bool    // is precise matching
}

func (n *node) String() string {
	return fmt.Sprintf("node{route=%s, part=%s, isPreciseMatch=%t}", n.route, n.part, n.isPreciseMatch)
}

func (n *node) insert(route string, parts []string, height int) {
	if len(parts) == height {
		n.route = route
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isPreciseMatch: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(route, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.route == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) travel(list *[]*node) {
	if n.route != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

// returns the first matched node
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isPreciseMatch {
			return child
		}
	}
	return nil
}

// returns all the matched node
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isPreciseMatch {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
