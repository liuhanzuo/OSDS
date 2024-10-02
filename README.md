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

## Tools

I implement some (not) useful tools in makefile for cleaning, you could run

```makefile
make cleantxt
```

to clean all txt files in ``./data`` ^V^

## Task1.2

Here I implement a file called ``multi_resize.go`` in ``.`` directory

Since copy the data again is such a complicated task, I use the images in ``./gitclone/tiny-imagenet-200`` straightly

as a result, you can view the result in ``./data/log%d.txt`` where ``%d`` stand for the thread number

Please implement the report, and make sure the data you get is correct.

Total operation time stand for operation of each thread in total

Total latency stand for the latency of each thread in total

Total elapsed time stand for the total time for both operating and waiting for the lock, **which is the time we were optimizing!**
