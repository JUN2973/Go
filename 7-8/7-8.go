// 7-8
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 3) //버퍼가 3개인 비동기 채널 생성
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true             //채널에 true를 보냄, 버퍼가 가득 차면 대기
			fmt.Println("고루틴 : ", i) //반복문의 변수 출력
		}
	}()

	for i := 0; i < count; i++ {
		<-done                    //버퍼에 값이 없으면 대기, 값을 꺼냄
		fmt.Println("메인함수 : ", i) //반복문의 변수 출력
	}

}
