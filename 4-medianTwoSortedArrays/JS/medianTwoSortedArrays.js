/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function (nums1, nums2) {
  let size = nums1.length + nums2.length;
  let sorted = [...nums1, ...nums2].sort(function (a, b) {
    return a - b;
  });
  half = size / 2;

  if (Number.isInteger(half)) {
    return (sorted[half - 1] + sorted[half]) / 2;
  }
  return sorted[Math.floor(half)];
};

module.exports = findMedianSortedArrays;
