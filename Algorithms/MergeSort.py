#!/usr/bin/env python 

import time
import numpy as np
import random
# import sort func we previous use
from SelectionSort import selectionSort

'''
    1. interate over two sorted list
    2. find the minimum element for each comparison and moving to the new list
    3. finish moving rest of elements to the new list 
    4. return 
'''

def mergeSort(seq1, seq2):
    sorted_list = []
    index_a = 0
    index_b = 0

    while index_a < len(seq1) and index_b < len(seq2):
        if seq1[index_a] < seq2[index_b]:
            sorted_list.append(seq1[index_a])
            index_a += 1
        else:
            sorted_list.append(seq2[index_b])
            index_b += 1

    while index_a < len(seq1):
        sorted_list.append(seq1[index_a])
        index_a +=1

    while index_b < len(seq2):
        sorted_list.append(seq2[index_b])
        index_b +=1

    return sorted_list 

if __name__ == '__main__':
    #randlist1 = np.random.randint(1,100, (10))
    #randlist2 = np.random.randint(1,100, (10))
    randlist1 = [ random.randint(1, 100) for i in range(10) ]
    randlist2 = [ random.randint(1, 100) for i in range(10) ]
    print(type(randlist1), type(randlist2))

    sorted_list1, _ = selectionSort(randlist1)
    sorted_list2, _ = selectionSort(randlist2)

    print(sorted_list1, sorted_list2)
    mergedSortList = mergeSort(sorted_list1, sorted_list2)
    print("MergedList", mergedSortList)


