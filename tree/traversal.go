package tree

func (node *Node) Traverse() {
	//in GO, you can check null only once at the entrance
	if node == nil {
		return
	}
	node.Left.Traverse() //you should write if(node.Left!= null) node.Left.Traverse() in java or C++
	node.Print()
	node.Right.Traverse()
}
