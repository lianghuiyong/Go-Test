package main

import "fmt"

func main() {
	test_array()
	test_slice()
}

/*-----测试数组-----*/
func test_array() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)
	arr[0]= 3
	arr[1]= 3
	arr[2]= 3
	arr[3]= 3
	fmt.Println(arr)
}

/*------测试切片-----*/
func test_slice() {

}
