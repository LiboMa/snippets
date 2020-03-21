#include <stdio.h>
#include<iostream>
#include <bitset>
//#include <random>
#include <cstdlib>
#include <ctime>
#include <string>
#include <array>

using namespace std;
// Function to swap two pointers
#define N 1000000
int O;

void swap(int *a, int *b)
{
    int temp = *a;
    *a = *b;
    *b = temp;
}

// Function to print an array of integers
void printArray(int arr[], int size)
{
    for (int i = 0; i < size; i++)
    {
        printf("%d ", arr[i]);
    }

    printf("\n");
}

// Function to run quicksort on an array of integers
// l is the leftmost starting index, which begins at 0
// r is the rightmost starting index, which begins at array length - 1
void quicksort(int arr[], int l, int r)
{
    // Base case: No need to sort arrays of length <= 1
    if (l >= r)
    {
        return;
    }
    O++;
    
    // Choose pivot to be the last element in the subarray
    int pivot = arr[r];

    // Index indicating the "split" between elements smaller than pivot and 
    // elements greater than pivot
    int cnt = l;

    // Traverse through array from l to r
    for (int i = l; i <= r; i++)
    {
        // If an element less than or equal to the pivot is found...
        if (arr[i] <= pivot)
        {
            // Then swap arr[cnt] and arr[i] so that the smaller element arr[i] 
            // is to the left of all elements greater than pivot
            swap(&arr[cnt], &arr[i]);

            // Make sure to increment cnt so we can keep track of what to swap
            // arr[i] with
            cnt++;
        }
    }
    
    // NOTE: cnt is currently at one plus the pivot's index 
    // (Hence, the cnt-2 when recursively sorting the left side of pivot)
    quicksort(arr, l, cnt-2); // Recursively sort the left side of pivot
    quicksort(arr, cnt, r);   // Recursively sort the right side of pivot
}

// Run a test case of quicksort
// Test case taken from http://geeksquiz.com/quick-sort/
int main() 
{
    int arrayInt[N];
    // const means immutabule
    const int SIZE = N;
    int size = sizeof(arrayInt)/sizeof(int);

    for (int i = 0;i<N;i++) {
        arrayInt[i] = std::rand() % N;
    }

    std::clock_t c_start = std::clock();

    quicksort(arrayInt, 0, N-1);
    printf("Sorted array: \n");
    printArray(arrayInt, N);

    int count;
    std::clock_t c_end = std::clock();
    cout <<"CPU time used " << 1000.0*(c_end - c_start)/CLOCKS_PER_SEC <<"ms\n";
    cout << "size of array: " <<size;
    cout << "\t Complexity O: " <<O;
    cout << endl;

    return 0;
}
