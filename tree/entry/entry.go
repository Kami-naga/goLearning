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

	//var root tree.Node
	//root = tree.Node{Value: 3}
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) // no more -> like root.right->left, all are .
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Traverse2()
	fmt.Println()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)
	//myRoot := myTreeNode{&root}
	//myRoot.postOrder()
	root.postOrder()

	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("Max node value: ", maxNode)
	//root.Right.Left.Print()
	//
	//root.Print()
	//root.SetValue(100)
	//pRoot := &root
	//pRoot.Print()
	//pRoot.SetValue(200)
	//pRoot.Print()
	//fmt.Println()
	//root.Traverse()
	// we can also create a slice
	//nodes := []TreeNode{
	//	{value: 4},
	//	{},
	//	{5, nil, &root},
	//}
	//fmt.Println(nodes)
}

//GO无继承，如何扩展别人写的包呢？3种办法
//1. 别名法（见container中的queue和sliceops）该办法最简单，但若之后需要过渡到组合法改起来很麻烦
//2. 组合法 最常见 除非确保以后不动了的代码用别名法更快乐，其余情况还是组合吧
//2.1 内嵌法 Embedding 特殊的语法糖 好处是能少写许多代码 但要注意 这种写法得大家都是熟练工的时候这么写，否则会让萌新们感到混乱的
//组合法把组合的结构体的名字去掉，就留个类型，就成内嵌了
//那没了名字，我们怎么找到它呢？GO会给他们起个名字，名字就是类型最后一个.后边的东西， 此处*tree.Node，即为Node
//然后此时还有个语法糖，访问外头那层直接可以.到里头包的（成员变量&方法都有）
//这玩意儿和继承很像呢!但它不是哦~本质还是组合（子类指针指向父类，可以发现指不过去，真想这么干，用接口!）
//那 有重载吗？确实类似，定义一个同名的函数，则调用xx.原来函数()时，原来的会被新的覆盖掉（shadowed），但你仍可以通过xx.Node.原来函数()来访问原来的
type myTreeNode struct {
	*tree.Node //此即内嵌，原为 node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}
