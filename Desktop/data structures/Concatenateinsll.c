#include <stdio.h>
#include <stdlib.h>

// Structure definition
typedef struct node {
    int data;
    struct node* next;
} Node;
typedef Node* NODE;

// Function to add a node at the end
NODE addNode(NODE head, int data) {
    NODE newNode = (NODE)malloc(sizeof(Node));  // Corrected sizeof to use Node, not NODE
    newNode->data = data;
    newNode->next = NULL;

    if (head == NULL) {
        head = newNode;
    } else {
        NODE temp = head;
        while (temp->next != NULL) {
            temp = temp->next;
        }
        temp->next = newNode;
    }
    return head;
}

// Function to display the list
void displayList(NODE head) {
    NODE temp = head;
    while (temp != NULL) {
        printf("%d --> ", temp->data);
        temp = temp->next;
    }
    printf("NULL\n");
}

// Function to concatenate two lists
NODE concatenate(NODE list1, NODE list2) {
    if (list1 == NULL) {
        return list2;
    }
    NODE temp = list1;
    while (temp->next != NULL) {
        temp = temp->next;
    }
    temp->next = list2;
    return list1;
}

// Main program
int main() {
    NODE list1 = NULL, list2 = NULL;
    int choice, x;

    while (1) {
        printf("\n----- MENU -----\n");
        printf("1. Create SLL1\n");
        printf("2. Create SLL2\n");
        printf("3. Display SLL1\n");
        printf("4. Display SLL2\n");
        printf("5. Concatenate SLL1 & SLL2\n");
        printf("6. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                printf("Enter elements for SLL1 (-1 to end): ");
                scanf("%d", &x);
                while (x != -1) {
                    list1 = addNode(list1, x);
                    scanf("%d", &x);
                }
                break;
            case 2:
                printf("Enter elements for SLL2 (-1 to end): ");
                scanf("%d", &x);
                while (x != -1) {
                    list2 = addNode(list2, x);
                    scanf("%d", &x);
                }
                break;
            case 3:
                printf("SLL1: ");
                displayList(list1);
                break;
            case 4:
                printf("SLL2: ");
                displayList(list2);
                break;
            case 5:
                list1 = concatenate(list1, list2);
                printf("Concatenated List (SLL1 + SLL2): ");
                displayList(list1);
                break;
            case 6:
                exit(0);
            default:
                printf("Invalid choice. Please try again.\n");
        }
    }

    return 0;
}