#include <assert.h>
#include <iostream>
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

ListNode* ln_builder(std::vector<int> nums)
{
  auto ln = new ListNode();
  auto cur = ln;
  auto alen = nums.size();

  for (auto i = 0; i < alen; i++)
  {
    cur->val = nums[i];
    if (i != alen - 1)
    {
      cur->next = new ListNode();
      cur = cur->next;
    }
  }
  return ln;
}

//Testing
int main()
{
  std::vector<std::vector<int>> inArrays = { {0}, {0}, {2,4,3}, {5,6,4}, {9,9,9}, {9,9,9,9,9} };
  std::vector<std::vector<int>> wants = { {0}, {2,4,3}, {7,0,8}, {4,6,4,1}, {8,9,9,0,0,1} };

  for (auto i = 0; i < inArrays.size() - 1; i++)
  {
    auto l1 = ln_builder(inArrays[i]);
    auto l2 = ln_builder(inArrays[i + 1]);
    auto want = ln_builder(wants[i]);
    auto got = Solution().addTwoNumbers(l1, l2);

    while (got || want)
    {
      std::cout << "Test " << i << " - got: " << got->val << " | want: " << want->val << "\n";
      assert(got->val == want->val);
      got = got->next;
      want = want->next;
    }
    std::cout << "Test " << i << " passed\n";
  }
}