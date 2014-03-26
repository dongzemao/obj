package main
import (
	"fmt"
)

type aaabbb struct {
	aaa int
	bbb string
}
type aabb struct {
	aaa int
}

func main() {
	// str := "1好1"
	// bytes := []byte(str)
	// a := fmt.Println
	// test()
	// a(a, bytes)
	aaa := aaabbb{1, "str1"}
	aa := aaa.(abc)
	fmt.Println(aaa, aa)
}

func (aaabbb) ooo() {
	fmt.Println("aaa, aa")
}

func test(nnn ...int) {
	fmt.Println(nnn)
}

type abc interface {
	ooo()
}
