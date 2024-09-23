package main

import(
	"fmt"
	"sync"
)

type queue struct {
	items    []string
	cap int
	sz  int
	lock sync.Mutex
	notempty sync.Cond
	notfull sync.Cond
}
func (q *queue) Init(capacity int) {
	q.items = make([]string, 0, capacity)
	q.cap = capacity
	q.sz = 0
	q.notempty.L = &q.lock
	q.notfull.L = &q.lock
}

func (q *queue) Enqueue(item string) (int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.sz == q.cap {
		fmt.Println("Waiting enqueue thread", item)
		q.notfull.Wait()
	}
	if q.sz == q.cap {
		fmt.Println("queue is full, cannot enqueue", item)
		return 1
	}
	q.items = append(q.items, item)
	q.sz++
	q.notempty.Signal()
	fmt.Println("Enqueued", item)
	return 0
}

func (q *queue) Dequeue() (string, int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for q.sz == 0 {
		fmt.Println("Waiting dequeue thread")
		q.notempty.Wait()
	}
	if q.sz == 0 {
		fmt.Println("queue is empty, cannot dequeue")
		return "", 1
	}
	item := q.items[0]
	q.items = q.items[1:]
	q.sz--
	q.notfull.Signal()
	fmt.Println("Dequeued", item)
	return item, 0
}

func (q *queue) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.sz
}

func (q *queue) Capacity() int {
	return q.cap
}

func main() {
	var wait sync.WaitGroup
	q := queue{}
	q.Init(3)
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func(val string) {
			defer wait.Done()
			q.Enqueue(val)
		}(fmt.Sprintf("item%d", i))
	}
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			_, err := q.Dequeue()
			if err != 0 {
				fmt.Println("Failed to dequeue")
			}
		}()
	}
	wait.Wait()
}
