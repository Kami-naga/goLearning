package tree

import "fmt"

//no inheritance, no polymorphism, only encapsulation
// so in GO, we only have struct, no class
type Node struct {
	Value       int
	Left, Right *Node
}

//no construction func in GO
//if you want to control the construction part, just use factory mode

func CreateNode(value int) *Node {
	//here you returned a local variable, it will cause problem in C++
	//but in GO, it's OK
	// so here comes the question, where is the treenode? stack or heap?
	//in GO ,you don't need to know, it depends on the runtime env.
	//compiler can decide where to put it
	return &Node{Value: value}
}

// the func of struct
//the difference is that before the name of the function, you give it a receiver(like this,self pointer in other languages like python)
//and it means that receiver can use this function
//e.g. func (xx X) foo() ---> xx.Print()
//     func foo(xx X) ---> Print(xx)
// and what it does is just the same as the normal function, only syntax differences here
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

//注意，此处若写node TreeNode，则会由于GO值传递的原因，他会拷贝，若不接它的返回值。则无法对外部产生什么影响
//只有指针才可以改变结构内容
//因此，一定要改成指针
//GO的编译器很聪明，你要指针他会拿指针，你要值他会从指针那里把值取出来
func (node *Node) SetValue(value int) {
	if node == nil {
		//非常神奇的是，nil指针也可以调用方法！！！
		fmt.Println("Setting Value to nil!")
		return
	}
	//nil指针虽然能调用方法，但如果要取一个nil指针的值，还是会炸的
	node.Value = value
}
