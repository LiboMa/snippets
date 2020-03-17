#!/usr/bin/env python

import random

rand_seq = [ random.randint(1,100) for _ in range(10) ]

#n = len(rand_seq)
#for i in range(n - 1):
#    print("OUTER LOOP: %d th" % i)
#    for j in range(i+ n - 1):
#        if rand_seq[j] > rand_seq[j+1]:
#            tmp = rand_seq[j] 
#            rand_seq[j] = rand_seq[j+1]
#            rand_seq[j+1] = tmp
