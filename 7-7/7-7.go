// 7-7
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool) //동기 채널 생성
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			done <- true                //고루틴에 true를 보냄. 값을 꺼낼 떄까지 대기
			fmt.Println("고루틴 : ", i)    //반복문의 변수 출력
			time.Sleep(1 * time.Second) //1초 대기
		}
	}()

	for i := 0; i < count; i++ {
		<-done                    //채널에 값이 들어올 떄까지 대기, 값을 꺼냄
		fmt.Println("메인함수 : ", i) //반복문의 변수 출력
	}

}
