import matplotlib.pyplot as plt
import numpy as np
import sys

a = [1, 2, 5, 10, 20, 50, 100, 200, 300]
b = 300
dqList = []
enList = []

for i in a:
    File = open("./data/THROUGHPUT_TN" + str(b) + "QL" + str(i) + ".txt", "r")
    lines_list = File.readlines()
    d, e = (int(val) for val in lines_list[0].split())
    dqList.append(d)
    enList.append(e)

plt.plot(a, dqList)
plt.plot(a, enList)
plt.title("Throughput vs Queue Length (10s, 300 threads)")
plt.xlabel("Queue Length")
plt.ylabel("Throughput")
plt.savefig("./fig/throughput_queuelength.png")
plt.show()

plt.clf()

dqList = []
enList = []

for i in a:
    File = open("./data/THROUGHPUT_TN" + str(i) + "QL" + str(i) + ".txt", "r")
    lines_list = File.readlines()
    d, e = (int(val) for val in lines_list[0].split())
    dqList.append(d)
    enList.append(e)

plt.plot(a, dqList)
plt.plot(a, enList)
plt.title("Throughput vs # of Threads (10s)")
plt.xlabel("# of Threads")
plt.ylabel("Throughput")
plt.savefig("./fig/throughput_threadnum.png")
plt.show()