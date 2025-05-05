#include<stdio.h>
void main() {
	int a[20], i, n, j, large, index;
	printf("Enter value of n : ");
	scanf("%d", &n);
	// Write the code to read an array elements
	for(int i = 0; i < n; i++)
		{
			printf("Enter element for a[%d] : ",i);
			scanf("%d",&a[i]);
		}
	
	printf("Before sorting the elements in the array are\n");
	// Write the code to print the given array elements before sorting
	for(int i = 0;i <n;i++)
		{
			printf("Value of a[%d] = %d\n",i,a[i]);
		}
	
	// Write the code for selection sort largest element method
	
	printf("After sorting the elements in the array are\n");
	// Write the code to print the given array elements after sorting
	for(int i = n - 1; i >= 1; i--)
		{
			large = 0;
			index = i;
	
			for(int j = i ; j >= 0;j--)
				{
					if(a[j] > a[index])
					{
						index = j;
					}
				}
		
	large = a[index];
	a[index] = a[i];
	a[i] = large;
		}
	for(int i = 0 ; i < n; i++)
		{
			printf("Value of a[%d] = %d\n",i,a[i]);
		}
	
}