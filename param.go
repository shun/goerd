package main

import (
	"fmt"
)

// ########################################################
//
// strunt
//

// Param aaa
type Param struct {
	key string
	val string
}

// ParamList aaa
type ParamList []Param

func (lst *ParamList) add(item Param) {
	*lst = append(*lst, item)
}

func (lst ParamList) toString() string {
	ret := ""
	sz := len(lst)
	for idx, item := range lst {
		ret += item.toString()
		if idx+1 != sz {
			ret += ", "
		} else {
			ret += " "
		}
	}

	return ret
}

func (prm Param) toString() string {
	ret := ""
	if isTextParam(prm.key) {
		ret += fmt.Sprintf("%s = \"%s\"", prm.key, prm.val)
	} else {
		ret += fmt.Sprintf("%s = %s", prm.key, prm.val)
	}

	return ret
}
