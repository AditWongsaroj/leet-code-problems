
import unittest
from intToRoman import Solution


class Test(unittest.TestCase):
    '''testing'''

    def tests(self):
        '''basic tests'''
        in_arrays = [3749, 58, 1994]
        wants = ["MMMDCCXLIX", "LVIII", "MCMXCIV"]

        for i, s in enumerate(in_arrays):
            got = Solution().intToRoman(s)
            want = wants[i]
            print(f'Test {i} - got: {got} | want: {want}')
            try:
                self.assertEqual(got, want)
            except:
                print(f'Test {i} - Failed')


if __name__ == '__main__':
    unittest.main()
