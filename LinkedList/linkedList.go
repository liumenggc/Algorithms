package main

type Node struct {
	data int
	next *Node
}

type ListedList struct {
	head *Node
}

func (list *ListedList) InsertFirst(i int)  {
	data := &Node{data: i}
	if list.head != nil {
		data.next = list.head
	}
	list.head =data
}