package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type queue struct {
	items    []string
	cap      int
	sz       int
	lock     sync.Mutex
	notempty sync.Cond
	notfull  sync.Cond

	closed bool
}

func (q *queue) Init(capacity int) {
	q.items = make([]string, 0, capacity)
	q.cap = capacity
	q.sz = 0
	q.notempty.L = &q.lock
	q.notfull.L = &q.lock
	q.closed = false
}

func (q *queue) Enqueue(item string) int /*, time.Duration*/ {
	q.lock.Lock()
	defer q.lock.Unlock()
	// start := time.Now()
	for q.sz == q.cap {
		// fmt.Println("Waiting enqueue thread", item)
		if q.closed {
			return 1
		}
		q.notfull.Wait()
	}
	// if q.sz == q.cap {
	// 	fmt.Println("queue is full, cannot enqueue", item)
	// 	return 1
	// }
	q.items = append(q.items, item)
	q.sz++
	// elapsed := time.Since(start)
	q.notempty.Signal()
	//fmt.Println("Enqueued", item)
	return 0 /*, elapsed*/
}

func (q *queue) Dequeue() (string, int /*, time.Duration*/) {
	q.lock.Lock()
	defer q.lock.Unlock()
	// start := time.Now()
	for q.sz == 0 {
		// fmt.Println("Waiting dequeue thread")
		if q.closed {
			return "", 1
		}
		q.notempty.Wait()
	}
	// if q.sz == 0 {
	// 	fmt.Println("queue is empty, cannot dequeue")
	// 	return "", 1
	// }
	item := q.items[0]
	q.items = q.items[1:]
	q.sz--
	// elapsed := time.Since(start)
	q.notfull.Signal()
	// fmt.Println("Dequeued", item)
	return item, 0 /*, elapsed*/
}

func (q *queue) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.sz
}

func (q *queue) Capacity() int {
	return q.cap
}

