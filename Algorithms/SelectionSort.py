#!/usr/bin/env python 
# find the smallest one and add to the new list
import time
import random

#seq = [2,3,51,23,12,5,100]
#sorted_seq = []
rand_seq = [ random.randint(1,10000) for _ in range(10000) ]

def selectionSort(seq):
    start = time.time()
    n = len(seq) - 1
    for i in range(n):
        minIndex = i
        #print(seq[minIndex])
        for j in range(i+1, n):
            if seq[j] < seq[minIndex]:
                minIndex = j
                #print(seq[minIndex])

        if minIndex != i:
            swap = seq[i]
            seq[i] = seq[minIndex]
            seq[minIndex] = swap
            #print("round: {}, min {}, minValue {} <=SWAP...==> i {}, iValue {}".format(i, minIndex,seq[minIndex], i,seq[i]))
        else:
            #print("round: {}, min {}, minValue {} , i {}, iValue {}".format(i, minIndex,seq[minIndex], i,seq[i]))
            pass
    elasped = time.time() - start
    return seq, elasped 

if __name__ == '__main__':
    seq, elasped = selectionSort(rand_seq)

    print("{} items sorted, time cost:{} s".format(len(seq), elasped))
