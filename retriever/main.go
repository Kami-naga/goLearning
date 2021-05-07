package main

import (
	"fmt"
	"goLearning/retriever/mock"
	"goLearning/retriever/real"
)

type Retriever interface { //retriever为实现者
	//interface中无需加func，因为里头全是func
	//接口的实现是隐式的，只要实现接口里的方法
	Get(url string) string //接口由使用者来定义（这点和传统方式不同，传统为实现者实现）
}

func download(r Retriever) string { //downloader为使用者
	return r.Get("https://www.imooc.com")
}
func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake retriever"}
	r = real.Retriever{}
	fmt.Println(download(r))

}
