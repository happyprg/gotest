package main

import "fmt"

func Hanoi(depth int) {
	fmt.Printf("Number of disks:%d  \r\n", depth)
	col1 := make([]int, depth)
	col2 := make([]int, depth)
	col3 := make([]int, depth)

	fmt.Printf("len:%d, cap:%d, col1:%d  \r\n", len(col1), cap(col1), col1)
	fmt.Printf("len:%d, cap:%d, col2:%d  \r\n", len(col2), cap(col2), col2)
	fmt.Printf("len:%d, cap:%d, col3:%d  \r\n", len(col3), cap(col3), col3)

	//for i := 0; i <= i; i++ {
	//	col1[i] = 1
	//	col2[i] = 0
	//	col3[i] = 0
	//}
	//
	//for i := 0; i <= i; i++ {
	//	fmt.Printf("col1-%d, col2-%d, col3-%d", col1[i], col2[i], col3[i])
	//}

	//for i := 1; i <= singCount; i++ {
	//	pants := i * 10
	//	sword := pants + 10
	//	fmt.Printf("타잔이 %d원짜리 팬티를 입고 %d원짜리 칼을 차고 노래를 한다.\r\n", pants, sword)
	//}
}

func main() {
	Hanoi(3)
}
