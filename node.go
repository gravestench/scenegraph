package scenegraph

import (
	"github.com/gravestench/mathlib"
)

// NewNode creates and initializes a new scene graph node
func NewNode() *Node {
	n := &Node{
		World:    mathlib.NewMatrix4(nil),
		Local:    mathlib.NewMatrix4(nil),
		children: make([]*Node, 0),
	}

	return n
}

// Node is a scene graph node.
type Node struct {
	parent   *Node
	World    *mathlib.Matrix4
	Local    *mathlib.Matrix4
	children []*Node
}

// SetParent sets the parent of this scene graph node
func (n *Node) SetParent(p *Node) *Node {
	if n.parent != nil {
		n.parent.removeChild(n)
	}

	n.parent = p

	if p != nil {
		n.parent.children = append(n.parent.children, n)
	}

	return n
}

func (n *Node) removeChild(m *Node) *Node {
	if m == nil {
		return n
	}

	for idx := len(n.children) - 1; idx >= 0; idx-- {
		if n.children[idx] != m {
			continue
		}

		n.children = append(n.children[:idx], n.children[idx+1:]...)
	}

	return n
}

// UpdateWorldMatrix updates this node's World matrix using the (optional) parent World matrix
func (n *Node) UpdateWorldMatrix(args ...*mathlib.Matrix4) *Node {
	// this is a hack so that we can just call `node.UpdateWorldMatrix()`
	parentWorldMatrix := (*mathlib.Matrix4)(nil)
	if len(args) > 0 {
		parentWorldMatrix = args[0]
	}

	n.World = parentWorldMatrix

	for idx := range n.children {
		n.children[idx].UpdateWorldMatrix(n.GetWorldMatrix())
	}

	return n
}

// GetWorldMatrix applies the local transform to the world matrix and returns it
func (n *Node) GetWorldMatrix() *mathlib.Matrix4 {
	if n.World == nil {
		return n.Local.Clone()
	}

	return n.World.Clone().Multiply(n.Local)
}
