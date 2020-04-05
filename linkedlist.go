package linkedCycle

type ListNode struct {
	value int
	next  *ListNode
}

type List struct {
	value int
	head  *ListNode
	tail  *ListNode
}

// Add element to the list
func (p *List) add(value int) {
	if p.head == nil {
		p.head = new(ListNode)
		p.head.value = value
		p.tail = p.head
		return
	}
	node := new(ListNode)
	p.tail.next = node
	p.tail = node
	node.value = value
}

// Point tail to some node
func (p *List) point(to int) {
	if to == -1 {
		return
	}
	where := p.head
	for i := 0; i < to; i++ {
		where = where.next
	}
	p.tail.next = where
}

// Pack arr elements to list
func FromArray(arr []int, pos int) *List {
	list := new(List)
	for _, item := range arr {
		list.add(item)
	}
	list.point(pos)
	return list
}

// Detect cylcle with a hashtable
func DetectCycleWithHash(l *List) bool {
	cur := l.head

	table := make(map[*ListNode]bool)
	for cur != nil {
		_, ok := table[cur]
		if ok {
			return true
		} else {
			table[cur] = true
		}
		cur = cur.next
	}
	return false
}

// Use two pointers advancing at different rates
// if there is a cycle - eventually they will meet
func DetectCycleTwoPointers(l *List) bool {
	first := l.head
	second := l.head
	if first == nil {
		return false
	}
	// there is only one self-looped node

	for first != nil && second != nil && second.next != nil {
		first = first.next
		second = second.next.next
		if first == second {
			return true
		}
	}
	return false
}

// Just like DetectCyleTwoPointers, but returns cycle start
func CycleStart(l *List) *ListNode {
	first := l.head
	second := l.head
	if first == nil {
		return nil
	}
	// there is only one self-looped node

	for first != nil && second != nil && second.next != nil {
		first = first.next
		second = second.next.next
		if first == second {
			return first
		}
	}
	return nil
}
