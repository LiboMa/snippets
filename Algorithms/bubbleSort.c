// C program for implementation of Bubble sort 
#include <stdio.h>
#include <time.h>
#include <stdlib.h>
  
#define NUMBERS 100000

int Random(int min, int max) {
    int number = min + rand() % (max - min);
    return number;
}

void swap(int *xp, int *yp) 
{ 
    int temp = *xp; 
    *xp = *yp; 
    *yp = temp; 
} 
  
// A function to implement bubble sort 
void bubbleSort(int arr[], int n) 
{ 
   int i, j; 
   for (i = 0; i < n-1; i++)       
  
       // Last i elements are already in place    
       for (j = 0; j < n-i-1; j++)  
           if (arr[j] > arr[j+1]) 
              swap(&arr[j], &arr[j+1]); 
} 
  
/* Function to print an array */
void printArray(int arr[], int size) 
{ 
    int i; 
    for (i=0; i < size; i++) 
        printf("%d ", arr[i]); 
    printf("\n"); 
} 
  
// Driver program to test above functions 
int main() 
{ 
    int arr[NUMBERS];
    srand(time(NULL));   // Initialization, should only be called once.
    //int r = rand();
    //
    clock_t t;
    t = clock();

    for (int index=0;index<NUMBERS; index++) {
        //printf("%d\t", Random(1,100));
        //days[index] = rand() %100;
        arr[index] = Random(1,NUMBERS);
        //printf("i %d => number %d\n", index, days[index]);
    }

    int n = sizeof(arr)/sizeof(arr[0]); 
    bubbleSort(arr, n); 
    printf("Sorted array: \n"); 
    //printArray(arr, n); 

    t = clock() - t;
    double time_taken = ((double)t)/CLOCKS_PER_SEC;
    printf("The program took %f seconds to execute, {%d} items sorted\n", time_taken, n);
    return 0; 
} 
