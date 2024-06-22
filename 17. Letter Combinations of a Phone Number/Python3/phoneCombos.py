from itertools import product


class Solution:
    def letterCombinations(self, digits: str):
        if digits == "":
            return []
        keys = ['']*8
        a = ord('a')
        for x in range(8):
            if x == 7 or x == 5:
                n = 4
            else:
                n = 3
            for _ in range(n):
                keys[x] += chr(a)
                a += 1
        l = len(digits)
        args = [keys[int(_)-2] for _ in digits]

        x = list(product(*args))
        out = [''.join(_) for _ in x]
        return out
