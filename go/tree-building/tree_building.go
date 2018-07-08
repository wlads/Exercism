package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type ByID []Record

func (r ByID) Len() int           { return len(r) }
func (r ByID) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByID) Less(i, j int) bool { return r[i].ID < r[j].ID }

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Sort(ByID(records))

	var root *Node
	nodes := make(map[int]*Node)

	for i, r := range records {
		node := &Node{r.ID, nil}

		if nodes[r.ID] != nil {
			return nil, errors.New("record set contains duplicate node")
		}

		nodes[r.ID] = node

		if r.ID == r.Parent {
			if root != nil {
				return nil, errors.New("record set contains more than one root node")
			}

			root = node
			continue
		}

		if i != r.ID {
			return nil, errors.New("node has a wrong ID")
		}

		if r.Parent > node.ID {
			return nil, errors.New("parent ID is greater than nodes ID")
		}

		parent, ok := nodes[r.Parent]

		if !ok {
			return nil, errors.New("parent does not exist")
		}

		if parent.Children == nil {
			parent.Children = make([]*Node, 0)
		}

		parent.Children = append(parent.Children, node)
	}

	if root == nil {
		return nil, errors.New("record set contains no root node")
	}

	return root, nil
}
