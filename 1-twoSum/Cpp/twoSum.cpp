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