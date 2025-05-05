#include<stdio.h>
//#include<conio.h>

#define MAX 20
#define INF 999
int cost[MAX][MAX],n,e,i,j,s,d,w;

void prims() {
	int selected[MAX] = {0},no_of_edges=0;
	int x,y,min,total_cost=0;

	selected[1]=1;

	while(no_of_edges < n-1){
		min = INF;
		for(i=1;i<=n ;i++){
			if(selected[i]){
				for(j=1;j<=n;j++){
					if(!selected[j] && cost[i][j] <min){
						min = cost[i][j];
						x=i;
						y=j;
					}
				}
			}
		}
	printf("Edge cost from %d to %d : %d\n",x,y,cost[x][y]);
	total_cost += cost[x][y];
	selected[y] =1;
	no_of_edges++ ;
	}
	printf("Minimum cost of spanning tree = %d\n",total_cost);
    return 0;
}

int main() {
	printf("Enter the number of vertices : ");
	scanf("%d", &n);
	printf("Enter the number of edges : ");
	scanf("%d", &e);
	for(i=1; i<=e; i++) {
		printf("Enter source : ");
		scanf("%d", &s);
		printf("Enter destination : ");
		scanf("%d", &d);
		printf("Enter weight : ");
		scanf("%d", &w);
		if(s <= 0 || d <= 0 || s > n || d > n || w < 0 ) {
			printf("Invalid data.Try again.\n");
			i--;
			continue;
		}
		cost[s][d] = w;
		cost[d][s] = w;
	}
	for(i=1; i <= n; i++) {
	    for(j=1; j<=n; j++) {
			if(cost[i][j] == 0)
			    cost[i][j] = 999;
	    }
	}
	printf("The edges of Minimum Cost Spanning Tree are : \n");
	prims();
    return 0;
}