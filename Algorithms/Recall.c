#include <stdio.h>
#include <time.h>
#include <stdlib.h>

//#include "timestamp.h"

#define NUMBERS 100000

int bubbleSort(int seq[], int n);
int sum(int seq[], int n);
int sump(int *seq, int *end);
int bubblePointerSort(int *seq, int *end);
int Random(int min, int max);
int selectionSort(int seq[], int n);
//void timestamp ( );
//double cpu_time ( );
//
int Random(int min, int max) {
    int number = min + rand() % (max - min);
    return number;
}

int sump(int *start, int *end) {

    int total = 0;

    while(start < end) {
        total += *start;
        start++;
    }
    return total;

}

int bubbleSort(int seq[], int n) {

    int i;
    int key = 0;
    //timestamp();
     
    for(i=0;i<n-1;i++) {
        //Outer loop
        for(int j = 0;j<(n-i-1);j++) {
            //printf("inner loop idx => %d, j => %d \n", j, seq[j]);
            if(seq[j]>seq[j+1]) {
                // swap j and j+1
                int tmp = seq[j];
                seq[j] = seq[j+1];
                seq[j+1] = tmp;
            }
        }
    }

    return 0;

}


int selectionSort(int seq[], int n) {

    // get timer
    clock_t t;
    t = clock();

    int minIndex;
    for(int i=0;i<n;i++) {
        //Outer loop
        minIndex = i;
        for(int j = i+1;j<n;j++) {
            //printf("inner loop idx => %d, j => %d \n", j, seq[j]);
            if(seq[j] < seq[minIndex]) {
                // swap j and j+1
                minIndex = j;
            }
        }

        if(minIndex != i) {
            int tmp = seq[i];
            seq[i] = seq[minIndex];
            seq[minIndex] = tmp;
        }
    }

   t = clock() - t;
   double time_taken = ((double)t)/CLOCKS_PER_SEC;
   printf("The program took %f seconds to execute\n", time_taken);

    return 0;

}

int bubblePointerSort(int *seq, int *end) {

    int total=0;
    int tmp;

    while(seq < end) {
        //total += *seq;
        seq++;
        while(seq < end) {
            printf("inner loop value: %d\n", seq[0]);
            seq++;
        }
    }
    return total;

}

int main(int argc, char *argv[]) {

    int days[NUMBERS];
    srand(time(NULL));   // Initialization, should only be called once.
    //int r = rand();

    for (int index=0;index<NUMBERS; index++) {
        //printf("%d\t", Random(1,100));
        //days[index] = rand() %100; 
        days[index] = Random(1,NUMBERS); 
        //printf("i %d => number %d\n", index, days[index]);
    }
    //sum the number 
    //
    clock_t t;
    t = clock();

    int total;
    //total = bubblePointerSort(days, days+NUMBERS);
    bubbleSort(days, NUMBERS);

    //selectionSort(days, NUMBERS);

    for(int z=0;z<NUMBERS;z++) {
        printf("%d ",days[z]);
        total +=1;
    }

    printf("total => %d\n", total);
    t = clock() - t;
    double time_taken = ((double)t)/CLOCKS_PER_SEC;
    printf("The program took %f seconds to execute\n", time_taken);

   // printf("\n ");
    //printf("using pointer func\n");
    //total = sump(days, days+NUMBERS);
    //printf("total => %d", total);
    //printf("%2d", r);

}
