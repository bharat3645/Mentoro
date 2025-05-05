#include <stdio.h>
#include <stdlib.h>

#define MAX 100

typedef struct node {
    int vertex;
    struct node* next;
} *GNODE;

GNODE graph[MAX];

// Function to insert a new vertex
void insertVertex(int *N) {
    (*N)++;
    graph[*N] = NULL;
    printf("Vertex %d inserted successfully.\n", *N);
}

// Function to insert a new edge
void insertEdge(int *N) {
    int s, d;
    printf("Enter source vertex: ");
    scanf("%d", &s);
    printf("Enter destination vertex: ");
    scanf("%d", &d);
    if(s > *N || d > *N || s <= 0 || d <= 0) {
        printf("Invalid vertices.\n");
        return;
    }

    GNODE q = (GNODE)malloc(sizeof(struct node));
    q->vertex = d;
    q->next = NULL;

    if(graph[s] == NULL)
        graph[s] = q;
    else {
        GNODE p = graph[s];
        while(p->next != NULL)
            p = p->next;
        p->next = q;
    }
    printf("Edge inserted successfully.\n");
}

// Function to delete a vertex
void deleteVertex(int *N) {
    int v;
    printf("Enter the vertex to be deleted: ");
    scanf("%d", &v);
    if(v > *N || v <= 0 || graph[v] == NULL) {
        printf("Invalid vertex.\n");
        return;
    }

    // Free the list at vertex v
    GNODE temp;
    while(graph[v] != NULL) {
        temp = graph[v];
        graph[v] = graph[v]->next;
        free(temp);
    }
    graph[v] = NULL;

    // Remove edges pointing to v
    for(int i = 1; i <= *N; i++) {
        if(graph[i] == NULL) continue;
        GNODE curr = graph[i], prev = NULL;
        while(curr != NULL) {
            if(curr->vertex == v) {
                if(prev == NULL)
                    graph[i] = curr->next;
                else
                    prev->next = curr->next;
                free(curr);
                break;
            }
            prev = curr;
            curr = curr->next;
        }
    }

    printf("After deleting vertex the adjacency list is:\n");
    print(N);
}

// Function to delete an edge
void deleteEdge(int *N) {
    int src, des;
    printf("Enter the source vertex of the edge: ");
    scanf("%d", &src);
    printf("Enter the destination vertex of the edge: ");
    scanf("%d", &des);

    if(src > *N || des > *N || src <= 0 || des <= 0 || graph[src] == NULL) {
        printf("Invalid edge.\n");
        return;
    }

    GNODE curr = graph[src], prev = NULL;
    while(curr != NULL && curr->vertex != des) {
        prev = curr;
        curr = curr->next;
    }

    if(curr == NULL) {
        printf("Edge not found.\n");
        return;
    }

    if(prev == NULL)
        graph[src] = curr->next;
    else
        prev->next = curr->next;

    free(curr);
    printf("After deleting edge the adjacency list is:\n");
    print(N);
}

// Function to print the adjacency list
void print(int *N) {
    for(int i = 1; i <= *N; i++) {
        printf("Vertex %d: ", i);
        GNODE temp = graph[i];
        while(temp != NULL) {
            printf("%d -> ", temp->vertex);
            temp = temp->next;
        }
        printf("NULL\n");
    }
}

// Main function
int main() {
    int x, op;
    int N, E, s, d, i;
    GNODE p, q;

    printf("Enter the number of vertices: ");
    scanf("%d", &N);
    printf("Enter the number of edges: ");
    scanf("%d", &E);

    for(i = 1; i <= E; i++) {
        printf("Enter source: ");
        scanf("%d", &s);
        printf("Enter destination: ");
        scanf("%d", &d);

        q = (GNODE)malloc(sizeof(struct node));
        q->vertex = d;
        q->next = NULL;

        if(graph[s] == NULL)
            graph[s] = q;
        else {
            p = graph[s];
            while(p->next != NULL)
                p = p->next;
            p->next = q;
        }
    }

    while(1) {
        printf("\n1.Insert vertex 2.Insert edge 3.Delete vertex 4.Delete edge 5.Print adjacency list 6.Exit\n");
        printf("Enter your option: ");
        scanf("%d", &op);
        switch(op) {
            case 1:
                insertVertex(&N);
                break;
            case 2:
                insertEdge(&N);
                break;
            case 3:
                deleteVertex(&N);
                break;
            case 4:
                deleteEdge(&N);
                break;
            case 5:
                print(&N);
                break;
            case 6:
                exit(0);
        }
    }
    return 0;
}