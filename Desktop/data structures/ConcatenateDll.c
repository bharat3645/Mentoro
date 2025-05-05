#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Node{
int data;
struct Node* prev;
struct Node* next;
};
struct DoublyLinkedList{
struct Node* head;
struct Node* tail;
};
void append(struct DoublyLinkedList * list, int data)
{
	struct Node* newNode = (struct Node*)malloc(sizeof(struct Node));
	newNode -> data = data;
	newNode -> next = NULL;
	if(list -> head == NULL)
	{
		newNode -> prev = NULL;
		list -> head = newNode;
		list -> tail = newNode;
	}
	else
	{
		newNode -> prev = list -> tail;
		list -> tail ->next = newNode;
		list -> tail = newNode;
	}
}
void displayList(struct DoublyLinkedList* list)
{
	struct Node* temp = list -> head;
	while(temp != NULL)
		{
			printf("%d <-> ",temp -> data);
			temp = temp -> next;
		}
	printf("NULL\n");
}

void concatenate(struct DoublyLinkedList* list1, struct DoublyLinkedList* list2)
{
	if(list1 -> tail != NULL && list2 -> head != NULL)
	{
		list1 -> tail -> next = list2 -> head;
		list2 -> head -> prev = list1 -> tail;
		list1 -> tail = list2 -> tail;
	}
}int main() {
    struct DoublyLinkedList* list1 = (struct DoublyLinkedList*)malloc(sizeof(struct DoublyLinkedList));
    list1->head = NULL;
    list1->tail = NULL;

    struct DoublyLinkedList* list2 = (struct DoublyLinkedList*)malloc(sizeof(struct DoublyLinkedList));
    list2->head = NULL;
    list2->tail = NULL;

    int choice, data;

    printf("Elements for List 1 (separated by spaces): ");
    while (1) {
        scanf("%d", &data);
        append(list1, data);
        char c = getchar();
        if (c == '\n') break;
    }

    printf("Elements for List 2 (separated by spaces): ");
    while (1) {
        scanf("%d", &data);
        append(list2, data);
        char c = getchar();
        if (c == '\n') break;
    }

    while (1) {
        printf("1. Display List 1\n");
        printf("2. Display List 2\n");
        printf("3. Concatenate\n");
        printf("4. Quit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                printf("List 1: ");
                displayList(list1);
                break;
            case 2:
                printf("List 2: ");
                displayList(list2);
                break;
            case 3:
                concatenate(list1, list2);
                printf("Lists Concatenated\n");
                break;
            case 4:
                free(list1);
                free(list2);
                return 0;
            default:
                printf("Invalid choice\n");
        }
    }

    return 0;
}

