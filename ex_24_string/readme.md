# 字符串处理

## 字符串操作

下面这些函数来自于strings包

* `func Contains(s , substr string) bool`

    字符串中是否包含substr，返回bool值

* `func Join(a []string, sep string) string`

    字符串连接，把slice a通过sep连接起来

* `func Index(s, sep string) int`

    字符串s中查找sep所在的位置，返回位置值，找不到返回-1

* `func Repeat(s string, count int) string`

    重复s字符串count次，最后返回重复的字符串

* `func Replace(s, old, new string, n int) string`

    在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换

* `func Split(s, sep string) []string`

    把s字符串按照sep分割 返回slice

* `func Trim(s string, cutset) string`

    在s字符串的头部和尾部去除cutset指定的字符串

* `func Fields(s string) []string`

    去除s字符串的空格服，并且按照空格分割返回slice


## 字符串转换

字符串转化的函数在strconv中，如下是一些常用的

* Append系列函数将整数转换为字符串后，添加到现有的字节数组中

