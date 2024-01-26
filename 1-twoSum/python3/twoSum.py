class Solution:
    def twoSum(self, nums:  [int], target: int) -> [int]:
        seen_dict = {}
        for i, num in enumerate(nums):
            x = target - num
            if x in seen_dict.keys():
                return [seen_dict[x], i]
            seen_dict[num] = i
        

import unittest        
class TestTwoSum(unittest.TestCase):

    def test_zero_case(self):
        data = [0, 0]
        target = 0
        result = Solution().twoSum(data, target)
        self.assertEqual(result, [0,1])

if __name__ == '__main__':
    unittest.main()