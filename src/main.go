package main

import "fmt"

var message string

const (
	first  = "the first"
	second = "the second"
	third  = "the third"
)

func main() {
	fmt.Println(message)
	declareInt()
	declareFloat()
	declareComplex()
}

func init() {
	message = "Hello World!"
}

func declareInt() {
	var myInt int
	myInt = 42
	fmt.Println(myInt)
}

func declareFloat() {
	var myFloat32 float32 = 42.
	println(myFloat32)

	myString := "Hello World"
	println(myString)
}

func declareComplex() {
	myComplex := complex(2, 3)
	println(myComplex)
}
