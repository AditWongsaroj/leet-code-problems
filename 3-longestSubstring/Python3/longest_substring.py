''' Given a string s, find the length of the 
    longest substring without repeating characters. ''' 

class Solution:
    '''from leetcode'''
    def length_of_longest_substring(self, s: str) -> int:
        '''renamed to conform to snakecase'''
        left, right, longest = 0, 0, 0
        ch_map = {}

        for i, c in enumerate(s):
            x = ch_map.get(c) if ch_map.get(c) is not None else 0
            left = max(x,  left)
            longest = max(longest, right-left +1)
            ch_map[c] = i+1
            right +=1
        return longest

    def ms26(self, s: str) -> int:
        ''' faster from letcode'''     
        start = result = 0
        seen = {}
        for i, c in enumerate(s):
            if seen.get(c, -1) >= start:
                start = seen[c] + 1
            result = max(result, i - start + 1)
            seen[c] = i
        return result

if __name__ == '__main__':
    print(Solution.length_of_longest_substring(None, "abcabcbb"))
