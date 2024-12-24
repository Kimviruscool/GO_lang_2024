package main

import "fmt"

func main() {
	fmt.Println("print test")
	fmt.Println("this is Go print test")
	변수() //함수 실행방법
	상수()
}

func 변수() { //함수 생성 // 변수 : 프로그램이 실행되는 동안 변경이 가능한 값
	//var a int 변수 정수 선언 방법
	//var f float32 = 11. 소수점 사용 변수 선언
	//var i, j, k int = 1, 2 ,3
	var i = 1
	var s = "Hi"
	fmt.Println(i, s)
}

func 상수() { //상수 : 프로그램이 실행되는 동안 변하지 않는 값
	const c int = 10
	const s string = "hi"
	fmt.Println(c, s)
}
