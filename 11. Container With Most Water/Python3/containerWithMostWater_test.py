
import unittest
from containerWithMostWater import Solution


class Test(unittest.TestCase):
    '''testing'''

    def tests(self):
        '''basic tests'''
        in_arrays = [[1, 8, 6, 2, 5, 4, 8, 3, 7], [1, 1]]
        wants = [49, 1]

        for i, s in enumerate(in_arrays):
            got = Solution().maxArea(s)
            want = wants[i]
            print(f'Test {i} - got: {got} | want: {want}')
            self.assertEqual(got, want)


if __name__ == '__main__':
    unittest.main()
