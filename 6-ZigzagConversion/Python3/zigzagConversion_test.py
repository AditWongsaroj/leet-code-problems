import unittest
from zigzagConversion import Solution


class Test(unittest.TestCase):
    '''testing'''

    def tests(self):
        '''basic tests'''
        in_arrays = ["PAYPALISHIRING", "PAYPALISHIRING",
                     "A"]
        in_rows = [3, 4, 1]
        wants = ["PAHNAPLSIIGYIR", "PINALSIGYAHRPI", "A"]
        for i, s in enumerate(in_arrays):
            got = Solution().convert(s, in_rows[i])
            want = wants[i]
            print(f'Test {i} - got: {got} | want: {want}')
            self.assertEqual(got, want)


if __name__ == '__main__':
    unittest.main()
