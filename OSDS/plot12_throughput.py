import matplotlib.pyplot as plt
import numpy as np
import sys
import os

a = [1, 2, 5, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100]
b = 1000
dList = []
experimentTimes = int(sys.argv[1])

for i in a:
    File = open("./data/RESIZE_THROUGHPUT_TN" + str(i) + "QL" + str(b) + ".txt", "r")
    lines_list = File.readlines()
    d = float(lines_list[0])
    for j in range(0, experimentTimes):
        d += float(lines_list[j])
    d /= experimentTimes
    dList.append(d)


plt.plot(a, dList)
plt.title("Throughput vs # of Threads (Queue Length = 1000)")
plt.xlabel("# of Threads")
plt.ylabel("Throughput (per second)")
plt.savefig("./fig/resize_throughput_threadnum.png")
plt.show()

plt.clf()

dList = []

a = [10, 15, 20, 25, 50, 75, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000]
b = 70

for i in a:
    File = open("./data/RESIZE_THROUGHPUT_TN" + str(b) + "QL" + str(i) + ".txt", "r")
    lines_list = File.readlines()
    d = float(lines_list[0])
    for j in range(0, experimentTimes):
        d += float(lines_list[j])
    d /= experimentTimes
    dList.append(d)


plt.plot(a, dList)
plt.title("Throughput vs Queue Length (70 Threads)")
plt.xlabel("Queue Length")
plt.ylabel("Throughput (per second)")
plt.savefig("./fig/resize_throughput_queuelength.png")
plt.show()