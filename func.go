package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a, b int, op string) int  {
	switch op {
	case "+":
		return a + b 
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation:" + op)
	}
}

func div(a,b int) (q, r int) {
	// return a / b, a % b
	q = a /b
	r = a % b
	return
}

func applv(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args " + "(%d , %d)", opName, a, b)
	return op(a, b)
}

func pow (a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(applv(pow, 3, 4))
}