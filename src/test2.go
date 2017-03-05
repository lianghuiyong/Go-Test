package main

import "fmt"

func main() {

	fmt.Println("--------------------test_for--------------------")
	test_for()

	fmt.Println("--------------------test_if--------------------")
	test_if()

	fmt.Println("--------------------test_switch--------------------")
	test_switch()

	fmt.Println("--------------------test_defer--------------------")
	test_defer()
}

func test_for() {
	sum := 0
	for i := 0; i < 10; i ++ {
		sum += i
	}

	fmt.Println("sum:", sum) // sum: 45
}

func test_if() {

	//表达式无需小括号
	if true {
		fmt.Println("is test if")
	} else {
		fmt.Println("is test if")
	}

	//表达式可执行一部分简单的语句（赋值、函数等）
	if one := 1; true {
		fmt.Println("is test if ", one, " teue")
	} else {
		fmt.Println("is test if")
	}

	//简单的调用
	if test_for(); false {
		fmt.Println("is test if ", "test_for")
	} else {
		fmt.Println("is test if")
	}
}

func test_switch() {

	//同样可执行简单语句，而且case可为bool值！
	switch i5 := 5; true {
	case i5 > 5:
		fmt.Println("i5 > 5")
	case i5 == 5:
		fmt.Println("i5 == 5")
	case i5 < 5:
		fmt.Println("i5 < 5")
	}

	i5 := 5
	switch  i5 {
	case 0:
		fmt.Println("i5 is 0")
	case 1:
		fmt.Println("i5 is 1")
	case 2:
		fmt.Println("i5 is 2")
	case 3:
		fmt.Println("i5 is 3")
	case 4:
		fmt.Println("i5 is 4")
	case 5:
		fmt.Println("i5 is 5")
	}
}

//defer会将函数进行入栈，当包含defer的函数返回时，被入栈的函数将一一出栈。
func test_defer()  {

	test_defer_hello()

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func test_defer_hello() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

/*
--------------------test_for--------------------
sum: 45
--------------------test_if--------------------
is test if
is test if  1  teue
sum: 45
is test if
--------------------test_switch--------------------
i5 == 5
i5 is 5
--------------------test_defer--------------------
hello
world
counting
done
9
8
7
6
5
4
3
2
1
0
*/