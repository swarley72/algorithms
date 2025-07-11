type Node struct {
	Value int
	Next  *Node
}

type MyLinkedList struct {
	Capacity int
	Head     *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Capacity {
		return -1
	}

	current := this.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}

	return current.Value
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Capacity, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Capacity {
		return
	}
	this.Capacity++

	if index == 0 {
		newNode := &Node{val, this.Head}
		this.Head = newNode
		return
	}

	current := this.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	oldNext := current.Next
	current.Next = &Node{val, oldNext}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Capacity {
		return
	}

	this.Capacity--

	if index == 0 {
		this.Head = this.Head.Next
		return
	}

	current := this.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
}
