import matplotlib.pyplot as plt
import numpy as np
import sys
import os

t = int(sys.argv[1])
for i in range(0, t):
    print(f"Running {i}-th experiments")
    os.system("bash run12_throughput.sh")