package queue

//type Queue []int //queue only support int type

// An FIFO queue.
type Queue []interface{} //queue can support all types!

//change q, so must be pointer receiver

// Pushes the element into the queue.
// 	e.g. q.Push(123) (add some space before you write the doc, then the view will changed -> in a block)
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

// Pops element from head.
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//下层interface{},上层限定int
//func (q *Queue) Push(v int) {
//	*q = append(*q, v)
//}
//
//func (q *Queue) Pop() int {
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head.(int)
//}

//下层interface{},上层限定int,但接口为interface{}呢？
//func (q *Queue) Push(v int) {
//	*q = append(*q, v.(int)) //但会在push其他类型参数时报运行时错误，虽然编译没错误
//}
//
//func (q *Queue) Pop() int {
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head.(int)
//}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
