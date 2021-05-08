package utils

import "strings"

func ListEqual(a []uint64, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, ai := range a {
		if ai != b[idx] {
			return false
		}
	}

	return true
}

func StringListEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, ai := range a {
		if strings.Compare(ai, b[idx]) != 0 {
			return false
		}
	}

	return true
}
