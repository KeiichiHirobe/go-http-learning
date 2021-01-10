package main

import "fmt"

func main() {
	// len=0 cap=4
	s := make([]int, 0, 4)
	s = append(s, 1, 2, 3)

	smap := map[string][]int{
		"a": s,
	}
	// 	1,2,3
	fmt.Println(smap["a"])

	s = append(s, 4)

	// 	1,2,3
	fmt.Println(smap["a"])

	// We must write explicitly
	smap["a"] = s
	// 	1,2,3,4
	fmt.Println(smap["a"])

	s[0] = 0
	// 0,2,3,4
	// We write implicitly
	fmt.Println(smap["a"])

	s = s[:1]
	// 0,2,3,4
	fmt.Println(smap["a"])

	// We must write explicitly
	smap["a"] = s
	// 0
	fmt.Println(smap["a"])
}
