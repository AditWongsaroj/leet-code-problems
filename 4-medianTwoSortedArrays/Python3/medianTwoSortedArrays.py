import math

class Solution:
    def findMedianSortedArrays(self, nums1: list[int], nums2: list[int]) -> float:
        ac = nums1 + nums2
        ac.sort()
        mid = len(ac)/2.0
        if not mid.is_integer():
            return ac[math.floor(mid)]
        return (ac[int(mid-1)] + ac[int(mid)])/2.0
