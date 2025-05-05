#include<stdio.h>
#include<stdlib.h>

typedef struct Node{
int data;
struct Node* next;
}Node;
 typedef Node* NODE;
 NODE insertAtPositionInCLL(NODE first, int pos, int x)
{
	NODE newNode = (NODE)malloc(sizeof(NODE));
	newNode -> data = x;
	newNode -> next = NULL;
	
	if(first == NULL)
	{
		if(pos == 1)
		{
			newNode -> next = newNode;
			return newNode;
		}
		else 
		{
			printf("Position not found\n");
			free(newNode);
			return first;
		}
	}
	int count = 0;
	NODE temp = first;
	do
	{
		count++;
		temp = temp -> next;
	}while(temp != first);
	if(pos < 1 || pos > count + 1)
	{
	printf("Position not found\n");
	free(newNode);
	return first;
	}
	if(pos == 1)
	{
		NODE last = first;
		while(last -> next != first)
			last = last -> next;
		newNode -> next = first;
		last -> next = newNode;
		return newNode;
	}
	temp = first;
	
	for(int i = 1; i < pos -1; i++)
		{
			temp = temp -> next;
		}
	newNode -> next = temp -> next;
	temp -> next = newNode;
	return first;
}
NODE deleteAtPositionInCLL(NODE first , int pos)
{
	if(first == NULL)
	{
		printf("CLL is empty\n");
		return NULL;
	}
	int count = 0;
	NODE temp = first;
	do
	{
		count++;
		temp = temp -> next;
	}while(temp != first);
	 if(pos < 1 || pos > count)
	{
		printf("Position not found\n");
		return first;
	}
	if(pos == 1)
	{
		if(first -> next == first)
		{
			printf("Deleted element: %d\n",first -> data);
			free(first);
			return NULL;
		}
		NODE last = first;
		while(last -> next != first)
			last = last -> next;
		NODE temp = first;
		first = first -> next;
		last -> next = first;
		printf("Deleted element: %d\n",temp -> data);
		free(temp);
		return first;
	}
	NODE prev = first;
	for(int i = 1 ; i < pos-1;i++ )
		{
			prev = prev -> next;
		}
	NODE toDelete = prev -> next;
	prev -> next = toDelete -> next;
	printf("Deleted element: %d\n",toDelete->data);
	free(toDelete);
	return first;
}
void traverseListInCLL(NODE first)
{
	NODE temp = first;
	do{
		printf("%d --> ",temp -> data);
		temp = temp -> next;
		
	}while(temp != first);
	printf("\n");
}
NODE reverseCLL(NODE first)
{
	if(first == NULL || first -> next == first)
	{
		return first;
	}
	NODE last = first;
	while(last -> next != first)
		last = last -> next;
	last -> next = NULL;

	NODE prev = NULL, curr = first, next = NULL;

	while(curr != NULL)
		{
			next = curr -> next;
			curr -> next = prev;
			prev = curr;
			curr = next;
		}
	first = prev;
	last = first;
	while(last -> next != NULL)
		last = last -> next;
	last -> next = first;
	return first;
}
void main() {
    NODE first = NULL;
    int x, pos, op;
    while (1) {
        printf("1.Insert 2.Delete 3.Display 4.Reverse 5.Exit\n");
        printf("choice: ");
        scanf("%d", &op);
        switch (op) {
            case 1:
                printf("Position: ");
                scanf("%d", &pos);
                if (pos <= 0) {
                    printf("Position not found\n");
                } else {
                    printf("Element: ");
                    scanf("%d", &x);
                    first = insertAtPositionInCLL(first, pos, x);
                }
                break;
            case 2:
                if (first == NULL) {
                    printf("CLL is empty\n");
                } else {
                    printf("Position: ");
                    scanf("%d", &pos);
                    first = deleteAtPositionInCLL(first, pos);
                }
                break;
            case 3:
                if (first == NULL) {
                    printf("CLL is empty\n");
                } else {
                    printf("Elements in CLL are: ");
                    traverseListInCLL(first);
                }
                break;
            case 4:
                if (first == NULL) {
                    printf("CLL is empty\n");
                } else {
                    first = reverseCLL(first);
                    printf("CLL reversed\n");
                }
                break;
            case 5:
                exit(0);
        }
    }
}