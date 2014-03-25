package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	str = "aaabbbababab"
	// num1 := strings.Count(str, "abaa")
	// fmt.Println(num1)
	// num2 := strings.Contains(str, "ababb")
	// fmt.Println(num2)
	// num3 := strings.Index(str, "bbba")
	// fmt.Println(num3)
	// ss := strings.SplitN(str, "", 2)
	// fmt.Println(ss)
	ss := strings.SplitAfterN(str, "abbba", 2)
	fmt.Println(ss)
}
