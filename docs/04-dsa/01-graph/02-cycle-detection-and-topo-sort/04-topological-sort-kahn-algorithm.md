# Topological Sort (`Kahn's algorithm`)

An iterative approach to get topological sort of the graph.

## Algorithm

```code
in_degree_list ← Vector containing in-degree of all the nodes at index
L ← Empty list that will contain the topologically sorted elements
Q ← Queue of all nodes with no incoming edge (in-degree 0)

while Q is not empty do
    remove a node n from Q
    add n to L
    for each node m with an edge e from n to m do
        remove edge e from the graph (in_degree_list[m]--;)
        if m has no other incoming edges (in_degree_list[m]==0);
            then insert m into Q

if graph has edges (`len(L) != no. of nodes`) then
    return error   (graph has at least one cycle)
else 
    return L   (a topologically sorted order)
```

---

## Code

```cpp
// Including necessary header file
#include <bits/stdc++.h>
using namespace std;

// Function to return list containing vertices in
// Topological order.
vector<int> topologicalSort(vector<vector<int> >& adj, int V)
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
		cout << "Graph contains cycle!" << endl;
		return {};
	}

	return result;
}
```