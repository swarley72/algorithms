type Node struct {
	Key   int
	Next  *Node
	Prev  *Node
	Value int
}

type LRUCache struct {
	Capacity int
	HashMap  map[int]*Node
	Head     *Node
	Tail     *Node
}

func Constructor(capacity int) LRUCache {
	hashMap := make(map[int]*Node, capacity)
	head := &Node{}
	tail := &Node{}
	// дувусвязный список
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Capacity: capacity,
		HashMap:  hashMap,
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) Remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) AddToHead(node *Node) {
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.HashMap[key]
	if !ok {
		return -1
	}

	this.Remove(node)
	this.AddToHead(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.HashMap[key]
	if ok {
		this.Remove(node)
		this.AddToHead(node)
		node.Value = value
		return
	}

	// вторая проверка возможно лишняя если по условию capacity >= 1
	if len(this.HashMap) == this.Capacity && len(this.HashMap) > 0 {
		node := this.Tail.Prev
		this.Remove(node)
		delete(this.HashMap, node.Key)
	}

	newNode := &Node{Key: key, Value: value}
	this.HashMap[key] = newNode
	this.AddToHead(newNode)
}
