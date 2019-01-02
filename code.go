package main

import "math"
import "fmt"

//定义变量

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n",a, s)
}

//常亮
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)

	var c int
	c = int(math.Sqrt(a * a + b * b ))
	fmt.Println(filename, c)
}

//枚举
func enums() {
    //普通枚举类型
	const (
		php = 1
		jarv = 2
	)

	//自增值枚举类型
	const (
		cpp = iota
		_
		java
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main()  {
	fmt.Println("Hello world")
	variableZeroValue()
	consts()
	enums()
} 