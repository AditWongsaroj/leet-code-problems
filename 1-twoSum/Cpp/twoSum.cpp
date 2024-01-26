#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {        
        unordered_map<int, int> smap{};
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