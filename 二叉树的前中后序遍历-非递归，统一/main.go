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
			Node* preRoot = nullptr;
			while (root != nullptr) {
				while (root != nullptr) {
					stk.push(root);
					// cout<<stk.top()->val<<endl; 先根遍历,一步输出
					root = root->left;
					// if (preRoot != nullptr) {
					//     cout<<preRoot->val<<endl;	//后根遍历输出根，两步输出
					// }
				}
	
				// Node *leftNode = stk.top();
				// cout<<stk.top()->val<<endl; 			//中/后根遍历输出左儿子，两步输出
				stk.pop();
	
				// if (preRoot != nullptr) {
				//     cout<<preRoot->val<<endl; 		//后根遍历输出根，两步输出
				// }
				if (!stk.empty() && stk.top()->right != nullptr) {
					root = stk.top()->right;
					// cout<<stk.top()->val<<endl; 		//中根遍历2，两步输出，不需要下面一行记录根节点
					preRoot = stk.top();
					stk.pop();
					// leftNode.right = stk.top();
					// stk.pop();
				}
			}
	
			return root;
		}
	};