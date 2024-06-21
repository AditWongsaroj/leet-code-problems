from itertools import cycle


class Solution:
    def convert(self, s: str, numRows: int) -> str:
        ordered = [str()] * numRows
        zig = cycle([*range(numRows), *range(numRows-2, 0, -1)])
        for _ in range(len(s)):
            ordered[next(zig)] += s[_]
        return ''.join(ordered)
