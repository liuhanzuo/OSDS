package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/nfnt/resize"
)

type Queue struct {
	c chan string
}

func (q *Queue) Init(capacity int) {
	q.c = make(chan string, capacity)
}
func (q *Queue) Enqueue(item string) int {
	q.c <- item
	return 0
}
func (q *Queue) Dequeue() (string, int) {
	return <-q.c, 0
}
func (q *Queue) size() int {
	return len(q.c)
}
func (q *Queue) capacity() int {
	return cap(q.c)
}

var wg sync.WaitGroup
var lock sync.Mutex
var q Queue
var queue_capacity int
var thread_num int
var all int
var latency []float64

func func2(i int) {
	for {
		id, error := q.Dequeue()
		if error != 0 {
			return
		}
		st := time.Now()
		filename := "./gitclone/tiny-imagenet-200/test/images/test_" + id + ".JPEG"
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
		filename = "./gitclone/tiny-imagenet-200/test/images/resize_demo_resized_" + id + ".JPEG"
		out, err := os.Create(filename)
		if err != nil {
			fmt.Println("Write picture failed", filename)
			log.Fatal(err)
		}
		defer out.Close()
		jpeg.Encode(out, m, nil)
		lock.Lock()
		tmp, _ := strconv.Atoi(id)
		if tmp%100 == 0 {
			fmt.Fprintf(os.Stderr, "DONE %s\n", id)
		}
		latency = append(latency, time.Since(st).Seconds())
		lock.Unlock()
		wg.Done()
	}
}
func run() {
	for i := 0; i < thread_num; i++ {
		go func2(i)
	}
	for i := 0; i < all; i++ {
		name := fmt.Sprintf("%d", i)
		q.Enqueue(name)
		wg.Add(1)
	}
	wg.Wait()
}
func main() {
	thread_num, _ = strconv.Atoi(os.Args[1])
	queue_capacity, _ = strconv.Atoi(os.Args[2])
	mode := os.Args[3]
	all = 1000
	start := time.Now()
	q.Init(queue_capacity)
	run()
	throughput := float64(all) / time.Since(start).Seconds()
	// latency_sum:=float64(0)
	// for i := 0; i < all; i++ {
	// 	latency_sum += latency[i]
	// }
	// time_in_second := time.Since(start).Seconds()
	// fmt.Println(mode)
	if mode == "THREADPUT" {
		FILENAME := fmt.Sprintf("./data/RESIZE_THROUGHPUT_TN%dQL%d", thread_num, queue_capacity)
		file, _ := os.OpenFile(FILENAME+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		fmt.Fprintf(file, "%f\n", throughput)
	} else if mode == "LATENCY" {
		fmt.Println(all)
		for i := 0; i < all; i++ {
			fmt.Println(latency[i])
		}
	}
	// fmt.Println("time multi-thread ", time_in_second)
	// fmt.Println("latency multi-thread ", latency_sum)
	// fmt.Println("throughput multi-thread ", throughput)
}
