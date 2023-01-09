#include <iostream>
#include <vector>
#include <queue>
#include <stack>

using namespace std;

/*

144. Binary Tree Preorder Traversal(Easy)

    Given the root of a binary tree, return the preorder traversal of its nodes' values.

    https://leetcode.com/problems/binary-tree-preorder-traversal/

*/


struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode* left, TreeNode* right) : val(x), left(left), right(right) {}
    
};

class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {

        vector<int> res;
        if (root == nullptr) return res;


        TreeNode* cur = root;
        stack<TreeNode*> s;
        s.push(cur);

        while (!s.empty()) {
            cur = s.top();
            s.pop();
            res.push_back(cur->val);

            if (cur->right != nullptr) s.push(cur->right);
            if (cur->left != nullptr) s.push(cur->left);
        }

        return res;
    }
};

int main()
{
    TreeNode* left = new TreeNode(1);
    TreeNode* right = new TreeNode(2);

    TreeNode* root = new TreeNode(3, left, right);

    Solution sol;
    vector<int> v = sol.preorderTraversal(root);

    for (int i = 0; i < v.size(); i++) {
        cout << v[i] << "\t";
    }
    cout << endl;
}
