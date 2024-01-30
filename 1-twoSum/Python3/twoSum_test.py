'''testing'''

import unittest
from twoSum import Solution

class TestTwoSum(unittest.TestCase):

    def test_zero_case(self):
        data = [0, 0]
        target = 0
        want = [0,1]
        got = Solution().twoSum(data, target)
        self.assertEqual(got, want)

    def test_multi_case(self):
        data = [2,7,12,4]
        targets = [9, 6, 19, 16]
        wants = [[0,1], [0,3], [1,2], [2,3]]
        for (target, want) in zip(targets, wants):
            got = Solution().twoSum(data, target)
            self.assertEqual(got, want)

if __name__ == '__main__':
    unittest.main()