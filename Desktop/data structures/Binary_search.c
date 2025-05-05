#include<stdio.h>
void main() {
	int a[20], i, j, n, key, flag = 0, low, high, mid, temp;
	printf("Enter value of n : ");
	scanf("%d", &n);

	for(int i = 0;i < n;i++)
	{ 
	
		printf("Enter element for a[%d] : ",i);
		scanf("%d",&a[i]);
	}

	printf("Enter key element : ");
	scanf("%d", &key);
for(int i = 0;i < n-1;i++)
	{
		for(int j = 0; j < n-i-1;j++)
			{
				if(a[j] > a[j + 1])
				{
					int temp;
					temp = a[j];
					a[j ] = a[j + 1];
					a[j + 1] = temp;
				}
			}
	}
	
	printf("After sorting the elements in the array are\n");
	for(int i = 0;i < n ; i++)
		{
			printf("Value of a[%d] = %d\n",i,a[i]);
		}
	
	low = 0; // Complete the statement
	high = n - 1; // Complete the statement
	
	// Write the code to search an element using  binary search process
	for(i = 0;i < n;i++)
		{
			mid = (high + low) /2;
			if(key == a[mid])
			{
				flag =1;
			}
			else if(key<a[mid])
			{
				high = mid-1;
			}
			else if(key>a[mid])
			{
				low = mid + 1;
			}
		}
	
	
	if (flag == 1 ) { // Write the condition part
	printf("The key element %d is found at the position %d\n", key,mid ); 
		// Complete the statement
	} else {
		printf("The Key element %d is not found in the array\n", key   ); // Complete the statement
}