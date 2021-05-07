package main

import (
	"fmt"
	"goLearning/retriever/mock"
	"goLearning/retriever/real"
	"time"
)

const url = "https://www.imooc.com"

type Retriever interface { //retriever为实现者
	//interface中无需加func，因为里头全是func
	//接口的实现是隐式的，只要实现接口里的方法
	Get(url string) string //接口由使用者来定义（这点和传统方式不同，传统为实现者实现）
}

//接口的组合
type Poster interface {
	Post(url string, form map[string]string) string
}

//因此在GO中
//interface{}可以表示任何类型
func download(r Retriever) string { //downloader为使用者
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name": "bbb",
			"age":  "18",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "this is a fake retriever"}
	r = &retriever

	inspect(r)

	r = &real.Retriever{ //接口几乎不会用到接口的指针
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	fmt.Printf("%T %v\n", r, r)

	//查看接口类型和值三种方法
	//1. print 如上fmt.Printf("%T %v\n", r, r)
	//2. switch v := r.(type)
	// type assertion
	inspect(r)
	//Type Assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever!")
	}
	//fmt.Println(download(r))
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
	//接口变量内含有实现者的类型，以及实现者的值或指针，指针指向一个实现者）
	//所以接口变量不要去用他的地址&
	//接口变量同样采用值传递，几乎不需要使用接口的指针
	//指针接收者实现只能以指针方式去使用，值接收者两者皆可
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
