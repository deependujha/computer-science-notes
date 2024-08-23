# Traversals in Binary Tree

![tree traversal](../../../images/dsa/tree/tree-traversal.png)

---

## Depth-first Search (DFS)

- Pre-order traversal
- In-order traversal
- Post-order traversal

=== "pre-order"

    ```c++
    // pre-order: root, left, right

    void preOrder(Node* node){
        if (node == NULL){
            return;
        }

        cout << node->data << " ";
        preOrder(node->left);
        preOrder(node->right);
    }
    ```

=== "In-order"

    ``` c++
    // in-order: left, root, right

    void inOrder(Node* node){
        if (node == NULL){
            return;
        }

        inOrder(node->left);
        cout << node->data << " ";
        inOrder(node->right);
    }
    ```

=== "Post-order"

    ```c++
    // post-order: left, right, root

    void postorder(Node* node){
        if (node == NULL){
            return;
        }

        postorder(node->left);
        postorder(node->right);
        cout << node->data << " ";
    }
    ```

---

## Breadth-first Search (BFS)

- Level-order traversal

=== "Level-order"

    ```c++
    // level-order: root, left, right

    void levelOrder(Node* node){
        if (node == NULL){
            return;
        }

        queue<Node*> q;
        q.push(node);

        while (!q.empty()){
            Node* t = q.front();
            q.pop();

            cout << t->data << " ";

            if (t->left != NULL){
                q.push(t->left);
            }
            if (t->right != NULL){
                q.push(t->right);
            }
        }
    }
    ```

---

## Iterative traversal

### pre-order iterative code

```cpp
void preOrderIterative(Node* node){
    if (node == NULL){
        return;
    }

    stack<Node*> s;
    s.push(node);

    while (!s.empty()){
        Node* t = s.top();
        s.pop();

        cout << t->data << " ";

        if (t->right != NULL){
            s.push(t->right);
        }

        if (t->left != NULL){
            s.push(t->left);
        }
    }
}
```

### in-order iterative code


### post-order iterative (with 2 stacks)

### post-order iterative (with 2 stacks)


---

## Pre-order, In-order, Post-order in one traversal
