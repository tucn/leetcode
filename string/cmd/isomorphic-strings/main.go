package main

import "fmt"

func main() {
	str1 := "egg"
	str2 := "add"
	fmt.Printf("%v\n", isIsomorphic(str1, str2))
	str1 = "foo"
	str2 = "bar"
	fmt.Printf("%v\n", isIsomorphic(str1, str2))
	str1 = "paper"
	str2 = "title"
	fmt.Printf("%v\n", isIsomorphic(str1, str2))
	str1 = "badc"
	str2 = "baba"
	fmt.Printf("%v\n", isIsomorphic(str1, str2))

}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chkMap1 := make(map[string]string)
	chkMap2 := make(map[string]string)
	for k, _ := range s {
		sVal := string(s[k])
		tVal := string(t[k])
		if chkVal, ok := chkMap1[sVal]; ok {
			if chkVal != tVal {
				return false
			}
			continue
		}
		if chkVal2, ok := chkMap2[tVal]; ok {
			if chkVal2 != sVal {
				return false
			}
			continue
		}
		chkMap1[sVal] = tVal
		chkMap2[tVal] = sVal
	}
	return true

}
