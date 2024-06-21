import unittest

from lps import Solution

class Test(unittest.TestCase):
    def test(self):
        in1 = "babad"
        want = "bab"
        got = Solution().longestPalindrome(in1)
        print(f'Test {1} - got: {got} | want: {want}')
        self.assertEqual(got, want)

    def test2(self):
        in1 = "abbc"
        want = "bb"
        got = Solution().longestPalindrome(in1)
        print(f'Test {2} - got: {got} | want: {want}')
        self.assertEqual(got, want)

if __name__ == "__main__":
    unittest.main()