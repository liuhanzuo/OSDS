package main

import (
	"os"
	"fmt"
	"log"
	"image/jpeg"
	"github.com/nfnt/resize"
	"time"
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

func (q *queue) Enqueue(item string) (int, time.Duration) {
	q.lock.Lock()
	defer q.lock.Unlock()
	start := time.Now()
	for q.sz == q.cap {
//		fmt.Println("Waiting enqueue thread", item)
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
//		fmt.Println("Waiting dequeue thread")
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
func rsize(id string){
	filename := "tiny-imagenet-200//test//images//test_"+id+".JPEG";
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Load picture failed", filename)
		log.Fatal(err)
	}
	img, err := jpeg.Decode(file)
	m := resize.Resize(128, 128, img, resize.Lanczos3)
	if err != nil {
		fmt.Println("Decode picture failed", filename)
		log.Fatal(err)
	}
	file.Close()
	filename = "tiny-imagenet-200//test//images//resize_demo_resized_"+id+".JPEG";
	out, err := os.Create(filename)
	if err != nil {
		fmt.Println("Write picture failed", filename)
		log.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)
}
var wait sync.WaitGroup
var q queue;
func func1(i int){
	for j:=i*100;j<i*100+100;j++{
		str:=fmt.Sprintf("%d",j)
		q.Enqueue(str)
	}
	wait.Done()
}
func func2(i int){
	for j:=0;j<100;j++{
		str,_,_:=q.Dequeue()
		rsize(str)
	}
	wait.Done()
}
func main(){
	start:=time.Now()
	q.Init(100)
	for i:=0;i<100;i++{
		wait.Add(1)
		go func1(i)
		wait.Add(1)
		go func2(i)
	}
	wait.Wait()
	fmt.Println("time used multi-thread %d",time.Since(start))
}
