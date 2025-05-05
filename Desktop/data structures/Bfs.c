#include<stdio.h>
#include<stdlib.h>
#define MAX 100
struct node{
int vertex;
struct node* next;
};
typedef struct node* GNODE;
GNODE graph[MAX];
int visited[MAX];
void DFS(int v)
{
	visited[v] = 1;
	printf("\n%d",v);
	GNODE p = graph[v];
	while(p != NULL)
		{
			if(!visited[p->vertex])
				DFS(p->vertex);
			p = p->next;
		}
}
void main() {
	int n, E, s, d, i, j, v;
	GNODE p, q;
	printf("Enter no of vertices: ");
	scanf("%d",&n);
	printf("Enter no of edges: ");
	scanf("%d",&E);
	for(i=1;i<=E;i++) {
		printf("Enter source: ");
		scanf("%d",&s);
		printf("Enter destination: ");
		scanf("%d",&d);
		q=(GNODE)malloc(sizeof(struct node));
		q->vertex=d;
		q->next=NULL;
		if(graph[s]==NULL) {
			graph[s]=q;
		} else {
			p=graph[s];
			while(p->next!=NULL)
				p=p->next;
			p->next=q;
		}
	}
	for(i=1;i<=n;i++)
		visited[i]=0;
	printf("Enter Start Vertex for BFS: ");
	scanf("%d", &v);
	printf("BFS of graph: ");
	BFS(v);
	printf("\n");
}