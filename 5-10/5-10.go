// 5-10
package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b

}

func diff(a int, b int) int {
	return a - b
}

func main() {
	f := []func(int, int) int{sum, diff} //슬라이스를 생성 후 함수로 초기화
	fmt.Println(f[0](1, 2))              //배열의 첫 번째 요소로 함수 호출
	fmt.Println(f[1](1, 2))              //배열의 두 번째 요소로 함수 호출
}
