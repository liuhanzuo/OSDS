import matplotlib.pyplot as plt
import numpy as np
import sys

if len(sys.argv) < 4:
    print('Usage: plot_distribution.py PLOT_TITLE X_LABEL Y_LABEL [SAVING_LOCATION]')
    sys.exit(1)
n = int(input())
a = []
for i in range(0, n):
    a.append(float(input()))
    if(a[-1] > 1000) :
       a.pop() 
plt.hist(a, bins = 100)
plt.title(sys.argv[1])
plt.xlabel(sys.argv[2])
plt.ylabel(sys.argv[3])
plt.savefig(sys.argv[4] if len(sys.argv) > 4 else './' + sys.argv[1] + '.png')
plt.show()