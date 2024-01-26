/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
  let seen_list = new Map();
  let n_len = nums.length;
  for (let i = 0; i < n_len; i++) {
    let x = target - nums[i];
    if (seen_list.has(x)) {
      return [seen_list.get(x), i];
    }
    seen_list.set(nums[i], i);
  }
};

module.exports = twoSum;
