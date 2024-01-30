#define DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN
#include "doctest.h"

#include "addTwoNumbers.cpp"

TEST_CASE("TESTING"){
  std::vector<std::vector<int>> inArrays = { {0}, {0}, {2,4,3}, {5,6,4}, {9,9,9}, {9,9,9,9,9} };
  std::vector<std::vector<int>> wants = { {0}, {2,4,3}, {7,0,8}, {4,6,4,1}, {8,9,9,0,0,1} };
  
  for (auto i = 0; i < wants.size(); i++)
  {
    auto l1 = ln_builder(inArrays[i]);
    auto l2 = ln_builder(inArrays[i + 1]);
    auto want = ln_to_array(ln_builder(wants[i]));
    auto got = ln_to_array(Solution().addTwoNumbers(l1, l2));

      std::cout << "Test " << i << " - got: [ ";
      for (const auto& i: got)
        std::cout << i << ", ";
      std::cout << " ] | want: [ ";
      for (const auto& i : want)
        std::cout << i << ", ";
      std::cout << "]\n";
      CHECK(got == want);

    std::cout << "Tests " << i << " passed\n";
  }
}