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
func (q *queue) resize() (string, string, int, time.Duration) {
	q.lock.Lock()
	defer q.lock.Unlock()
	start := time.Now()
	for q.sz == 0 {
		fmt.Println("Waiting dequeue thread")
		q.notempty.Wait()
	}
	filename := q.items[0]
	q.items = q.items[1:]
	q.sz--
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
	store_filename := filename - ".JPEG" + "_resized_.JPEG"
	out, err := os.Create(store_filename)
	if err != nil {
		fmt.Println("Write picture failed", store_filename)
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, m, nil)
	elapsed := time.Since(start)
	fmt.Println("Dequeued", filename)
	q.notfull.Signal()
	return filename, store_filename, 0, elapsed
}
func (q *queue) multi_resize(len int)([]string,[]string, []int,[]time.Duration,[]time.Duration){
	var filenames []string
	var store_filenames []string
	var errs []int
	var op_times []time.Duration
	var latency []time.Duration
	for i:=0;i<len;i++ {
		start = time.Now()
		filename, store_filename, err, op_time := q.resize()
		latency := time.Since(start)
		filenames = append(filenames, filename)
		store_filenames = append(store_filenames, store_filename)
		errs = append(errs, err)
		op_times = append(op_times, op_time)
		latency = append(latency, latency)
	}
	return filenames, store_filenames, errs, op_times, latency
}
type log{
	filename string
	store_filename string
	err int
	op_time time.Duration
	latency time.Duration
}
func main(){
	q:=queue{}
	q.Init(100)
	for i:=0;i<100;i++ {
		q.Enqueue(i)
	}
	log:=log{}
	for i:=0;i<10;i++{
		filenames, store_filenames, errs, op_times, latency := q.multi_resize(100)
		for i:=0;i<10;i++{
			log.append(log{filenames[i], store_filenames[i], errs[i], op_times[i], latency[i]})
		}
	}
}