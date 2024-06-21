
class Solution:
    rStone = {
        'I':  1,
        'V':  5,
        'X':  10,
        'L':  50,
        'C':  100,
        'D':  500,
        'M':  1000
    }

    def romanToInt(self, s: str) -> int:
        result = 0
        for i in range(len(s)):
            n = self.rStone[s[i]]
            next_n = 0
            if i < len(s)-1:
                next_n = self.rStone[s[i+1]]

            if next_n > n:
                result -= n
            else:
                result += n
        return result
