package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var a int // declaration
	// a = 9 // assign value
	// var a int = 9 // declaratiom & assign value
	a := 9 // short form. declaratiom & assign value
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	b := 2.71
	//c := 'Z'
	//d, e := 4.10, 3.99
	//f := "문자열"
	// fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(d),
	// reflect.TypeOf(e), f, reflect.TypeOf(f))
	K := "변수명이 대문자로 시작하면 다른 패키지에서 이 변수를 사용할수 없음"
	// Koreanzombie := "정찬성"
	koreanzombie := "정찬성"

	fmt.Println(float64(a) * b)
	fmt.Println(a * int(b))
	fmt.Println(float64(a) > b)
	var g int
	var h float32
	var i bool
	var j string
	fmt.Println(j, g, h, i, K, koreanzombie)
	fmt.Printf("%s%f%d%s", j, g, h, i)
	// fmt.Println(reflect.TypeOf('Z')) // rune, int32
	// fmt.Println(reflect.TypeOf(99))
	// fmt.Println(reflect.TypeOf(2.7))
	// fmt.Println(reflect.TypeOf(false))
	// fmt.Println(reflect.TypeOf("go!"))
	// fmt.Println('A', 'a', '0', '김', '인', '하')
	// fmt.Println(math.Floor(3.17))
	// fmt.Println(strings.Title("head first go"))
}
