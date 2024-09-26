# OSDS

This is the repository for OS&amp;DS courses in 2024 autumn semester

## Part1--queue

I impllement the queue in ./queue.go with the needed functions.

In sync.go thread, I implement concurrently working dequeue and enqueue threads

Update1: I count time for each operation, especially for operation time and latency.

The place to store it is in "./data/", and the file name stand for TN -- thread number. QL -- queue length

The later ones stand for enqueue_operation time and latency time/ dequeue operation/ latency time, you can see it in ./data

For makefile, you can run 

```makefile
make syncall
```

to do the concurrency enqueue&dequeue(maybe for simplification need to update to make syncall 10 100)

TODO: plot the figure   write the report    write a makefile for auto generation


