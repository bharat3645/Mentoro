#include <stdio.h>
#include <stdlib.h>

// Define the BST node structure
struct node {
    int data;
    struct node* left;
    struct node* right;
};

typedef struct node* BSTNODE;

// Function to create a new BST node
BSTNODE newNodeInBST(int item) {
    BSTNODE temp = (BSTNODE)malloc(sizeof(struct node));
    temp->data = item;
    temp->left = temp->right = NULL;
    return temp;
}

// Function to insert a node in BST
BSTNODE insertNodeInBST(BSTNODE node, int ele) {
    if (node == NULL) {
        printf("Successfully inserted.\n");
        return newNodeInBST(ele);
    } else if (ele < node->data) {
        node->left = insertNodeInBST(node->left, ele);
    } else if (ele > node->data) {
        node->right = insertNodeInBST(node->right, ele);
    } else {
        printf("Element already exists in BST.\n");
    }
    return node;
}

// In-order traversal
void inorderInBST(BSTNODE root) {
    if (root != NULL) {
        inorderInBST(root->left);
        printf("%d ", root->data);
        inorderInBST(root->right);
    }
}

// Pre-order traversal
void preorderInBST(BSTNODE root) {
    if (root != NULL) {
        printf("%d ", root->data);
        preorderInBST(root->left);
        preorderInBST(root->right);
    }
}

// Post-order traversal
void postorderInBST(BSTNODE root) {
    if (root != NULL) {
        postorderInBST(root->left);
        postorderInBST(root->right);
        printf("%d ", root->data);
    }
}

// Search in BST
BSTNODE searchNodeInBST(BSTNODE root, int ele) {
    if (root == NULL || root->data == ele) {
        return root;
    } else if (ele < root->data) {
        return searchNodeInBST(root->left, ele);
    } else {
        return searchNodeInBST(root->right, ele);
    }
}

// Main function
int main() {
    BSTNODE root = NULL;
    int choice, x;

    while (1) {
        printf("\nBinary Search Tree Menu:\n");
        printf("1. Insert an element\n");
        printf("2. In-order traversal\n");
        printf("3. Pre-order traversal\n");
        printf("4. Post-order traversal\n");
        printf("5. Search for an element\n");
        printf("6. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                printf("Enter an element to insert: ");
                scanf("%d", &x);
                root = insertNodeInBST(root, x);
                break;
            case 2:
                if (root == NULL) {
                    printf("Binary Search Tree is empty.\n");
                } else {
                    printf("Elements of the BST (in-order traversal): ");
                    inorderInBST(root);
                    printf("\n");
                }
                break;
            case 3:
                if (root == NULL) {
                    printf("Binary Search Tree is empty.\n");
                } else {
                    printf("Elements of the BST (pre-order traversal): ");
                    preorderInBST(root);
                    printf("\n");
                }
                break;
            case 4:
                if (root == NULL) {
                    printf("Binary Search Tree is empty.\n");
                } else {
                    printf("Elements of the BST (post-order traversal): ");
                    postorderInBST(root);
                    printf("\n");
                }
                break;
            case 5:
                printf("Enter an element to search: ");
                scanf("%d", &x);
                if (searchNodeInBST(root, x) == NULL) {
                    printf("Element not found in the binary search tree.\n");
                } else {
                    printf("Element found in the binary search tree.\n");
                }
                break;
            case 6:
                exit(0);
            default:
                printf("Invalid choice. Please try again.\n");
        }
    }

    return 0;
}