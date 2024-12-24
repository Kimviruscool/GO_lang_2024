package main

import "fmt"

func main() {
	//데이터 타입
	/*
		bool : True / False 타입
		String : 문자열 타입
		int 8 ~ 64 : 정수형 타입
		unit 8~64 , ptr : 정수형 타입
		float32 ~64 : 복소수 타입
		complex64~ 128 : 복소수 타입
		byte : unit 8 과 동일하며 바이트 코드에 사용
		rune : int 32 와 동일하며 유니코드 코드포인트에 사용
	*/
	//bool
	var a bool = 1 > 3 //타입 지정
	b := 1 > 2         //타입 생략 자동 지정
	fmt.Println(a)     //출력
	fmt.Println(b)     //출력

	//String
	var c string = "문자열 타입"
	fmt.Println(c)
	d := "문자열타입"
	fmt.Println(d)

	//int , unit , float
	var e int = 1
	fmt.Println(e)
	f := 2
	fmt.Println(f)

	var g int16 = 12345               //5자리 1~10000까지
	var h int32 = 1234567890          //10 ~ 10억까지
	var i int64 = 1234567890123456789 // ~~ 경
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	var aa uint = 12345678901234567890
	fmt.Println(aa)

	var bb float32 = 0.1234567890123456789012345678901234567890123456789012345678901234567890123456789 //0.12345679
	fmt.Println(bb)
}
