package main

import "fmt"

type queue struct {
	items    []string
	cap int
	sz  int
}

func (q *queue) Init(capacity int) {
	q.items = make([]string, 0, capacity)
	q.cap = capacity
	q.sz = 0
}

func (q *queue) Enqueue(item string) (int) {
	if q.sz == q.cap {
		fmt.Println("queue is full, cannot enqueue", item)
		return 1
	}
	q.items = append(q.items, item)
	q.sz++
	return 0
}

func (q *queue) Dequeue() (string, int) {
	if q.sz == 0 {
		fmt.Println("queue is empty, cannot dequeue")
		return "", 1
	}
	item := q.items[0]
	q.items = q.items[1:]
	q.sz--
	return item, 0
}

func (q *queue) Size() int {
	return q.sz
}

func (q *queue) Capacity() int {
	return q.cap
}

func main() {
	q := queue{}
	q.init(2)
	q.enqueue("mlr")
	fmt.Println(q.Size())
	q.enqueue("wsq")
	fmt.Println(q.Size())
	fmt.Println(q.Capacity())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
