#include <iostream>
#include <vector>
#include <stack>

using namespace std;

/*
100. Same Tree(Easy)

    Given the roots of two binary trees p and q, write a function to check if they are the same or not.
    Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

    https://leetcode.com/problems/same-tree/
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
    bool isSameTree(TreeNode* p, TreeNode* q) {

        TreeNode* cur1 = p;
        TreeNode* cur2 = q;

        if (cur1 == nullptr && cur2 == nullptr) return true;

        stack<TreeNode*> s1;
        stack<TreeNode*> s2;

        s1.push(cur1);
        s2.push(cur2);

        while (!s1.empty() && !s2.empty()) {
            cur1 = s1.top();
            s1.pop();

            cur2 = s2.top();
            s2.pop();

            if (cur1 == nullptr && cur2 != nullptr || cur1 != nullptr && cur2 == nullptr) return false;

            if (cur1->left != nullptr && cur2->left == nullptr || cur1->left == nullptr && cur2->left != nullptr) return false;
            if(cur1->right != nullptr && cur2->right == nullptr || cur1->right == nullptr && cur2->right != nullptr) return false;

            if (cur1->left != nullptr && cur2->left != nullptr && (cur1->left->val != cur2->left->val)) return false;

            if (cur1->right != nullptr && cur2->right != nullptr && (cur1->right->val != cur2->right->val)) return false;

            if (cur1->val != cur2->val) return false;

            if (cur1->right != nullptr) s1.push(cur1->right);
            if (cur1->left != nullptr) s1.push(cur1->left);

            if (cur2->right != nullptr) s2.push(cur2->right);
            if (cur2->left != nullptr) s2.push(cur2->left);
        }

        if (!s1.empty() || !s2.empty()) return false;

        return true;
    }
};

int main()
{
    
}
