import unittest
from reverseInteger import Solution


class Test(unittest.TestCase):
    '''testing'''

    def tests(self):
        '''basic tests'''
        in_arrays = [123, -123,
                     120]
        wants = [321, -321, 21]
        for i, s in enumerate(in_arrays):
            got = Solution().reverse(s)
            want = wants[i]
            print(f'Test {i} - got: {got} | want: {want}')
            self.assertEqual(got, want)


if __name__ == '__main__':
    unittest.main()
