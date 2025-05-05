#include <stdio.h>
#include <stdlib.h>

struct Node{
int data;
struct Node* prev;
struct Node* next;
};
struct Node* head = NULL;
struct Node* tail = NULL;
void insert(int data)
{
	struct Node* newNode = (struct Node*)malloc(sizeof(struct Node));
	newNode -> data = data;
	newNode -> next = NULL;
	newNode -> prev = NULL;

	if(head == NULL)
	{
		newNode -> prev = NULL;
		head = newNode;
		tail = newNode;
	}
	else {
		struct Node* temp = head;
		while(temp -> next != NULL)
			{
				temp = temp -> next;
			}
		temp -> next = newNode;
		newNode -> prev = temp;
		tail = newNode;
	}
} 
void remov(int data)
{
	struct Node* temp = head;
	while(temp != NULL)
		{
			if(temp -> data == data)
			{
				if(temp -> prev != NULL)
					{
						temp -> prev -> next = temp ->  next ;
					}
					else
					{
					head = temp -> next;	
					}
				if(temp -> next != NULL)
				{
					temp -> next -> prev = temp -> prev;
				}
				free(temp);
				return;
				}
			temp = temp -> next;
			}
	printf("%d not found\n",data);
}

void display()
{
	struct Node* temp = head;
	while(temp != NULL)
		{
			printf("%d\t",temp -> data);
			temp = temp -> next;
		}
	printf("\n");
}void reverse()
{
	struct Node* temp = NULL;
	struct Node* current = head;

while(current != NULL)
{
	temp = current -> prev;
	current -> prev = current -> next;
	current -> next = temp;
	current = current -> prev;
}
if(temp != NULL)
	{
	head = temp -> prev;
	}
}

int main() {
    int n, ch;
    do {
        printf("Operations on doubly linked list\n");
        printf("1.Insert\n2.Remove\n3.Display\n4.Reverse\n0.Exit\n");
        printf("Enter Choice 0-4: ");
        scanf("%d", &ch);
        switch (ch) {
            case 1:
                printf("Enter number: ");
                scanf("%d", &n);
                insert(n);
                break;
            case 2:
                printf("Enter number to delete: ");
                scanf("%d", &n);
                remov(n);
                break;
            case 3:
                display();
                break;
            case 4:
                reverse();
                printf("List reversed\n");
                break;
        }
    } while (ch != 0);
    return 0;
}


