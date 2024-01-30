'''Problem 2 https://leetcode.com/problems/add-two-numbers/'''

class ListNode:
    '''from leetcode; next line disables pylint squigle in vscode'''
    # pylint: disable-next=redefined-builtin
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    '''from leetcode'''
    def add_two_numbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        '''my attempt'''
        ln = ListNode()
        cur = ln
        carry = 0

        while l1 is not None or l2 is not None or carry == 1:
            l1v = 0
            l2v = 0

            if l1:
                l1v = l1.val
                l1 = l1.next
            if l2:
                l2v = l2.val
                l2 = l2.next

            s = l1v + l2v + carry
            carry = 0

            if s > 9:
                carry = 1
                s %= 10

            cur.val = s

            if l1 is not None or l2 is not None or carry == 1:
                cur.next = ListNode()
                cur = cur.next

        return ln
