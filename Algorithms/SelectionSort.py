#!/usr/bin/env python 


# find the smallest one and add to the new list

seq = [2,3,512,23,12,3]


sorted_seq = []


def selectionSort(seq):
    n = len(seq)
    for i in range(n - 1):
        for j in range(i+1, n):
            print(j)

selectionSort(seq)

    

