package main

import (
	"fmt"
)

// ########################################################
//
// strunt
//

const (
	unknown = iota
	pk
	fk
)

// Column aaa
type Column struct {
	name    string
	keytype int
}

// Table aaa
type Table struct {
	name string
	cols []*Column
}

func (t *Table) toString() string {
	ret := t.name + " [ label = <\n"
	ret += "<table border=\"0\" cellborder=\"1\" cellspacing=\"0\">\n"

	// add table header
	ret += fmt.Sprintf("<tr><td align=\"center\" port=\"tn\" bgcolor=\"yellow\">%s</td></tr>\n", t.name)

	// add rows
	for rowno, col := range t.cols {
		ret += fmt.Sprintf("<tr><td align=\"left\" port=\"r%d\">%s</td></tr>\n", rowno+1, col.name)

	}
	ret += "</table>> ];"
	return ret
}

// Tables aaa
type Tables []*Table

func (lst Tables) toString() string {
	ret := ""
	for _, t := range lst {
		ret += t.toString() + "\n"
	}
	return ret
}

func (lst Tables) getPort(tblname string, colname string) int {
	ret := -1
	for _, t := range lst {
		if tblname != t.name {
			continue
		}
		ret = t.getPort(colname)
		if ret > 0 {
			return ret
		}
	}
	return ret
}

func (t *Table) getPort(colname string) int {
	for idx, col := range t.cols {
		if col.name == colname {
			return idx + 1
		}
	}

	return -1
}
