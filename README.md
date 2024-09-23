# OSDS

This is the repository for OS&amp;DS courses in 2024 autumn semester

## Part1--queue

I impllement the queue in ./queue.go with the needed functions.

In sync.go thread, I implement concurrently working dequeue and enqueue threads

However, since I do not know what to do when the queue is full/empty(whether dequeue threads should immediately return or should wait until next enqueue threads to put some data in?)

note that the enviroment condition is for the later condition thus if TA tell us we should use the former one just delete the code with enviroment condition
