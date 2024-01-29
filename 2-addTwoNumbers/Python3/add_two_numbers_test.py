import unittest
from add_two_numbers import ListNode, Solution

class Test(unittest.TestCase):
    '''testing'''
    def ln_cons (self, nums: [int])-> ListNode:
        '''listnode from array'''
        ln = ListNode()
        cur = ln
        num_len = len(nums)
        for n in range(0, num_len):
            cur.val = nums[n]
            if n is not num_len-1:
                cur.next = ListNode()
                cur = cur.next
        return ln

    def ln_to_array(self, ln: ListNode)-> [int]:
        '''for testing'''
        arr = []
        while ln is not None:
            arr.append(ln.val)
            ln = ln.next
        return arr


    def tests(self):
        '''basic tests'''
        in_arrays = [ [0], [0], [2,4,3], [5,6,4], [9,9,9], [9,9,9,9,9]  ]
        wants = [ [0], [2,4,3], [7,0,8], [4,6,4,1], [8,9,9,0,0,1] ]
        for i in range(0, len(in_arrays)-1):
            l1 = self.ln_cons(in_arrays[i])
            l2 = self.ln_cons(in_arrays[i+1])
            want = self.ln_to_array(self.ln_cons(wants[i]))
            got = self.ln_to_array(Solution().add_two_numbers(l1,l2))

            print(f'Test {i} - got: {got} | want: {want}')
            self.assertEqual(got, want)

if __name__ == '__main__':
    unittest.main()
