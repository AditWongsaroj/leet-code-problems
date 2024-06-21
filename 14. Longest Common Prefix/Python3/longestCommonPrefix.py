class Solution:
    def longestCommonPrefix(self, strs) -> str:
        pre = ''
        word = strs[0]
        for p in range(len(word)):
            letter = word[p]
            for w in strs[1:]:
                try:
                    if w[p] != letter:
                        return pre
                except:
                    return pre
            pre += letter
        return pre
