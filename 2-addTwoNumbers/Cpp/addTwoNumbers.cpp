#include <vector>

struct ListNode
{
  int val = 0;
  ListNode* next = nullptr;
  ListNode() = default;
  explicit ListNode(int x) : val(x){}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};

class Solution
{
public:
  ListNode* addTwoNumbers(ListNode* l1, ListNode* l2)
  {
    auto result = new ListNode();
    ListNode* cur = result;
    int carry = 0;

    while (l1 || l2 || carry)
    {
      int l1v = 0;
      int l2v = 0;

      if (l1)
      {
        l1v = l1->val;
        l1 = l1->next;
      }
      if (l2)
      {
        l2v = l2->val;
        l2 = l2->next;
      }

      int sum = l1v + l2v + carry;
      carry = 0;

      if (sum > 9)
      {
        carry = 1;
        sum %= 10;
      }

      cur->val = sum;

      if (l1 || l2 || carry)
      {
        cur->next = new ListNode();
        cur = cur->next;
      }

    }
    return result;
  }
};