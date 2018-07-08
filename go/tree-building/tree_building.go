// Package tree contains functions for building a tree from a list of abstract
// records.
package tree

import (
	"errors"
	"sort"
)

// Record struct with ID and Parent
type Record struct {
	ID, Parent int
}

// Node struct with ID and Childrens
type Node struct {
	ID       int
	Children []*Node
}

// Build constructs a tree from the given list of records and returns a pointer
// to its root.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, errors.New("Illegal root node")
	}

	root := &Node{}
	nodes := make(map[int]*Node)
	nodes[0] = root

	for i := 1; i < len(records); i++ {
		r := records[i]

		if r.ID < r.Parent {
			return nil, errors.New("Record with ID lower than Parent")
		}

		if r.ID == r.Parent {
			return nil, errors.New("Non-root record with self Parent")
		}

		if r.ID < 1 || r.ID >= len(records) {
			return nil, errors.New("Illegal record ID")
		}

		if _, exists := nodes[r.ID]; exists {
			return nil, errors.New("Duplicate record")
		}

		nn := &Node{ID: r.ID}
		nodes[r.ID] = nn

		if parent, exists := nodes[r.Parent]; exists {
			parent.Children = append(parent.Children, nn)
		}
	}

	return root, nil
}
