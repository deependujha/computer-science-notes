# Shortest path in DAG (`directed acyclic graph`)

- Do standard BFS traversal, just don't use a boolean visited array.
- Rather, create a distance array, which by default is -1, but stores the minimum distance between source node & index node.

!!! warning
    - In the below code, we're using queue, which increases time complexity.
    - The better approach is to simply use priority queue, and the algorithm is called, **`Dijkstra algorithm`**.
    - The code won't work for graphs having negative weights.

---

## Code to calculate distance from source node

!!! danger "Geeks for Geeks question"
    <a href="https://www.geeksforgeeks.org/problems/shortest-path-in-undirected-graph/1" target="_blank">Shortest path in Directed Acyclic Graph
</a>

```cpp
class Solution {
  public:
     vector<int> shortestPath(int N,int M, vector<vector<int>>& edges){
        // code here
        vector<pair<int,int>> adj[N];
        for(auto &e:edges){
            int x=e[0], y=e[1], w=e[2];
             vector<int> shortestPath(int N, int M, vector<vector<int>> &edges)
            adj[x].push_back({y,w});
        }
        
        // initialize `ans` with -1.
        vector<int> ans(N,-1);
        ans[0]=0; // 0 is the src node.
        
        queue<int> q;
        q.push(0);
        
        while(!q.empty()){
            int node = q.front();
            q.pop();
            
            for(auto&e:adj[node]){
                if(ans[e.first]==-1){
                    ans[e.first]=ans[node]+e.second;
                    q.push(e.first);
                }else if(ans[e.first]>ans[node]+e.second){
                    ans[e.first]=ans[node]+e.second;
                    q.push(e.first);
                }
            }
        }
        
        return ans;
    }
};
```

---

## Exact  route followed for shortest path

- We can simply use `unordered_map parent` to keep track of the parent node that lead the current node to its shortest path.

- Whenever we need to update the distance array, we will also update the `parent`, to account for new parent with shortest route.

- Finally, simply backtrack from destination to source node.


```cpp
    vector<int> shortestPath(int N, int M, vector<vector<int>> &edges)
    {
        // code here
        vector<pair<int, int>> adj[N];
        for (auto &e : edges)
        {
            int x = e[0], y = e[1], w = e[2];

            adj[x].push_back({y, w});
        }

        vector<int> ans(N, -1);
        ans[0] = 0;


        // to keep track of exact path, use parent map
        unordered_map<int, int> parent;
        parent[0] = -1;

        queue<int> q;
        q.push(0);

        while (!q.empty())
        {
            int node = q.front();
            q.pop();

            for (auto &e : adj[node])
            {
                if (ans[e.first] == -1)
                {
                    ans[e.first] = ans[node] + e.second;
                    q.push(e.first);
                    parent[e.first] = node;
                }
                else if (ans[e.first] > ans[node] + e.second)
                {
                    ans[e.first] = ans[node] + e.second;
                    q.push(e.first);
                    parent[e.first] = node;
                }
            }
        }

        // print minimum path between 0 to 2. (should be 0 => 4 => 2)
        int curr = 2;
        cout << "shortest route from (0=>2):\n";
        while (curr != -1 && curr != 0)
        {
            cout << curr << " => ";
            curr = parent[curr];
        }
        if(curr==0)cout << curr << endl;

        return ans;
    }
```
