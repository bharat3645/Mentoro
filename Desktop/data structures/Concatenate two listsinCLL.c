#include <stdio.h>
#include <stdlib.h>
typedef struct Node{
int data;
struct Node* next;
}Node;
typedef Node* NODE;
NODE insertEnd(NODE head, int data)
{
	NODE newNode = (NODE)malloc(sizeof(Node));
	newNode -> data = data;
	newNode -> next = newNode;
	if(head == NULL)
		return newNode;
	NODE temp = head;
	while(temp -> next != head)
		{
			temp = temp -> next;
		}
	temp -> next = newNode;
	newNode -> next = head;
	return head;

	temp -> next = newNode;
	newNode -> next = head;
	return head;

	temp -> next = newNode;
	newNode -> next = head;
	return head;
}
void displayList(NODE head)
{
	if (head == NULL)
	{
		printf("Circular Linked List is empty\n");
		return;
	}
	NODE temp = head;
	do{
		printf("%d ",temp-> data);
		temp = temp-> next;
	}while(temp != head);
	printf("\n");
}NODE concatenateLists(NODE head1, NODE head2)
{
	if(head1 == NULL || head2 == NULL)
	{
		printf("Nothing to concatenate\n");
		return head1;
	}
	NODE temp1 = head1;
	while(temp1 -> next != head1)
		temp1 = temp1 -> next;

	NODE temp2 = head2;
	while(temp2 -> next != head2)
		temp2 = temp2 -> next;
	
	temp1 -> next = head2;
	temp2 -> next = head1;
	printf("Lists concatenated.\n");
	return head1;
}

int main() {
    NODE head1 = NULL, head2 = NULL;
    int choice, data;

    do {
        printf("1. Create Circular Linked List 1\n");
        printf("2. Create Circular Linked List 2\n");
        printf("3. Display Circular Linked List 1\n");
        printf("4. Display Circular Linked List 2\n");
        printf("5. Concatenate Lists\n");
        printf("6. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                printf("Enter data to insert into List 1 (-1 to stop): ");
                scanf("%d", &data);
                while (data != -1) {
                    head1 = insertEnd(head1, data);
                    scanf("%d", &data);
                }
                break;
            case 2:
                printf("Enter data to insert into List 2 (-1 to stop): ");
                scanf("%d", &data);
                while (data != -1) {
                    head2 = insertEnd(head2, data);
                    scanf("%d", &data);
                }
                break;
            case 3:
                printf("Circular Linked List 1: ");
                displayList(head1);
                break;
            case 4:
                printf("Circular Linked List 2: ");
                displayList(head2);
                break;
            case 5:
                head1 = concatenateLists(head1, head2);
                break;
            case 6:
                printf("Exiting\n");
                break;
            default:
                printf("Invalid choice\n");
        }
    } while (choice != 6);

    return 0;
}

