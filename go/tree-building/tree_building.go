package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// Build reconstructs the tree structure from the given records.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node, len(records))

	for i, record := range records {
		if record.ID != i {
			return nil, errors.New("records are out of order or have missing IDs")
		}
		if record.ID == 0 && record.Parent != 0 {
			return nil, errors.New("root node has invalid parent")
		}

		if record.ID != 0 && record.Parent >= record.ID {
			return nil, errors.New("child node has invalid parent")
		}

		nodes[record.ID] = &Node{ID: record.ID}
		if record.ID != 0 { 
			parentNode := nodes[record.Parent]
			parentNode.Children = append(parentNode.Children, nodes[record.ID])
		}
	}

	return nodes[0], nil
}
