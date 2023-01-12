#include <iostream>
using namespace std;

/*
21. Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

https://leetcode.com/problems/merge-two-sorted-lists/description/

*/

 struct ListNode {
      int val;
      ListNode *next;
      ListNode() : val(0), next(nullptr) {}
      ListNode(int x) : val(x), next(nullptr) {}
      ListNode(int x, ListNode *next) : val(x), next(next) {}
 };
 

 void Add(ListNode*& root, int val)
 {
     ListNode* n = new ListNode(val);

     if (root == nullptr) root = n;
     else if (root->next == nullptr) root->next = n;

     else {
         ListNode* cur = root;
         while (cur->next) {
             cur = cur->next;
         }

         cur->next = n;

     }

 }

 void Show(ListNode* root) {
     ListNode* cur = root;
     while (cur) {
         cout << cur->val << "\t";
         cur = cur->next;
     }
     cout << endl;
 }

class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {

        ListNode* res = NULL;

        ListNode* cur1 = list1, * cur2 = list2;

        for (;;) 
        {
            if (cur1 == nullptr && cur2 == nullptr) {
                break;
            }

            if (cur1 != nullptr && cur2 != nullptr) 
            {
                if (cur1->val <= cur2->val) {
                    Add(res, cur1->val);
                    cur1 = cur1->next;
                }
                else
                {
                    Add(res, cur2->val);
                    cur2 = cur2->next;
                }
            }
            else if (cur1 == nullptr && cur2 != nullptr) {
                Add(res, cur2->val);
                cur2 = cur2->next;
            }
            else if (cur1 != nullptr && cur2 == nullptr) {
                Add(res, cur1->val);
                cur1 = cur1->next;

            }
        }

        return res;
    }

};

int main()
{
    ListNode* list1 = nullptr, * list2 = nullptr;

    Add(list1, 1);
    Add(list1, 2);
    Add(list1, 4);

    Show(list1);
}
