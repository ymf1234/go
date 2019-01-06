package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}
//切片
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:6]
	fmt.Println("arr[2:] = ", arr[2:])
	s2 := arr[:]
	fmt.Println("arr[:] = ", arr[:])

	updateSlice(s1)
	fmt.Println(arr)

	s2 = s2[:5]
	fmt.Println(s2)

	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))


	//切片操作
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 =", s3, s4, s5)
	fmt.Println("arr =", arr)
}