package main

import "fmt"
var x int = 43


func main(){
	fmt.Println("Closures")
	// a := 0
	// increment := func() int{
	// 	a++
	// 	return a
	// }
	// fmt.Println(increment())
	// fmt.Println(increment())
	// increment := wrapper()
	// fmt.Println(increment())
	// fmt.Println(increment())
	// addSix := createBase(6)
	// fmt.Println(addSix(21))

	c := counter()
	fmt.Println(c(5))
	fmt.Println(c(7))
	fmt.Println(c(10))
}

func createBase(base int) func(num int) int{

	return func(num int) int {
		return base + num
	}
}

func counter() func(num int) int{
	count := 0

	return func (num int) int {
		count += num
		return count
	} 
}

func wrapper() func() int{
	x := 0
	return func() int {
		x++
		return x
	}
}

func printX(){
	fmt.Println(x)
}

