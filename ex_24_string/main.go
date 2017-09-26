/*
	经常需要对字符串进行分割、连接、转换等操作

	可以通过Go标准库中的strings和strconv两个包中的函数来讲解如何进行有效快速的操作
*/


package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(strings.Contains("winger", "wing"))
	fmt.Println(strings.Contains("winger", "Wing"))
	fmt.Println(strings.Contains("winger", ""))
	fmt.Println(strings.Contains("", ""))

	s := []string{"foo", "bar", "hello"}
	fmt.Println(strings.Join(s, "-"))

	fmt.Println(strings.Index("hello", "o"))
	fmt.Println(strings.Index("hello", "ww"))

	fmt.Println("ba" + strings.Repeat("na", 10))

	fmt.Println(strings.Replace("hello good", "o", " nihao", 1))

	fmt.Println(strings.Split("a, b, c", ","))

	fmt.Println(strings.Trim("!!!test!!!!", "!"))

	fmt.Println(strings.Fields(" cdcd "))

	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	// a := strconv.FormatBool(false)
	// b := strconv.FormatFloat(123.23, 'g', 12, 64)
	// c := strconv.FormatInt(1234, 10)
	// d := strconv.FormatUint(12345, 10)
	// e := strconv.Itoa(1023)
	// fmt.Println(a, b, c, d, e)

	a, _ := strconv.ParseBool("false")
	b, _ := strconv.ParseFloat("123.23", 64)
	c, _ := strconv.ParseInt("1234", 10, 64)
	d, _ := strconv.ParseUint("12345", 10, 64)
	e, _ := strconv.Atoi("1023")
	fmt.Println(a, b, c, d, e)

}