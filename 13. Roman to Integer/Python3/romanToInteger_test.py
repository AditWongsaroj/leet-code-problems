
import unittest
from romanToInteger import Solution


class Test(unittest.TestCase):
    '''testing'''

    def tests(self):
        '''basic tests'''
        in_arrays = ["III", "MMMDCCXLIX", "LVIII", "MCMXCIV"]
        wants = [3, 3749, 58, 1994]

        for i, s in enumerate(in_arrays):
            got = Solution().romanToInt(s)
            want = wants[i]
            print(f'Test {i} - got: {got} | want: {want}')
            self.assertEqual(got, want)


if __name__ == '__main__':
    unittest.main()
