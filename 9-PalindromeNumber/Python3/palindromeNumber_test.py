
import unittest
from palindromeNumber import Solution


class Test(unittest.TestCase):
    '''testing'''

    def tests(self):
        '''basic tests'''
        in_arrays = [121, -121, 120]
        wants = [True, False, False]

        for i, s in enumerate(in_arrays):
            got = Solution().isPalindrome(s)
            want = wants[i]
            print(f'Test {i} - got: {got} | want: {want}')
            self.assertEqual(got, want)


if __name__ == '__main__':
    unittest.main()
