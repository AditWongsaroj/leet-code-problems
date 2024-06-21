class Solution:
    def reverse(self, x: int) -> int:
        if abs(x) > 0x7fffffff or x == 0:
            return 0
        else:
            rev = str(x)[::-1].lstrip('0')
            if rev[-1] == '-':
                rev = int(rev[-1]+rev[:-1])
            else:
                rev = int(rev)
            if abs(rev) > 0x7fffffff:
                return 0
            return int(rev)
