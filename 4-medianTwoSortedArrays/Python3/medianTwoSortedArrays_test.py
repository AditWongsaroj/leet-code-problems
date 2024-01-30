import unittest
from medianTwoSortedArrays import Solution

class Test(unittest.TestCase):
    def test0(self):
        in1 = [0]
        in2 = [1,2]
        want = 1
        got = Solution().findMedianSortedArrays(in1, in2)
        print(f'Test {1} - got: {got} | want: {want}')
        self.assertEqual(got, want)
    def test1(self):
        in1 = [0]
        in2 = [1]
        want = 0.5
        got = Solution().findMedianSortedArrays(in1, in2)
        print(f'Test {1} - got: {got} | want: {want}')
        self.assertEqual(got, want)
    def test2(self):
        in1 = [0]
        in2 = [1,2]
        want = 1
        got = Solution().findMedianSortedArrays(in1, in2)
        print(f'Test {1} - got: {got} | want: {want}')
        self.assertEqual(got, want)	
    def test3(self):
        in1 = [0, 2]
        in2 = [1,2]
        want = 1.5
        got = Solution().findMedianSortedArrays(in1, in2)
        print(f'Test {1} - got: {got} | want: {want}')
        self.assertEqual(got, want)
    def test4(self):
        in1 = [-1,2]
        in2 = [1,-4]
        want = 0.0
        got = Solution().findMedianSortedArrays(in1, in2)
        print(f'Test {1} - got: {got} | want: {want}')
        self.assertEqual(got, want)

if __name__ == '__main__':
    unittest.main()
