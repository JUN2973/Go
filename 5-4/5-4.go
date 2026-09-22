// 5-4
package main

import (
	"fmt"
)

func SumAndDIff(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	sum, diff := SumAndDIff(6, 2)
	fmt.Println(sum, diff)
}
