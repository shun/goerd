package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ########################################################
//
// global values
//

var reTable = regexp.MustCompile(`^t:([^ ]+) \[`)
var reColumn = regexp.MustCompile(`^.?:([^ ]+)`)
var reRelation = regexp.MustCompile(`([^:]+):([^:]+) (.+--.+) ([^:]+):([^:]+)`)
var tables Tables
var relations Relations

var defaultGraph ParamList
var defaultNode ParamList

func isTextParam(prm string) bool {
	textParams := []string{"charset", "bgcolor", "fontname"}
	for _, tp := range textParams {
		if tp == prm {
			return true
		}
	}

	return false
}

// ########################################################
//
// functions
//

func setDefaultParam() {
	defaultGraph.add(Param{key: "charset", val: "UTF-8"})
	defaultGraph.add(Param{key: "bgcolor", val: "#EDEDED"})
	defaultGraph.add(Param{key: "rankdir", val: "LR"})

	defaultNode.add(Param{key: "shape", val: "plaintext"})
	defaultNode.add(Param{key: "fontname", val: "Migu 1M"})
	defaultNode.add(Param{key: "fontsize", val: "12"})

}

func parse(lines []string) {

	var ret []string
	var table *Table
	var relation *Relation
	for i := 0; i < len(lines); i++ {
		if lines[i][0] == ']' {
			// found end of the table definition
			tables = append(tables, table)
			continue
		}

		ret = reTable.FindStringSubmatch(lines[i])
		if len(ret) > 0 {
			// found table definition
			table = new(Table)
			table.name = ret[1]
			continue
		}

		ret = reRelation.FindStringSubmatch(lines[i])
		if len(ret) > 0 {
			// found relation definition
			relation = new(Relation)
			relation.t1 = ret[1]
			relation.key1 = ret[2]
			relation.t2 = ret[4]
			relation.key2 = ret[5]
			relation.label = ret[3]
			relations = append(relations, relation)
			continue
		}

		ret = reColumn.FindStringSubmatch(lines[i])
		if len(ret) > 0 {
			// found column definition
			items := strings.Split(lines[i], ":")
			key := unknown
			if items[0] == "p" {
				key = pk
			} else if items[0] == "f" {
				key = fk
			}
			col := new(Column)
			col.keytype = key
			col.name = items[1]
			table.cols = append(table.cols, col)

			continue
		}
	}
}

func readfile(fin *os.File) []string {
	lines := []string{}
	sc := bufio.NewScanner(fin)

	reComment := regexp.MustCompile(`#.*$`)
	for sc.Scan() {
		line := reComment.ReplaceAllString(strings.TrimSpace(sc.Text()), "")
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	return lines
}

func execute() {
	output := "digraph G {\n"

	output += "graph [ "
	output += defaultGraph.toString()
	output += "];\n"

	output += "node [ "
	output += defaultNode.toString()
	output += "];\n"

	output += tables.toString()
	output += relations.toString()

	output += "}\n"

	fmt.Println(output)
}

func main() {
	setDefaultParam()
	fin, err := os.Open(os.Args[1])
	if err != nil {
		panic("file err")
	}
	lines := readfile(fin)

	parse(lines)
	execute()

	fin.Close()
}
