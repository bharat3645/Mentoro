#include <stdio.h>

// Function to display the array
void display(int arr[], int n) {
    for(int i = 0; i < n; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

// Function to merge two sorted subarrays
void merge(int arr[], int low, int mid, int high) {
    int i = low, h = low, j = mid + 1, k;
    int temp[50];

    while(h <= mid && j <= high) {
        if(arr[h] <= arr[j]) {
            temp[i] = arr[h];
            h++;
        } else {
            temp[i] = arr[j];
            j++;
        }
        i++;
    }

    // Copy remaining elements from left subarray (if any)
    if(h <= mid) {
        for(k = h; k <= mid; k++) {
            temp[i] = arr[k];
            i++;
        }
    } else {
        // Copy remaining elements from right subarray (if any)
        for(k = j; k <= high; k++) {
            temp[i] = arr[k];
            i++;
        }
    }

    // Copy the sorted elements back to the original array
    for(k = low; k <= high; k++) {
        arr[k] = temp[k];
    }
}

// Recursive function to split the array and then merge
void splitAndMerge(int arr[], int low, int high) {
    if(low < high) {
        int mid = (low + high) / 2;
        splitAndMerge(arr, low, mid);
        splitAndMerge(arr, mid + 1, high);
        merge(arr, low, mid, high);
    }
}

// Main function
int main() {
    int arr[15], i, n;
    printf("Enter array size (max 15): ");
    scanf("%d", &n);
    printf("Enter %d elements: ", n);
    for(i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }

    printf("Before sorting the elements are: ");
    display(arr, n);

    splitAndMerge(arr, 0, n - 1);

    printf("After sorting the elements are: ");
    display(arr, n);

    return 0;
}
