package main

import (
	"fmt"
	"goLearning/tree"
)

//在GO里，我们使用CamelCase
//首字母大写说明 public
//首字母消息， 则为private
//这个public和private是针对package来说的 （package main）
//每个目录一个包，但和java不同，包名可以和目录名不同，但每个目录只能有一个包
//main包包含了一个可执行入口，里面有一个main函数，
//所以若目录中有一个main函数，则该目录下只能有一个main包，否则我们就可以给目录取其他的名字
//为struct定义的方法必须放在同一个包内，但可以是不同文件

func main() {
	//GO中值接收者和指针接收者用哪个呢
	//改变内容一定要指针
	//结构过大也考虑指针
	//一致性：如有指针接收者，最好都是指针接收者
	//值接受者是GO特有的，别的语言莫得，别的只有指针接收者
	//值/指针接收者可以接受值/指针都ok，所以定义方法的人可以随便改接收者为值or指针接收者，不影响调用者

	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) // no more -> like root.right->left, all are .
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse()
	fmt.Println()

	root.Right.Left.Print()

	root.Print()
	root.SetValue(100)
	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()
	fmt.Println()
	root.Traverse()
	// we can also create a slice
	//nodes := []TreeNode{
	//	{value: 4},
	//	{},
	//	{5, nil, &root},
	//}
	//fmt.Println(nodes)
}
