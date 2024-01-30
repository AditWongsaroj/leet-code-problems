#define DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN
#include "doctest.h"

#include "twoSum.cpp"


TEST_CASE("testing the twoSum function") {
    std::vector<int> inArray = {2,7,12,4,-12,-13};
    std::vector<int> inTargets = {9,6,19,11,0,-25};
    std::vector<std::vector<int>> wants = {{0,1}, {0,3}, {1,2}, {1,3}, {2,4}, {4,5}};

    for(int i =0; i < inTargets.size(); i++){
        auto got = Solution().twoSum(inArray, inTargets[i]);
        
        CHECK(got == wants[i]);
    };
}

    