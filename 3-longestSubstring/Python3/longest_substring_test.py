''' Given a string s, find the length of the 
    longest substring without repeating characters. '''

import unittest
from longest_substring import Solution

class Test(unittest.TestCase):
    '''testing'''

    def tests(self):
        '''basic tests'''
        in_arrays = ["", "abc", "bb", "dvds", "abcabcbb", "abccbafa"]
        wants = [0, 3, 1, 3, 3, 4]
        for i, s in enumerate(in_arrays):
            got = Solution().length_of_longest_substring(s)
            want = wants[i]
            print(f'Test {i} - got: {got} | want: {want}')
            self.assertEqual(got, want)

if __name__ == '__main__':
    unittest.main()
