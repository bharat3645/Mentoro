#include<stdio.h>
void main() {
	int a[20], i, n, j, temp, pos;
	printf("Enter value of n : ");
	scanf("%d", &n);
	// write the for loop to read array elements
	for(int i = 0; i < n; i++)
		{
			printf("Enter element for a[%d] : ",i);
			scanf("%d",&a[i]);
		}
	// write the for loop to display array elements before sorting
	printf("Before sorting the elements in the array are\n");
	for(int i = 0; i < n; i++)
	{
	printf("Value of a[%d] = %d\n",i,a[i]);
	}
	// write the code to sort elements
	printf("After sorting the elements in the array are\n");
	// write the for loop to display array elements after sorting
	for(int i = 0;i < n -1; i++)
		{
			
			pos = i;
			for(int j = i + 1; j < n ; j++)
				{
					if(a[j] < a[pos])
					{
						
						temp = a[pos];
						a[pos] = a[j];
						a[j] = temp;
					}
				}
		}
	
	for(int i = 0 ; i < n; i++)
		{
	printf("Value of a[%d] = %d\n",i,a[i]);
		}
	
}