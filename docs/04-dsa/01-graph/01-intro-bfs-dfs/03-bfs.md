# Breadth-First Search (BFS)

![animated bfs gif](../../../images/dsa/graph/Animated_BFS.gif)

- **BFS is the level order traversal of tree**.

---

## BFS Vs DFS

![bfs vs dfs](../../../images/dsa/graph/bfs-vs-dfs.png)

---

## Trick to remember

!!! danger "Totka 🦹‍♂️"
    bfs queue (kyu)?

    dfs queue (kyu) nhi!

    ---

    BFS - queue ✅

    DFS - stack (or recursion) ✅

---

## BFS Code (`using Queue`)

- Time complexity of BFS traversal is: **`O (n + e)`**.

!!! code

    === "C++"

        ```cpp
        #include <bits/stdc++.h>
        using namespace std;

        class Solution {
        public:
            vector<int> bfsOfGraph(int V, vector<int> adj[])
            {
                // V: no of vertex (nodes)

                int vis[V] = {0};
                vector<int> bfs;

                // iterate for non-visited nodes (to handle disconnected components)
                for (int currIdx = 0; currIdx < V; currIdx++)
                {
                    if (vis[currIdx] == 1)
                        continue;
                    vis[currIdx] = 1;
                    queue<int> q;
                    // push the initial starting node
                    q.push(currIdx);
                    // iterate till the queue is empty
                    while (!q.empty())
                    {
                        // get the topmost element in the queue
                        int node = q.front();
                        q.pop();
                        bfs.push_back(node);
                        // traverse for all its neighbors
                        for (auto it : adj[node])
                        {
                            // if the neighbor has previously not been visited,
                            // store in Q and mark as visited
                            if (!vis[it])
                            {
                                vis[it] = 1;
                                q.push(it);
                            }
                        }
                    }
                }
                return bfs;
            }

        };

        void addEdge(vector <int> adj[], int u, int v) {
            adj[u].push_back(v);
            adj[v].push_back(u);
        }

        void printAns(vector <int> &ans) {
            for (int i = 0; i < ans.size(); i++) {
                cout << ans[i] << " ";
            }
        }

        int main() 
        {

            vector <int> adj[6];
            
            addEdge(adj, 0, 1);
            addEdge(adj, 1, 2);
            addEdge(adj, 1, 3);
            addEdge(adj, 0, 4);

            Solution obj;
            vector <int> ans = obj.bfsOfGraph(5, adj);
            printAns(ans);

            return 0;
        }

        ```

    === "Python"

        ```python
        print("Hello World")
        ```
