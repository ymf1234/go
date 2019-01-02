package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)
func convertToBin(n int) string {
	var  result string
	if n == 0 {
		return "0"
	}
	for  ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forerve() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5), // 101
		convertToBin(13), // 1101
		convertToBin(0),
	)

	printFile("code.go")
}