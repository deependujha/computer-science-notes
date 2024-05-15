# Prim's Algorithm (`minimum spanning tree`)

## Algorithm

1. Create a boolean visited array that **stores if the node is already present in MST**.
2. Create an empty min-priority queue.
3. Priority queue should store three numbers: {weight, node, parent_node}.
4. This structure makes sure, at the top of the priority queue, we have the least weighted node.
5. Push first node into priority queue and mark it as visited.
6. Now, iterate till the priority queue is not empty.
7. Pick the top item from priority queue and pop it. Now, iterate for all the nodes connected with it and insert them into priority queue in the order: {weight, child_node, pq_top_node}
8. Once done, check if pq_top_node is not already visited, and then either insert into your MST vector answer (node, parent_node) or add to the sum of MST (add weight).
9. Repeat from step 6.

---

## Code

!!! danger "Codestudio Prims algorithm MST"
    <a href="https://www.naukri.com/code360/problems/prim-s-mst_1095633" target="_blank">Codestudio Prim's algorithm MST</a>

```cpp
#include <bits/stdc++.h> 
vector<pair<pair<int, int>, int>> calculatePrimsMST(int n, int m, vector<pair<pair<int, int>, int>> &g)
{
    // Write your code here.
    int mst_sum = 0;
    vector<pair<pair<int, int>, int>> ans;

   vector<pair<int,int>> adj[n+1];
    for(auto &e:g){
        int x =e.first.first, y=e.first.second, w = e.second;


        adj[x].push_back({y, w});
        adj[y].push_back({x, w});
    }

    bool visited[n+1]={false};

    // max heap code
    // priority_queue<pair<pair<int,int>,int>> pq; // weight, node, parent

    // min heap code
    priority_queue<pair<pair<int,int>,int>, vector<pair<pair<int,int>,int>>, greater<pair<pair<int,int>,int>>> pq;

    pq.push({{1,1},-1});
    visited[1]=true;

    while(!pq.empty()){
        auto node = pq.top();
        pq.pop();

        int weight = node.first.first;
        int curr_node = node.first.second;
        int parent_node=node.second;

        for(auto&e:adj[curr_node]){
            if (visited[e.first] == false)
                pq.push({{e.second, e.first}, curr_node});
        }

        if(visited[curr_node]==false){
            visited[curr_node]=true;
            ans.push_back({{curr_node,parent_node},weight});
            mst_sum+=weight;
        }
    }

    // or return mst_sum (depending on the problem statement)
    return ans;
}

```
