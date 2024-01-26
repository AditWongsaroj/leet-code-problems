#include <assert.h>
#include <iostream>
#include <unordered_map>
#include <vector>
class Solution {
public:
    std::vector<int> twoSum(std::vector<int>& nums, int target) {        
        std::unordered_map<int, int> smap{};
        auto nsize = nums.size();
        for(auto i = 0; i < nsize; i++){
            auto x = target - nums[i];
            if(smap.find(x) !=smap.end()){
                return {smap[x], i};
            }
            smap.insert({nums[i], i});
        }
        return {0,0};
    }
};

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