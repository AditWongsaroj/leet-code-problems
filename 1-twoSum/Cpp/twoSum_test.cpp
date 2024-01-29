#include <assert.h>
#include <iostream>

#include "twoSum.cpp"

//Testing
int main(){
    std::vector<int> inArray = {2,7,12,4};
    std::vector<int> inTargets = {9,6,19,11};
    std::vector<std::vector<int>> wants = {{0,1}, {0,3}, {1,2}, {1,3}};

    for(int i =0; i < inTargets.size(); i++){
        assert(Solution().twoSum(inArray, inTargets[i]) == wants[i]);
        std::cout << "Test " << i << " passed\n";
    };
}