package _0_isNumber

import (
	"strconv"
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	var strList []string
	strList = append(strList, s)
	if find := strings.Contains(s, "e"); find {
		strList = strings.Split(s, "e")
	} else if find := strings.Contains(s, "e"); find {
		strList = strings.Split(s, "E")
	}
	if len(strList) > 2 {
		return false
	} else if len(strList) == 2 {
		if _, err := strconv.Atoi(strList[1]); err != nil {
			return false
		}
	}
	if _, err := strconv.ParseFloat(strList[0], 32); err != nil {
		return false
	}
	return true
}
