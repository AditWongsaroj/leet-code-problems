import math
class Solution:
    # def longestPalindrome_old(self, s: str) -> str: 
    #     longest = ""
    #     for i, _ in enumerate(s):
    #         size = len(s)
    #         while size > i:
    #             cur_s = s[i:size]
    #             if self.isPalindrome(cur_s):
    #                 if len(cur_s) > len(longest):
    #                     longest = cur_s
    #             size -= 1
    #     return longest
    def longestPalindrome(self, s: str) -> str: 
        size = len(s)
        mid, oddity = divmod(size, 2)

        left_end = mid - 1
        right_start = mid + oddity

        for i in reversed(range(mid)):
            self.isPalindrome(s[:mid], s[mid:])

        for i, _ in enumerate(s):
            self.isPalindrome()


        return ""

    def isPalindrome(self, s1: str, s2: str) -> bool:
        return s1 == s2[::-1]

        # size = len(s)
        # mid = size/2
        # if size % 2 ==0:
        #     mid = int(mid)
        #     return s[0:mid] == s[mid:][::-1]
        # mid = math.floor(mid)
        # return  s[0:mid] == s[mid+1:][::-1]


