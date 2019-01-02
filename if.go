package main

import "io/ioutil"
import "fmt"


func grade(score int ) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	// default:
	// 	panic(fmt.Sprintf(
	// 		"Wrong score: %d", score))
	}

	return g
}

func main() {
	const filename string = "code.go"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(22),
		grade(99),
		grade(100),
		grade(100),
	)
}