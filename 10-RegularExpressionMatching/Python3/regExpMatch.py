import re


class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        try:
            return re.match(p, s).span() == re.match(".*", s).span()
        except:
            return False
