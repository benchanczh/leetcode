package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	result := [][]string{}

	for _, str := range strs {
		stringList := strings.Split(str, "")
		sort.Strings(stringList)
		key := strings.Join(stringList, "")

		if _, ok := m[key]; ok {
			m[key] = append(m[key], str)
		} else {
			m[key] = []string{str}
		}
	}

	for _, strList := range m {
		result = append(result, strList)
	}
	return result
}
