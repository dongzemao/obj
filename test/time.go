package main
import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Unix()

	fmt.Println(t)
	fmt.Println(time.Unix(t, 0))
	fmt.Println(time.Unix(t, 0))
	fmt.Println(time.Unix(t, 0))
	fmt.Println(t)
}
