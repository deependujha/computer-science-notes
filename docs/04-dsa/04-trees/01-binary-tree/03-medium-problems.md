# Medium Problems related to Binary Tree

## Height of a Binary Tree

- https://leetcode.com/problems/maximum-depth-of-binary-tree/

```cpp
class Solution {
public:
    int maxDepth(TreeNode* root) {
        if(root==NULL)return 0;
        return max(maxDepth(root->left), maxDepth(root->right))+1;
    }
};
```

---

## Check if a Binary Tree is height-balanced

- https://leetcode.com/problems/balanced-binary-tree/description/

```cpp
class Solution {
    int depth(TreeNode* root){
        if(root==NULL)return 0;
        int l = depth(root->left);
        int r = depth(root->right);

        if(l==-1 || r==-1){
            return -1; // we already know tree is not height-balanced
        }

        if(abs(l-r)>1)return -1;

        return max(l,r) + 1;
    }
public:
    bool isBalanced(TreeNode* root) {
        int val = depth(root);
        return val!=-1;
    }
};
```

---

## Diameter of a Binary Tree

- https://leetcode.com/problems/diameter-of-binary-tree/

```cpp
class Solution {
    int depth(TreeNode*root, int* ans){
        if(root==NULL)return 0;
        int l = depth(root->left, ans);
        int r = depth(root->right, ans);
        
        *ans = max(*ans, l+r);

        return max(l,r)+1;
    }
public:
    int diameterOfBinaryTree(TreeNode* root) {
        int ans = 0;
        depth(root, &ans);
        return ans;
    }
};

```

---

## Maximum path sum in a Binary Tree

---

## Check if two Binary Trees are identical

- https://leetcode.com/problems/same-tree/

```cpp
class Solution {
public:
    bool isSameTree(TreeNode* p, TreeNode* q) {
        if(p==NULL && q==NULL)return true;
        if(p==NULL || q==NULL)return false;

        if(p->val != q->val)return false;

        bool l = isSameTree(p->left, q->left);
        bool r = isSameTree(p->right, q->right);

        return l && r;
    }
};
```

---

## Zig-Zag traversal

- https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

```cpp
class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int>> ans;
        if(root==NULL)return ans;

        queue<TreeNode*>q;
        q.push(root);

        int level = 1;
        int levelLength = q.size();

        vector<int> v;

        while(!q.empty()){
            TreeNode* tmp = q.front();
            v.push_back(tmp->val);
            if(tmp->left)q.push(tmp->left);
            if(tmp->right)q.push(tmp->right);
            q.pop();
            levelLength--;

            if(levelLength==0){
                if(level%2){
                    // odd => left to right
                    ans.push_back(v);
                }
                else{
                    // even => right to left
                    reverse(v.begin(), v.end());
                    ans.push_back(v);
                }
                v.clear();

                level++;
                levelLength = q.size();
            }

        }
        return ans;
    }
};
```

---

## Boundary traversal

---

## Vertical order traversal

---

## Top-view of a Binary Tree

---

## Bottom-view of a Binary Tree

---

## Left/right-view of a Binary Tree

---

## Symmetric Binary Tree
