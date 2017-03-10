package main

import (
	"fmt"
	"encoding/json"
	"log"
)

func main() {
	test_array()
	test_slice()
	test_map()
	test_json()
}

/*-----测试数组-----*/
func test_array() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)
	arr[0] = 3
	arr[1] = 3
	arr[2] = 3
	arr[3] = 3
	fmt.Println(arr)
}

/*------测试切片-----*/
func test_slice() {
	mouth := []string{1: "一月", 2: "二月", 3: "三月", 4: "四月", 5: "五月", 6: "六月", 7: "七月", 8: "八月", 9: "九月", 10: "十月", 11: "十一月", 12: "十二月"}

	slice1 := mouth[1:3]
	slice2 := mouth[2:5]

	fmt.Println(slice1)
	fmt.Println(slice2)

	fmt.Println(mouth)
	reverse(mouth[:])
	fmt.Println(mouth)

	fmt.Println(slice1)
	fmt.Println(slice2)

}

/*--------数组反转--------*/
func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func test_map() {
	//创建map1
	map1 := map[int]string{}

	//初始化赋值
	map1[1] = "hello1"
	map1[2] = "world2"
	map1[3] = "world3"
	map1[4] = "world4"
	map1[5] = "world5"
	map1[6] = "world6"
	map1[7] = "world7"
	map1[8] = "world8"
	map1[9] = "world9"
	map1[10] = "!!!!!!10"

	//取值
	fmt.Println(map1)
	fmt.Println(map1[1])

	// Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。
	// 在实践中，遍历的顺序是随机的，每一次遍历的顺序都不相同。这是故意的，
	// 每次都使用随机的遍历顺序可以强制要求程序不会依赖具体的哈希函数实现。
	for _, value := range map1 {
		println("\t", value, "\n")
	}

	//删除
	delete(map1, 2)
	fmt.Println(map1)
	fmt.Println(map1[4] + "1")
}

func test_json() {
	/*	type BaseResponse struct {
			code int
			msg string
			data T
		}*/

	type BaseResponse struct {
		code int
		msg  string
		//data T
	}

	type Movie struct {
		Title  string
		Year   string
		Actors []  string
	}

	movies := []Movie{
		{"金刚狼", "2017", []string{"休·杰克曼", "达芙妮·基恩", "帕特里克·斯图尔特"}},
		{"极限特工", "2017", []string{"范·迪塞尔", "露比·罗丝", "妮娜·杜波夫"}},
		{"功夫瑜伽", "2017", []string{"成龙", "李治廷", "张艺兴"}},
		{"生化危机：终章", "2017", []string{"米拉·乔沃维奇", "伊恩·格雷", "艾丽·拉特"}},
	}

	//一个是不带缩进显示的，一个是带缩进显示的
	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "    ")

	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

}
