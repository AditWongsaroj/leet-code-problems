
import unittest
from phoneCombos import Solution


class Test(unittest.TestCase):
    '''testing'''

    def tests(self):
        '''basic tests'''
        in_arrays = ["23", "", "2"]
        wants = [["ad", "ae", "af", "bd", "be",
                  "bf", "cd", "ce", "cf"], [], ["a", "b", "c"]]

        for i, s in enumerate(in_arrays):
            got = Solution().letterCombinations(s)
            want = wants[i]
            print(f'Test {i} - got: {got} | want: {want}')
            self.assertEqual(got, want)


if __name__ == '__main__':
    unittest.main()
