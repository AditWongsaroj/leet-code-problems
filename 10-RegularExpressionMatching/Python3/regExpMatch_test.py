
import unittest
from regExpMatch import Solution


class Test(unittest.TestCase):
    '''testing'''

    def tests(self):
        '''basic tests'''
        in_arrays = ["aa", "aa", "ab"]
        in_p = ["a", "a*", ".*"]
        wants = [False, True, True]

        for i, s in enumerate(in_arrays):
            got = Solution().isMatch(s, in_p[i])
            want = wants[i]
            print(f'Test {i} - got: {got} | want: {want}')
            self.assertEqual(got, want)


if __name__ == '__main__':
    unittest.main()
