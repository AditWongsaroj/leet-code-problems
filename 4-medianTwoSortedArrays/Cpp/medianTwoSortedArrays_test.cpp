#define DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN
#include "doctest.h"

#include "medianTwoSortedArrays.cpp"

TEST_CASE("TESTING"){
    std::vector<std::vector<int>> inArrays = { {0}, {0}, {2,4,3}, {5,6,4}, {8}, {23,-10,-9} };
    std::vector<float> wants = { 0, 2.5, 4, 5.5, -0.5 };
  
    for (auto i = 0; i < wants.size(); i++)
    {
        auto a1 = inArrays[i];
        auto a2 = inArrays[i+1];
        auto want = wants[i];
        auto got = medianTwoSortedArrays(a1,a2);

        CHECK(got == want);
    }
}