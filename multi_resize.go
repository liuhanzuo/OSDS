package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"sync"
	"time"

	"github.com/nfnt/resize"
)

type queue struct {
	items    []string
	cap      int
	sz       int
	lock     sync.Mutex
	notempty sync.Cond
	notfull  sync.Cond
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
	fmt.Println("enqueue starting", item)
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

func (q *queue) resize() (string, string, int, time.Duration) {
	//	fmt.Println("resize starting")
	start := time.Now()
	filename, _, _ := q.Dequeue()
	fmt.Println(filename)
	file, err := os.Open(filename)
	//fmt.Println("checkpoint2")
	if err != nil {
		fmt.Println("Load picture failed", filename)
		log.Fatal(err)
	}
	img, err := jpeg.Decode(file)
	//fmt.Println("checkpoint3")
	m := resize.Resize(128, 128, img, resize.Lanczos3)
	//fmt.Println("checkpoint4")
	if err != nil {
		fmt.Println("Decode picture failed", filename)
		log.Fatal(err)
	}
	//fmt.Println("checkpoint5")
	file.Close()
	store_filename := filename[0:len(filename)-5] + "_resized_.JPEG"
	out, err := os.Create(store_filename)
	if err != nil {
		fmt.Println("Write picture failed", store_filename)
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)
	elapsed := time.Since(start)
	//fmt.Println("Dequeued", filename)
	q.notfull.Signal()
	return filename, store_filename, 0, elapsed
}
func (q *queue) multi_resize(len int) ([]string, []string, []int, []time.Duration, []time.Duration) {
	var filenames []string
	var store_filenames []string
	var errs []int
	var op_times []time.Duration
	var latencys []time.Duration
	for i := 0; i < len; i++ {
		start := time.Now()
		filename, store_filename, err, op_time := q.resize()
		latency := time.Since(start)
		filenames = append(filenames, filename)
		store_filenames = append(store_filenames, store_filename)
		errs = append(errs, err)
		op_times = append(op_times, op_time)
		latencys = append(latencys, latency)
	}
	return filenames, store_filenames, errs, op_times, latencys
}

type logg struct {
	filename       []string
	store_filename []string
	err            []int
	op_time        []time.Duration
	latency        []time.Duration
}

func main() {
	q := queue{}
	TOTAL_NUM := 1000
	q.Init(TOTAL_NUM)
	THREAD_NUM := 20
	for i := 0; i < TOTAL_NUM; i++ {
		q.Enqueue(fmt.Sprintf("./gitclone/tiny-imagenet-200/test/images/test_%d.JPEG", i))
	}
	log := logg{}
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < THREAD_NUM; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			filenames, store_filenames, errs, op_times, latencys := q.multi_resize(TOTAL_NUM / THREAD_NUM)
			for j := range filenames {
				log.filename = append(log.filename, filenames[j])
				log.store_filename = append(log.store_filename, store_filenames[j])
				log.err = append(log.err, errs[j])
				log.op_time = append(log.op_time, op_times[j])
				log.latency = append(log.latency, latencys[j])
			}
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	logFile, _ := os.Create(fmt.Sprintf("./data/log%d.txt", THREAD_NUM))
	defer logFile.Close()
	total_latency := time.Duration(0)
	total_optime := time.Duration(0)
	for i := range log.filename {
		logEntry := fmt.Sprintf("Filename: %s, Store Filename: %s, Error: %d, Operation Time: %v, Latency: %v\n",
			log.filename[i], log.store_filename[i], log.err[i], log.op_time[i], log.latency[i])
		logFile.WriteString(logEntry)
		total_latency += log.latency[i]
		total_optime += log.op_time[i]
	}
	fmt.Println("Total Elapsed Time: %v\n", elapsed)
	logFile.WriteString(fmt.Sprintf("Total Operation Time: %v\n", total_optime))
	logFile.WriteString(fmt.Sprintf("Total Latency: %v\n", total_latency))
	logFile.WriteString(fmt.Sprintf("Total Elapsed Time: %v\n", elapsed))
}
