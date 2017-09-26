/*
	regexp包中含有三个函数用来判断是否匹配，如果匹配返回true 否则返回false
	fun Match(pattern string, b []byte) (matcher bool, error error)
	func MatchReader(oattern string, r io.RuneReader) (bool, error)
	func MatchString(pattern string, s string) (bool, error)

*/

package main

import (
	"fmt"
	"regexp"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
	} else if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("数字")
	} else {
		fmt.Println("不是数字")
	}
}