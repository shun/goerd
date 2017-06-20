package main

import (
	"fmt"
)

// ########################################################
//
// strunt
//

// Relation aaa
type Relation struct {
	t1    string
	t2    string
	key1  string
	key2  string
	label string
}

// Relations aaa
type Relations []*Relation

func (lst Relations) toString() string {
	ret := ""
	sz := len(lst)
	for idx, item := range lst {
		p1 := tables.getPort(item.t1, item.key1)
		p2 := tables.getPort(item.t2, item.key2)
		ret += fmt.Sprintf("%s:r%d -> %s:r%d [ label = \"[ %s ]\", arrowhead = none, arrowtail = none, weight = 5 ];", item.t1, p1, item.t2, p2, item.label)

		if sz > idx {
			ret += "\n"
		}
	}

	return ret
}
