package main

import "fmt"

func lengthOfNinRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastOccurred[ch] >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}
func main() {
	fmt.Println(lengthOfNinRepeatingSubStr("aaaaaa"))
	fmt.Println(lengthOfNinRepeatingSubStr("abc"))
	fmt.Println(lengthOfNinRepeatingSubStr(""))
	fmt.Println(lengthOfNinRepeatingSubStr("b"))
	fmt.Println(lengthOfNinRepeatingSubStr("lololollol"))
	fmt.Println(lengthOfNinRepeatingSubStr("我困"))

}
