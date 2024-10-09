# Hard Problems related to Binary Tree

## Lowest Common Ancestor (LCA) in a Binary Tree

- https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/

```cpp
class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        if(root==NULL || root==p || root==q)return root;

        TreeNode* left = lowestCommonAncestor(root->left, p, q);

        TreeNode* right = lowestCommonAncestor(root->right, p, q);

        if(left!=NULL && right!=NULL){
            // both left and right has values
            // means, current node is LCA
            return root;
        }
        else if(left!=NULL){
            return left;
        }
        return right; // could be a node, or null (both valid)
        
    }
};
```

---

## Root to Leaf Paths

- https://www.geeksforgeeks.org/problems/root-to-leaf-paths/1

```cpp
class Solution {
    void solver(Node*root, vector<vector<int>> &ans, vector<int> currPath){
        if(root==NULL){
            return;
        }
        
        currPath.push_back(root->data);
        if(root->left)solver(root->left, ans, currPath);
        if(root->right)solver(root->right, ans, currPath);
        
        if(root->left==NULL && root->right==NULL)
            ans.push_back(currPath);
    }
  public:
    vector<vector<int>> Paths(Node* root) {
        // code here
        vector<vector<int>> ans;
        vector<int> currPath;
        solver(root, ans, currPath);
        return ans;
    }
};
```

---
