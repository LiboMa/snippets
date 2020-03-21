#!/usr/bin/env python

import random
import time

rand_seq = [ random.randint(1,10000) for _ in range(100000) ]

def bubbleSort(seq):
    now = time.time()
    n = len(seq)
    for i in range(n):
        # print("OUTER : ->", i)
        for j in range( n - i - 1 ):
            # check j th of next one who is bigger
            if rand_seq[j] > rand_seq[j+1]:
                #print("swap")
                swap = rand_seq[j]
                rand_seq[j] = rand_seq[j+1]
                rand_seq[j+1] = swap
                
    end = time.time() - now
    return seq, end

if __name__ == '__main__':
    l, t = bubbleSort(rand_seq)
    
    print("{}, {} items sorted, time cost: {:.7f}".format(l, len(l), t))
