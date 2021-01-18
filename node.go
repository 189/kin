package kin

import "strings"

type node struct {
	pattern string
	wild bool
	part string
	children []*node
}

// 从前缀树中，查找最末级，匹配 URL.Path 的 node
func (n *node) search(parts []string, height int) *node{
	// 最后一层，或者当前层 part 为 *file 直接返回
	if len(parts) == height || strings.HasPrefix(parts[height], "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	// 收集当前height 层级中 children 里符合 part 的 node
	nodes := n.matchChildren(parts[height])

	// 继续从 下级 children 中查找，最后或得 最末级 匹配 part 的 node
	for _, node := range nodes {
		result := node.search(parts, height + 1)
		if result != nil {
			return result
		}
	}
	return nil
}

// 找当前层级最近的一个匹配 part 的节点， 一旦找到直接返回
func (n *node) matchChild (part string) *node {
	for _, child := range n.children {
		// 找到一个，直接返回
		if child.part == part || child.wild {
			return child
		}
	}
	return nil
}

// 找当前层级的所有匹配 part 的节点, 收集当前层级所有匹配 part 的 *node 不会再深入下层遍历
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.wild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}


// 生成前缀树
func (n *node) insert (pattern string, parts []string, height int) {
	// 最后一层, 设置完 pattern 后直接返回
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)

	// 生成当前节点
	if child == nil {
		child = &node{ part: part, wild: part[0] == ':' || part[0] == '*'}
		child.children = append(child.children, child)
	}
	// 递归 继续生成下一层节点，形成一个分叉
	child.insert(pattern, parts, height + 1)
}

// 收集所有 pattern 不为空的 node
func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}


