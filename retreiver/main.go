package main

import (
	"basic/retreiver/mock"
	"basic/retreiver/real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url  = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name": "ccmouse",
			"course": "golang",
		})
}

//接口组合
type RetrieverPoster interface {
	Retriever
	Poster
}
func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked ",
	})

	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is"}

	r = &retriever
	//Type assertion
	mockRetriever := r.(*mock.Retriever)
	fmt.Println(mockRetriever.Contents)

	inspect(r)
	r = real.Retriever{
		UserAgent: "Mozilla/5.0",
	}
	inspect(r)

	//Type assertion
	realRetriever := r.(real.Retriever)
	fmt.Println(realRetriever.TimeOut)
	//fmt.Println(download(r))

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
