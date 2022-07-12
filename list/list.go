package list

// 單向鏈結串列
type SinglyLinkedList struct {
	// 指向第一個節點
	first *Node
	// 紀錄目前串列節點數
	length int
}

// 建立新串列
func New() *SinglyLinkedList {
	return new(SinglyLinkedList)
}

// 取得節點數
func (l *SinglyLinkedList) Length() int {
	return l.length
}

// 清空節點
func (l *SinglyLinkedList) Clear() {
	l.first = nil
	l.length = 0
}

// 取得第一個節點
func (l *SinglyLinkedList) First() *Node {
	return l.first
}

// 取得最後一個節點
func (l *SinglyLinkedList) Last() *Node {
	for node := l.first; node != nil; node = node.Next() {
		if node.Next() == nil {
			return node
		}
	}
	return nil
}

// 取得索引i的節點
func (l *SinglyLinkedList) Get(i int) *Node {
	if i == l.length-1 {
		return l.Last()
	}

	if i < 0 || i >= l.length {
		panic("index out of range")
	}

	node := l.First()
	for j := 0; j < i; j++ {
		node = node.Next()
	}
	return node
}

// 在串列最前插入新的值
func (l *SinglyLinkedList) InsertFirst(v string) *Node {
	node := &Node{
		Value: v,
		node:  l.first,
	}

	l.first = node
	l.length++
	return node
}

// 在串列最後插入新的值
func (l *SinglyLinkedList) InsertLast(v string) *Node {
	node := &Node{
		Value: v,
	}

	l.Get(l.length - 1).node = node
	l.length++
	return node
}

// 在串列索引i的位置插入新的值
func (l *SinglyLinkedList) Insert(i int, v string) *Node {
	if i == 0 {
		return l.InsertFirst(v)
	}
	if i == l.length {
		return l.InsertLast(v)
	}

	next := l.Get(i)
	prev := l.Get(i - 1)
	node := &Node{
		Value: v,
		node:  next,
	}
	prev.node = node
	l.length++
	return node
}

// 刪除索引i的節點
func (l *SinglyLinkedList) Delete(i int) *Node {
	if i == 0 {
		deletedNode := l.first
		l.first = l.first.node
		deletedNode.node = nil
		l.length--
		return deletedNode
	}
	if i == l.length-1 {
		deletedNode := l.Last()
		l.Get(l.length - 2).node = nil
		l.length--
		return deletedNode
	}
	deletedNode := l.Get(i)
	next := deletedNode.node
	prev := l.Get(i - 1)
	prev.node = next
	deletedNode.node = nil
	l.length--
	return deletedNode
}

// 鏈結串列的節點
type Node struct {
	// 資料欄位，儲存資料
	Value string
	// 節點欄位，儲存下個節點的指標
	node *Node
}

// 取得下一個節點
func (n *Node) Next() *Node {
	return n.node
}
