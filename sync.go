package main

import(
	"fmt"
	"sync"
	"time"
	"os"
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

func (q *queue) Enqueue(item string) (int, time.Duration) {
	q.lock.Lock()
	defer q.lock.Unlock()
	start := time.Now()
	for q.sz == q.cap {
		fmt.Println("Waiting enqueue thread", item)
		q.notfull.Wait()
	}
	// if q.sz == q.cap {
	// 	fmt.Println("queue is full, cannot enqueue", item)
	// 	return 1
	// }
	q.items = append(q.items, item)
	q.sz++
	elapsed := time.Since(start)
	q.notempty.Signal()
	//fmt.Println("Enqueued", item)
	return 0, elapsed
}

func (q *queue) Dequeue() (string, int, time.Duration) {
	q.lock.Lock()
	defer q.lock.Unlock()
	start := time.Now()
	for q.sz == 0 {
		fmt.Println("Waiting dequeue thread")
		q.notempty.Wait()
	}
	// if q.sz == 0 {
	// 	fmt.Println("queue is empty, cannot dequeue")
	// 	return "", 1
	// }
	item := q.items[0]
	q.items = q.items[1:]
	q.sz--
	elapsed := time.Since(start)
	q.notfull.Signal()
	//fmt.Println("Dequeued", item)
	return item, 0, elapsed
}

func (q *queue) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.sz
}

func (q *queue) Capacity() int {
	return q.cap
}
func synccall(queue_len int,thread_num int, filename string) {
	var wait sync.WaitGroup
	Total_thread_num := thread_num
	q := queue{}
	q.Init(queue_len)
	FILENAME := filename
	fmt.Printf("queue length is %d, thread number is %d, result is stored in %s\n", queue_len, thread_num, filename)
	file, _:= os.OpenFile(FILENAME+".txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	fmt.Fprintf(file, "queue length is %d, thread number is %d, result is stored in %s\n", queue_len, thread_num, filename)
    file.Close()
	for i := 0; i < Total_thread_num; i++ {
		wait.Add(1)
		go func(val string) {
			defer wait.Done()
			start := time.Now()
			file, _ := os.OpenFile(FILENAME+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			defer file.Close()
			_ ,optime := q.Enqueue(val)
			latency := time.Since(start)
			output := fmt.Sprintf("Enqueue %s took operation %s and latency %s\n", val, optime, latency)
			if _, err := file.WriteString(output); err != nil {
				fmt.Println("Failed to write to file:", err)
			}
			enqueue_opfile, _:= os.OpenFile(FILENAME+"en_op.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			enqueue_optime := fmt.Sprintf("%s ",optime)
			enqueue_opfile.WriteString(enqueue_optime)
			enqueue_opfile.Close()
			enqueue_latfile, _:= os.OpenFile(FILENAME+"en_la.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			enqueue_latency := fmt.Sprintf("%s ",latency)
			enqueue_latfile.WriteString(enqueue_latency)
			enqueue_latfile.Close()
		}(fmt.Sprintf("item%d", i))
	}
	for i := 0; i < Total_thread_num; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			start := time.Now()
			file, _ := os.OpenFile(FILENAME+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			defer file.Close()
			_, _, optime := q.Dequeue()
			latency := time.Since(start)
			output := fmt.Sprintf("Dequeue took operation %s and latency %s\n", optime, latency)
			if _, err := file.WriteString(output); err != nil {
				fmt.Println("Failed to write to file:", err)
			}
			dequeue_opfile, _:= os.OpenFile(FILENAME+"de_op.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			dequeue_optime := fmt.Sprintf("%s ",optime)
			dequeue_opfile.WriteString(dequeue_optime)
			dequeue_opfile.Close()
			dequeue_latfile, _:= os.OpenFile(FILENAME+"de_la.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			dequeue_latency := fmt.Sprintf("%s ",latency)
			dequeue_latfile.WriteString(dequeue_latency)
			dequeue_latfile.Close()
		}()
	}
	wait.Wait()
}
func main() {
	THREAD_NUM := 1000
	QUEUE_LEN := 50
	synccall(QUEUE_LEN, THREAD_NUM, fmt.Sprintf("./data/TN%dQL%d", THREAD_NUM, QUEUE_LEN))
}
