class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen_dict = {}
        for i, num in enumerate(nums):
            x = target - num
            if x in seen_dict.keys():
                return [seen_dict[x], i]
            seen_dict[num] = i
        
        