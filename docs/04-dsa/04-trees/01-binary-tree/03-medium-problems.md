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

- https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/description/

```cpp
class Solution {
    void solver(TreeNode* root, int h, int v, map<int, map<int, vector<int>>>*mp){
        if(root==NULL)return;
        (*mp)[h][v].push_back(root->val);
        solver(root->left, h-1, v+1, mp);
        solver(root->right, h+1, v+1, mp);
    }
public:
    vector<vector<int>> verticalTraversal(TreeNode* root) {
        map<int, map<int, vector<int>>>mp; // horizontal, vertical, values
        solver(root, 0, 0, &mp);

        vector<vector<int>> ans;
        for(auto &e:mp){
            vector<int> v;
            for(auto &f:e.second){
                sort(f.second.begin(), f.second.end());
                for(auto&z:f.second)v.push_back(z);
            }
            ans.push_back(v);
        }

        return ans;

    }
};
```

---

## Top-view of a Binary Tree

- https://www.geeksforgeeks.org/problems/top-view-of-binary-tree/1

```cpp
class Solution
{
    void solver(Node* root, int x, int y, map<int,pair<int,int>> *mp){
        if(root==NULL)return;

        if(((*mp).find(x)==(*mp).end()) || (*mp)[x].second > y)
            (*mp)[x]={root->data, y};
        
        solver(root->left, x-1, y+1, mp);
        solver(root->right, x+1, y+1, mp);
    }
    public:
    //Function to return a list of nodes visible from the top view 
    //from left to right in Binary Tree.
    vector<int> topView(Node *root)
    {
        //Your code here
        map<int,pair<int,int>> mp; // mp[x]={val, y}
        solver(root, 0, 0, &mp);
        vector<int> ans;
        for(auto &e: mp)
            ans.push_back(e.second.first);
        return ans;
    }

};
```

---

## Bottom-view of a Binary Tree

- https://www.geeksforgeeks.org/problems/bottom-view-of-binary-tree/1

```cpp
class Solution {
    
    void solver(Node*root, int x, int y, map<int, pair<int,int>>*mp){
        if(root==NULL)return;
        // inorder traversal, so rightmost will come at end
        solver(root->left, x-1, y+1, mp);
        
        // '=' to account for same (x,y) node and bcoz of inorder traversal, rightmost will come at end
        if(((*mp).find(x)==(*mp).end()) || (*mp)[x].second <= y) 
            (*mp)[x]={root->data, y};

        solver(root->right, x+1, y+1, mp);
    }
    
  public:
    vector <int> bottomView(Node *root) {
        // Your Code Here
        map<int,pair<int,int>>mp; // mp[x]={val, y}
        solver(root, 0, 0, &mp);
        
        vector<int> ans;

        for(auto &e:mp){
            ans.push_back(e.second.first);
        }
        return ans;
    }
};
```

---

## Left/right-view of a Binary Tree

- https://leetcode.com/problems/binary-tree-right-side-view/description/

```cpp
class Solution {
    vector<int> levelOrderTraversal(TreeNode* root){
        vector<int> ans;
        if(root==NULL)return ans;
        queue<TreeNode*> q;

        q.push(root);

        while(!q.empty()){
            ans.push_back(q.back()->val);

            int N = q.size();
            for(int i=0;i<N;i++){
                TreeNode* currNode = q.front();
                q.pop();

                if(currNode->left)q.push(currNode->left);
                if(currNode->right)q.push(currNode->right);
            }
        }

        return ans;
    }
public:
    vector<int> rightSideView(TreeNode* root) {
        // do level-wise traversal and insert the last element of each level
        return levelOrderTraversal(root);
    }
};

```

---

## Symmetric Binary Tree

- https://leetcode.com/problems/symmetric-tree/description/

- We can also do this by checking if level order traversal array is same as reversed level order traversal array. But, not so straightforward. We've to account for null nodes.

```cpp
class Solution {
    bool isHelper(TreeNode* left, TreeNode* right){
        if(left==NULL && right==NULL)return true;
        if(left==NULL || right==NULL)return false;

        if(left->val != right->val)return false;
        return isHelper(left->left, right->right) && isHelper(left->right, right->left);
        
    }

public:
    bool isSymmetric(TreeNode* root) {
        if(root==NULL)return true;
        return isHelper(root->left, root->right);
    }
};
```
