
import unittest
from longestCommonPrefix import Solution


class Test(unittest.TestCase):
    '''testing'''

    def tests(self):
        '''basic tests'''
        in_arrays = [["flower", "flow", "flight"],
                     ["dog", "racecar", "car"], ["tim", "time", "timid"]]
        wants = ["fl", "", "tim"]

        for i, s in enumerate(in_arrays):
            got = Solution().longestCommonPrefix(s)
            want = wants[i]
            print(f'Test {i} - got: {got} | want: {want}')
            self.assertEqual(got, want)


if __name__ == '__main__':
    unittest.main()
