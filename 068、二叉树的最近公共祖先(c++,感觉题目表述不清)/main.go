/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */


 int findP = 0;
 int findQ = 0;

int findNode(TreeNode* root, TreeNode* p, TreeNode* q) {

    if (root == p) {
        findP = 1;
    }
    if (root == q) {
        findQ = 2;
    }

    if (root == nullptr) {
        return 0;
    }
    findNode(root->left,p,q);
    findNode(root->right,p,q);

    return findP + findQ;

}

class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        
        if (root == nullptr || p == nullptr || q == nullptr) {
            return nullptr;
        }
        
        int l = 0;
        int r = 0;


        if (root == p || root == q) {
            return root;
        }else {

            //左右都有瓶p，q怎么办？？
            l = findNode(root->left,p,q);
            findQ = 0;
            findP = 0;
            if (l == 3) {
                return lowestCommonAncestor(root->left,p,q);
            }

            r = findNode(root->right,p,q);
            findQ = 0;
            findP = 0;
            if (r == 3) {
                return lowestCommonAncestor(root->right,p,q);
            }


        }

        return root;
    }
};


// class Solution {
// public:
//     TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {

//         if (root == nullptr || root == p || root == q) {
//             return root;
//         }
//         TreeNode* left =  lowestCommonAncestor(root->left,p,q);
//         TreeNode* right = lowestCommonAncestor(root->right,p,q);

//         //  没懂，查找的明明是两个，如果left,right都为null,不是也可能左右个一个的情况吗？为啥直接返回nullptr
//         //  而且，当left，right都不为空的时候，不应该进一步判断谁的深度更大吗？
//         if (!left && !right) return nullptr;
//         if (left == nullptr) return right;
//         if (right == nullptr) return left;

//         return root;
//     }
// };