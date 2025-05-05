#include<stdio.h>
#include<string.h>
void main() {
	char a[20][20];
	int i, n, j;
	char temp[20];
	printf("Enter value of n : ");
	scanf("%d", &n);
	for(int i = 0; i< n; i++)
		{
			printf("Enter string for a[%d] : ",i);
			scanf("%s",a[i]);
		}
	printf("Before sorting the strings in the array are\n");
	//Write the code to print the strings before sorting
	for(int j = 0; j < n;j++)
		{
			printf("String at a[%d] = %s\n",j,a[j]);
		}
	
	//Write the code to sort the strings using bubble sort technique
	for(int i = 0; i < n-1; i++ )
		{
			for(int j = 0; j<n-i-1;j++)
				{
					if(strcmp(a[j] ,a[j+1]) > 0)
					{
						strcpy(temp,a[j]);
						strcpy(a[j],a[j+1]);
						strcpy(a[j+1],temp);
					}
				}
		}
	
	printf("After sorting the strings in the array are\n");
	for(int i =0;i < n; i++)
		{
			printf("String at a[%d] = %s\n",i,a[i]);
		}
	
}