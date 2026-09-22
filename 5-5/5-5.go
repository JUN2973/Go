// 5-5
package main

import (
	"fmt"
)

func SumAndDIff(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main() {
	sum, diff := SumAndDIff(6, 2)
	fmt.Println(sum, diff)
}