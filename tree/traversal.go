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

func (node *Node) Traverse2() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
}

func (node *Node) TraverseFunc(f func(node *Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
