package diff

import (
	"reflect"
	"sort"
)

type NodeType string

const (
	Added     NodeType = "added"
	Removed   NodeType = "removed"
	Changed   NodeType = "changed"
	Unchanged NodeType = "unchanged"
	Nested    NodeType = "nested"
)

type Node struct {
	Key      string      `json:"key"`
	Type     NodeType    `json:"type"`
	Value    interface{} `json:"value,omitempty"`
	OldValue interface{} `json:"old_value,omitempty"`
	Children []Node      `json:"children,omitempty"`
}

func BuildDiff(a, b map[string]interface{}) []Node {
	keysMap := make(map[string]struct{})

	for k := range a {
		keysMap[k] = struct{}{}
	}
	for k := range b {
		keysMap[k] = struct{}{}
	}

	keys := make([]string, 0, len(keysMap))
	for k := range keysMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var nodes []Node
	for _, k := range keys {
		va, oka := a[k]
		vb, okb := b[k]

		switch {
		case oka && !okb:
			nodes = append(nodes, Node{
				Key:   k,
				Type:  Removed,
				Value: va,
			})
		case !oka && okb:
			nodes = append(nodes, Node{
				Key:   k,
				Type:  Added,
				Value: vb,
			})
		default:
			ma, isa := va.(map[string]interface{})
			mb, isb := vb.(map[string]interface{})

			if isa && isb {
				children := BuildDiff(ma, mb)
				nodes = append(nodes, Node{
					Key:      k,
					Type:     Nested,
					Children: children,
				})
				continue
			}

			if reflect.DeepEqual(va, vb) {
				nodes = append(nodes, Node{
					Key:   k,
					Type:  Unchanged,
					Value: va,
				})
			} else {
				nodes = append(nodes, Node{
					Key:      k,
					Type:     Changed,
					OldValue: va,
					Value:    vb,
				})
			}
		}
	}
	return nodes
}
