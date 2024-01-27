class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        ln = ListNode()
        cur = ln
        carry = 0

        while l1 != None or l2 != None or carry == 1:
            l1v = 0
            l2v = 0

            if(l1):
                l1v = l1.val
                l1 = l1.next
            if(l2):
                l2v = l2.val
                l2 = l2.next
            
            sum = l1v + l2v + carry
            carry = 0

            if sum > 9:
                carry = 1
                sum %= 10

            cur.val = sum

            if l1 != None or l2 != None or carry == 1:
                cur.next = ListNode()
                cur = cur.next    

        return ln


import unittest
class TestAddTwoNumbers(unittest.TestCase):
    def ln_cons (self, nums: [int])-> ListNode:
        ln = ListNode()
        cur = ln
        num_len = len(nums)
        for n in range(0, num_len):
            cur.val = nums[n]
            if(n != num_len-1):
                cur.next = ListNode()
                cur = cur.next
        return ln

# 
# 
    def test_zero_case(self):
        inArrays = [ [0], [0], [2,4,3], [5,6,4], [9,9,9], [9,9,9,9,9]  ]
        wants = [ [0], [2,4,3], [7,0,8], [4,6,4,1], [8,9,9,0,0,1] ]
        for i in range(0, len(inArrays)-1):
            l1 = self.ln_cons(inArrays[i])
            l2 = self.ln_cons(inArrays[i+1])
            want = self.ln_cons(wants[i])
            got = Solution().addTwoNumbers(l1,l2)
            while got != None or want != None:
                print(f'Test {i} - got: {got.val} | want: {want.val}')
                self.assertEqual(got.val, want.val)
                got = got.next
                want = want.next

if __name__ == '__main__':
    unittest.main()