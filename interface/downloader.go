package main

import (
	"fmt"
	"goLearning/interface/infra"
)

//func getRetriever() infra.Retriever {
//	return infra.Retriever{} //若想返回testing.Retriever就发现不行，所以有一个类型耦合
//}

func getRetriever() retriever {
	return infra.Retriever{} //现在可以随心所欲改 infra, testing都可以
}

type retriever interface { //Something that can get
	//这里可以发现我们定义了接口但并没有给出接口的实现，他并没有地方说实现了该接口，他是怎么找到实现的呢？这里就是GO相较Java的特殊&先进之处了
	//这种做法好在哪里呢？
	Get(string) string
}

func main() {
	//类型仍耦合
	//Go为强类型系统
	//对于弱类型系统or动态绑定语言系统来说（php,python,javascript等），变量写的时候，程序是不知道类型的，
	//只有运行到这里，传入时才知道类型，所以对于动态语言而言，做到这里已是解耦了
	//但对于静态语言而言，还有个类型需要去处理（下边的infra.Retriever）
	//var retriever infra.Retriever = getRetriever() //若换个Retriever，这里的类型也要改

	//var retriever ？ = getRetriever() //所以这里的类型给改成别的啥，让代码和逻辑一致，要个东西能取url
	//interface!! 如此一来外边类型那儿就全是接口了，真正传哪个类型可以在getRetriever中随意更改
	var r retriever = getRetriever()
	fmt.Println(r.Get("https://www.imooc.com"))
}
