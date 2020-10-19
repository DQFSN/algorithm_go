/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;

    Node() {}

    Node(int _val) {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node* _left, Node* _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/
class Solution {
	public:
		Node* treeToDoublyList(Node* root) {
	
			stack<Node*> stk;
			Node* head = nullptr;
			Node* cur = nullptr;
			Node* pre = nullptr;

            if (root == nullptr) {
                return root;
            }

            while (root != nullptr || !stk.empty()) {
				while (root != nullptr) {
					stk.push(root);
                    // cout<<root->val<<endl;
					root = root->left;
				}

				if (!stk.empty()) {
                    if (cur == nullptr) {
                        cur = stk.top();
                        head = stk.top();
                    }else {
                        pre = cur;
                        cur = stk.top();
                        pre->right = cur;
                        cur->left = pre;
                    }
					root = stk.top()->right;
					stk.pop();
				}
			}

            cur->right = head;
            head->left = cur;

			return head;
		}
	};