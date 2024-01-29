#include <assert.h>
#include <iostream>
#include <vector>

#include "addTwoNumbers.cpp"

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

std::vector<int> ln_to_array(ListNode* ln){
  auto arr = std::vector<int>{};
  while (ln != nullptr){
    arr.push_back(ln->val);
    ln = ln->next;
  }
  return arr;
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
    auto want = ln_to_array(ln_builder(wants[i]));
    auto got = ln_to_array(Solution().addTwoNumbers(l1, l2));

    // while (got || want)
    // {
      std::cout << "Test " << i << " - got: [ ";
      for (const auto& i: got)
        std::cout << i << ", ";
      std::cout << " ] | want: [ ";
      for (const auto& i : want)
        std::cout << i << ", ";
      std::cout << "]\n";
      assert(got == want);
      // got = got->next;
      // want = want->next;
    // }
    std::cout << "Tests " << i << " passed\n";
  }
}