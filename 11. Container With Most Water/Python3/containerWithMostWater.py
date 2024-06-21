class Solution:
    def maxArea(self, height) -> int:
        mxm, p, s, e = 0, 0, 0, len(height)-1
        while s < e:
            _sort = min(height[s], height[e])
            if p < _sort:
                mxm = max(mxm, _sort * (e-s))
                p = _sort
            if _sort == height[e]:
                e -= 1
            else:
                s += 1
        return mxm
