/*包名*/
package main

/*导包*/
import (
	"fmt"
)

func main() {
	fmt.Println("Go : Hello World!")
	fmt.Println(add(25, 25))
	fmt.Println(swap("huiyong", "liang"))
	fmt.Println(split(25))
}

/*
1、参数类型相同(x int, y int)，可简写成(x, y int)
2、返回类型为 int
*/
func add(x, y int) int {
	return x + y
}

/*
1、返回类型为 (string, string)，
*/
func swap(x, y string) (string, string) {
	return y, x
}

/**
1、返回类型(x, y int)，已经声明了返回变量x，y
2、指定了返回变量时，return x, y 可简写成 return、
3、 `:=` 简洁赋值语句用在明确类型的地方，用于替代 `var`

等同
func split(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