func printChannel[T any](printList chan T, filename string) {
	file, _ := os.OpenFile(filename+".txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	fmt.Fprintf(file, "%d\n", len(printList))
	for elem := range printList {
		fmt.Fprintln(file, elem)
	}
}

var dequeTimeList [100000]int64
var enqueTimeList [100000]int64

func latencyTest(queue_len int, thread_num int, filename string) {
	var wait sync.WaitGroup
	Total_thread_num := thread_num
	q := queue{}
	q.Init(queue_len)
	// FILENAME := filename
	// fmt.Printf("queue length is %d, thread number is %d, result is stored in %s\n", queue_len, thread_num, filename)
	// file, _ := os.OpenFile(FILENAME+".txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	// fmt.Fprintf(file, "queue length is %d, thread number is %d, result is stored in %s\n", queue_len, thread_num, filename)
	// file.Close()
	wait.Add(2 * Total_thread_num)
	/*for i := 0; i < 2*Total_thread_num; i++ {
		// wait.Add(1)
		if (i & 1) == 0 {
			go func(val string, id int) {
				defer wait.Done()
				start := time.Now()
				_ = q.Enqueue(val)
				latency := time.Since(start)
				enqueTimeList[id] = latency.Nanoseconds()
				// file, _ := os.OpenFile(FILENAME+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				// defer file.Close()
				// output := fmt.Sprintf("Enqueue %s took operation %s and latency %s\n", val, optime, latency)
				// if _, err := file.WriteString(output); err != nil {
				//	fmt.Println("Failed to write to file:", err)
				//}
				// enqueue_opfile, _ := os.OpenFile(FILENAME+"en_op.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				// enqueue_optime := fmt.Sprintf("%s\n", optime)
				// enqueue_opfile.WriteString(enqueue_optime)
				// enqueue_opfile.Close()
				// enqueue_latfile, _ := os.OpenFile(FILENAME+"en_la.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				// enqueue_latency := fmt.Sprintf("%s\n", latency)
				// enqueue_latfile.WriteString(enqueue_latency)
				// enqueue_latfile.Close()
			}(fmt.Sprintf("item%d", i), i)
		} else {
			// wait.Add(1)
			go func(id int) {
				defer wait.Done()
				start := time.Now()

				_, _ = q.Dequeue()
				latency := time.Since(start)
				dequeTimeList[id] = latency.Nanoseconds()
				// file, _ := os.OpenFile(FILENAME+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				// defer file.Close()
				// output := fmt.Sprintf("Dequeue took operation %s and latency %s\n", optime, latency)
				// if _, err := file.WriteString(output); err != nil {
				// 	fmt.Println("Failed to write to file:", err)
				// }

				// dequeue_opfile, _ := os.OpenFile(FILENAME+"de_op.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				// dequeue_optime := fmt.Sprintf("%s\n", optime)
				// dequeue_opfile.WriteString(dequeue_optime)
				// dequeue_opfile.Close()
				// dequeue_latfile, _ := os.OpenFile(FILENAME+"de_la.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				// dequeue_latency := fmt.Sprintf("%s\n", latency)
				// dequeue_latfile.WriteString(dequeue_latency)
				// dequeue_latfile.Close()
			}(i)

		}
	}*/

	for i := 0; i < Total_thread_num; i++ {
		// wait.Add(1)
		go func(val string, id int) {
			defer wait.Done()
			start := time.Now()
			_ = q.Enqueue(val)
			latency := time.Since(start)
			enqueTimeList[id] = latency.Nanoseconds()
			// file, _ := os.OpenFile(FILENAME+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			// defer file.Close()
			// output := fmt.Sprintf("Enqueue %s took operation %s and latency %s\n", val, optime, latency)
			// if _, err := file.WriteString(output); err != nil {
			//	fmt.Println("Failed to write to file:", err)
			//}
			// enqueue_opfile, _ := os.OpenFile(FILENAME+"en_op.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			// enqueue_optime := fmt.Sprintf("%s\n", optime)
			// enqueue_opfile.WriteString(enqueue_optime)
			// enqueue_opfile.Close()
			// enqueue_latfile, _ := os.OpenFile(FILENAME+"en_la.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			// enqueue_latency := fmt.Sprintf("%s\n", latency)
			// enqueue_latfile.WriteString(enqueue_latency)
			// enqueue_latfile.Close()
		}(fmt.Sprintf("item%d", i), i)
	}
	for i := 0; i < Total_thread_num; i++ {
		// wait.Add(1)
		go func(id int) {
			defer wait.Done()
			start := time.Now()

			_, _ = q.Dequeue()
			latency := time.Since(start)
			dequeTimeList[id] = latency.Nanoseconds()
			// file, _ := os.OpenFile(FILENAME+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			// defer file.Close()
			// output := fmt.Sprintf("Dequeue took operation %s and latency %s\n", optime, latency)
			// if _, err := file.WriteString(output); err != nil {
			// 	fmt.Println("Failed to write to file:", err)
			// }

			// dequeue_opfile, _ := os.OpenFile(FILENAME+"de_op.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			// dequeue_optime := fmt.Sprintf("%s\n", optime)
			// dequeue_opfile.WriteString(dequeue_optime)
			// dequeue_opfile.Close()
			// dequeue_latfile, _ := os.OpenFile(FILENAME+"de_la.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			// dequeue_latency := fmt.Sprintf("%s\n", latency)
			// dequeue_latfile.WriteString(dequeue_latency)
			// dequeue_latfile.Close()
		}(i)
	}
	wait.Wait()
	file, _ := os.OpenFile(filename+".txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	fmt.Fprintln(file, 2*thread_num)
	for i := 0; i < thread_num; i++ {
		fmt.Fprintln(file, dequeTimeList[i])
	}
	for i := 0; i < thread_num; i++ {
		fmt.Fprintln(file, enqueTimeList[i])
	}
	file.Close()
}

func throughputTest(queue_len int, thread_num int, timeLimit float64, filename string) {
	fmt.Printf("OK")
	var wait sync.WaitGroup
	Total_thread_num := thread_num
	q := queue{}
	q.Init(queue_len)
	FILENAME := filename
	fmt.Printf("queue length is %d, thread number is %d, result is stored in %s\n", queue_len, thread_num, filename)
	totEnqueueCount := 0
	totDequeueCount := 0
	var enqueueCountLock sync.Mutex
	var dequeueCountLock sync.Mutex
	start := time.Now()
	for i := 0; i < Total_thread_num; i++ {
		wait.Add(1)
		go func(val string) {
			enqueueCount := 0
			defer wait.Done()
			// file, _ := os.OpenFile(FILENAME+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			// defer file.Close()
			for {
				_ = q.Enqueue(val)
				endTime := time.Since(start).Seconds()
				if endTime < float64(timeLimit) {
					enqueueCount++
				} else {
					break
				}
			}
			enqueueCountLock.Lock()
			defer enqueueCountLock.Unlock()
			totEnqueueCount += enqueueCount
		}("")
	}
	for i := 0; i < Total_thread_num; i++ {
		wait.Add(1)
		go func(val string) {
			dequeueCount := 0
			defer wait.Done()
			// file, _ := os.OpenFile(FILENAME+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			// defer file.Close()

			for {
				_, _ = q.Dequeue()
				endTime := time.Since(start).Seconds()
				if endTime < timeLimit {
					dequeueCount++
				} else {
					break
				}
			}
			dequeueCountLock.Lock()
			defer dequeueCountLock.Unlock()
			totDequeueCount += dequeueCount
		}("")
	}
	time.Sleep(time.Duration(timeLimit) * time.Second)
	q.closed = true
	q.notempty.Broadcast()
	q.notfull.Broadcast()
	wait.Wait()

	file, _ := os.OpenFile(FILENAME+".txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	fmt.Fprintf(file, "%d %d", totDequeueCount, totEnqueueCount)
}

func main() {
	THREAD_NUM, _ := strconv.Atoi(os.Args[2])
	QUEUE_LEN, _ := strconv.Atoi(os.Args[3])
	fmt.Printf("%s", os.Args[1])
	if os.Args[1] == "LATENCYTEST" {
		latencyTest(QUEUE_LEN, THREAD_NUM, fmt.Sprintf("./data/LATENCY_TN%dQL%d", THREAD_NUM, QUEUE_LEN))
	} else if os.Args[1] == "THROUGHPUTTEST" {
		fmt.Printf("IN")
		TIME_LIMIT, _ := strconv.ParseFloat(os.Args[4], 64)
		throughputTest(QUEUE_LEN, THREAD_NUM, TIME_LIMIT, fmt.Sprintf("./data/THROUGHPUT_TN%dQL%d", THREAD_NUM, QUEUE_LEN))
	} else {
		fmt.Printf("NOT")
	}
}
