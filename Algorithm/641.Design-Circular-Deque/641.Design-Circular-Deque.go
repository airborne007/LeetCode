package designcirculardeque

type linkNode struct {
	val  int
	prev *linkNode
	next *linkNode
}

type MyCircularDeque struct {
	head, tail *linkNode
	size, cap  int
}

func Constructor(k int) MyCircularDeque {
	q := MyCircularDeque{
		head: &linkNode{},
		tail: &linkNode{},
		cap:  k,
	}
	q.head.next = q.tail
	q.tail.prev = q.head
	return q
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}
	node := &linkNode{val: value}
	temp := q.head.next
	q.head.next = node
	node.prev = q.head
	node.next = temp
	temp.prev = node
	q.size++
	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}
	node := &linkNode{val: value}
	temp := q.tail.prev
	q.tail.prev = node
	node.next = q.tail
	node.prev = temp
	temp.next = node
	q.size++
	return true
}

func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}
	temp := q.head.next
	q.head.next = temp.next
	temp.next.prev = q.head
	q.size--
	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}
	temp := q.tail.prev
	q.tail.prev = temp.prev
	temp.prev.next = q.tail
	q.size--
	return true
}

func (q *MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}
	return q.head.next.val
}

func (q *MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.tail.prev.val
}

func (q *MyCircularDeque) IsEmpty() bool {
	return q.size == 0
}

func (q *MyCircularDeque) IsFull() bool {
	return q.size == q.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
