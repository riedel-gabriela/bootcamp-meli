package main

import (
	"fmt"
)

func main() {
	// //apply “defer” to the invocation of an anonymous function
	// defer func() {
	// 	fmt.Println("This function is executed despite a panic occurring")
	// }()
	// //create a panic with a message that it occurred
	// // os.Exit(2)
	// panic("panic occurred!!!")
	num := 3
	isPair(num)
	fmt.Println("execution completed")
}

func isPair(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if num%2 != 0 {
		panic("not an even number")
	}

	fmt.Println(num, "is an even number!")
}
