package main

import "fmt"

func main() {
	m := map[string]string {
		"name" : "ccmouse",
		"course": "golang",
		"site": "mmm",
		"quality": "notbad",
	}

	m2 := make(map[string]string) //m2 == empty map

	var m3 map[string]int // m3 == nil
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, err := m["course"]
	fmt.Println(courseName, err)

	if coursName, err := m["cours"]; err {
		fmt.Println(coursName, err)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")

	name, ok = m["name"]
	fmt.Println(name, ok)

}