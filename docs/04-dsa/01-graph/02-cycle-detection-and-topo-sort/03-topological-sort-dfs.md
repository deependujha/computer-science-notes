# Topological Sort (using `DFS`)

## What is topological sorting?

- Topological sorting for Directed Acyclic Graph (DAG) is a linear ordering of vertices such that for every directed edge u-v, vertex u comes before v in the ordering.
- **not possible if the graph is not a DAG**.

![topo sort graph](../../../images/dsa/graph/topo-sort-graph.png)

Topo sort of the above graph:
    5 4 2 3 1 0

    Explanation:
        since, there's an edge from (5->0),
        so in topo sort of the graph, 5 must come become 0.

---

## Approach (using DFS)

1. Create an empty stack, we will use it store the nodes in their topo sort.
2. Use DFS traversal to traverse the graph (same code as you would for cycle detection (`bool dfs_visited[]`)).
3. When a node has traversed all its connected nodes, push the node into the stack.
4. Repeat.

> ✅ Stack from top contains one of the possible topo sort of the graph.

??? hint
    - Instead of stack, you can also use vector.
    - If you already know `graph is not cyclic`, you can push the nodes from end in vector.
    - If you're not sure, push nodes in vector, and once done, reverse the vector.

---

## Leetcode question

!!! danger "Leetcode topological sort question"
    <a href="https://leetcode.com/problems/course-schedule-ii/description/" target="_blank">Course Schedule-2 leetcode medium</a>

```cpp
class Solution {
public:

    bool dfs_topo_sort(vector<int> adj_list[], int n, bool visited[], bool dfs_visited[], int node, vector<int> &ans){
        visited[node]=true;
        dfs_visited[node]=true;

        for(auto&e:adj_list[node]){
            if(visited[e]==false){
                bool return_val = dfs_topo_sort(adj_list,n,visited,dfs_visited,e, ans);
                if(return_val==false){return false;}
            }
            else if(dfs_visited[e]==true){
                ans.clear();
                // false return - cyclic graph (no topo sort possible)
                return false;
            }
        }
        ans.push_back(node);
        dfs_visited[node]=false;
        return true;
    }

    vector<int> findOrder(int numCourses, vector<vector<int>>& prerequisites) {
        vector<int> adj_list[numCourses];
        for(auto &e:prerequisites){
            int x = e[0], y = e[1];

            //  for pair [0, 1] => 1 before 0
            adj_list[y].push_back(x);
        }

        bool visited[numCourses];
        bool dfs_visited[numCourses];
        for(auto&e:visited)e=false;
        for(auto&e:dfs_visited)e=false;

        vector<int> ans;


        for(int i=0;i<numCourses;i++){
            if(visited[i]==true)continue;
            bool return_val = dfs_topo_sort(adj_list,numCourses,visited,dfs_visited,i,ans);
            if(return_val==false)break;
        }

        reverse(ans.begin(),ans.end());
        return ans;
    }
};
```
