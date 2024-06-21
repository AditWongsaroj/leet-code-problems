def namer(n, ltr):
    if n == 9:
        return ltr[0]
    elif n > 4:
        return ltr[1] + ltr[2] * (n % 5)
    elif n == 4:
        return ltr[2] + ltr[1]
    else:
        return ltr[2] * n


class Solution:
    def intToRoman(self, num: int) -> str:
        ltrs = [["CM", "D", "C"], ["XC", "L", "X"], ["IX", "V", "I"]]
        rn = 'M' * int(num//1000)
        for i in range(3):
            try:
                rn += namer(num//(10**(2-i)) % 10, ltrs[i])
            except:
                pass
        return rn
