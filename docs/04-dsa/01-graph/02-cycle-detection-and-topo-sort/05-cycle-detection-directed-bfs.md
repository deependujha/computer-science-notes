# Cycle detection in Directed Graph (BFS)

We use the very same code as in `Kahn's algorithm`.

- If the length of the topo sort array is equal to no. of nodes, no cycle is present.
- Else, graph is cyclic.

---

## Code

!!! danger "Code studio detect cycle in a directed graph question"
    <a href="https://www.naukri.com/code360/problems/detect-cycle-in-a-directed-graph_1062626?leftPanelTabValue=PROBLEM" target="_blank">CodeStudio detect cycle in a directed graph</a>

```cpp
// Function to return list containing vertices in
// Topological order.
bool bfs_cycle_detection_kahn_algorithm_topologicalSort(vector<int> adj[], int V)
{
    // Vector to store indegree of each vertex
    vector<int> indegree(V);
    for (int i = 0; i < V; i++) {
        for (auto it : adj[i]) {
            indegree[it]++;
        }
    }

    // Queue to store vertices with indegree 0
    queue<int> q;
    for (int i = 0; i < V; i++) {
        if (indegree[i] == 0) {
            q.push(i);
        }
    }
    vector<int> result;
    while (!q.empty()) {
        int node = q.front();
        q.pop();
        result.push_back(node);

        // Decrease indegree of adjacent vertices as the
        // current node is in topological order
        for (auto it : adj[node]) {
            indegree[it]--;

            // If indegree becomes 0, push it to the queue
            if (indegree[it] == 0)
                q.push(it);
        }
    }

    // Check for cycle
    if (result.size() != V) {
        return true;
    }

    return false;
}

int detectCycleInDirectedGraph(int n, vector < pair < int, int >> & edges) {
  // Write your code here.
  vector<int> adj_list[n+1];
  for(auto &e:edges){
    int x= e.first;
    int y= e.second;
    adj_list[x].push_back(y);
  }

 return bfs_cycle_detection_kahn_algorithm_topologicalSort(adj_list, n+1);

}
```