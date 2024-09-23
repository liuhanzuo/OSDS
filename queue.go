package main

import "fmt"

type queue struct {
	items    []string
	cap int
	sz  int
}

func (q *queue) init(capacity int) {
	q.items = make([]string, 0, capacity)
	q.cap = capacity
	q.sz = 0
}

func (q *queue) enqueue(item string) (int) {
	if q.sz == q.cap {
		fmt.Println("queue is full, cannot enqueue", item)
		return 1
	}
	q.items = append(q.items, item)
	q.sz++
	return 0
}

func (q *queue) dequeue() (string, int) {
	if q.sz == 0 {
		fmt.Println("queue is empty, cannot dequeue")
		return "", 1
	}
	item := q.items[0]
	q.items = q.items[1:]
	q.sz--
	return item, 0
}

func (q *queue) size() int {
	return q.sz
}

func (q *queue) capacity() int {
	return q.cap
}

func main() {
	q := queue{}
	q.init(2)
	q.enqueue("mlr")
	fmt.Println(q.size())
	q.enqueue("wsq")
	fmt.Println(q.size())
	fmt.Println(q.capacity())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}
