package main

import "fmt"

func shout(singCount int) {
	for i := 1; i <= singCount; i++ {
		pants := i * 10
		sword := pants + 10
		fmt.Printf("타잔이 %d원짜리 팬티를 입고 %d원짜리 칼을 차고 노래를 한다.\r\n", pants, sword)
	}
}

func main() {
	shout(10)
}
