#include <vector>

struct ListNode //from leetcode
{
  int val = 0;
  ListNode* next = nullptr;
  ListNode() = default;
  explicit ListNode(int x) : val(x){}
  ListNode(int x, ListNode* next) : val(x), next(next) {}
};


//helper methods for tests
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

std::vector<int> ln_to_array(ListNode* ln)
{
  auto arr = std::vector<int>{};
  while (ln != nullptr){
    arr.push_back(ln->val);
    ln = ln->next;
  }
  return arr;
}