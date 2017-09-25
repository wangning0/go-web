/*
	Go语言追求简洁优雅，所以Go语言不支持传统的try...catch...finally这种异常

	Go语言设计者们认为，将异常与控制结构混在一起会很容易使得代码变得混乱，

	Go中引入的Exception处理： defer panic recover
*/

// Go中可以先抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理

package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println("e")

		}
		fmt.Println("f")
	}()
	f()
}

func f() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b")
	fmt.Println("c")
}

// a d e f