package main

import "fmt"

//no inheritance, no polymorphism, only encapsulation
// so in GO, we only have struct, no class
type treeNode struct {
	value       int
	left, right *treeNode
}

//no construction func in GO
//if you want to control the construction part, just use factory mode

func createNode(value int) *treeNode {
	//here you returned a local variable, it will cause problem in C++
	//but in GO, it's OK
	// so here comes the question, where is the treenode? stack or heap?
	//in GO ,you don't need to know, it depends on the runtime env.
	//compiler can decide where to put it
	return &treeNode{value: value}
}

func main() {
	//GO中值接收者和指针接收者用哪个呢
	//改变内容一定要指针
	//结构过大也考虑指针
	//一致性：如有指针接收者，最好都是指针接收者
	//值接受者是GO特有的，别的语言莫得，别的只有指针接收者
	//值/指针接收者可以接受值/指针都ok，所以定义方法的人可以随便改接收者为值or指针接收者，不影响调用者
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) // no more -> like root.right->left, all are .
	root.left.right = createNode(2)
	root.right.left.setValue(4)
	root.traverse()
	fmt.Println()

	root.right.left.print()

	root.print()
	root.setValue(100)
	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()
	fmt.Println()
	root.traverse()
	// we can also create a slice
	//nodes := []treeNode{
	//	{value: 4},
	//	{},
	//	{5, nil, &root},
	//}
	//fmt.Println(nodes)
}

// the func of struct
//the difference is that before the name of the function, you give it a receiver(like this,self pointer in other languages like python)
//and it means that receiver can use this function
//e.g. func (xx X) foo() ---> xx.print()
//     func foo(xx X) ---> print(xx)
// and what it does is just the same as the normal function, only syntax differences here
func (node treeNode) print() {
	fmt.Println(node.value, " ")
}

//注意，此处若写node treeNode，则会由于GO值传递的原因，他会拷贝，若不接它的返回值。则无法对外部产生什么影只有指针响
//只有指针才可以改变结构内容
//因此，一定要改成指针
//GO的编译器很聪明，你要指针他会拿指针，你要值他会从指针那里把值取出来
func (node *treeNode) setValue(value int) {
	if node == nil {
		//非常神奇的是，nil指针也可以调用方法！！！
		fmt.Println("Setting value to nil!")
		return
	}
	//nil指针虽然能调用方法，但如果要取一个nil指针的值，还是会炸的
	node.value = value
}

func (node *treeNode) traverse() {
	//in GO, you can check null only once at the entrance
	if node == nil {
		return
	}
	node.left.traverse() //you should write if(node.left!= null) node.left.traverse() in java or C++
	node.print()
	node.right.traverse()
}
