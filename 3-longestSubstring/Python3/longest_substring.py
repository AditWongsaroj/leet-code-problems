''' Given a string s, find the length of the 
    longest substring without repeating characters. '''

class Solution:
    '''from leetcode'''
    def length_of_longest_substring(self, s: str) -> int:
        '''renamed to conform to snakecase'''
        left, right, longest = 0, 0, 0
        ch_map = {}

        for i, c in enumerate(s):
            left = max(ch_map.get(c, 0),  left)
            longest = max(longest, right-left +1)
            ch_map[c] = i+1
            right +=1
        return longest
