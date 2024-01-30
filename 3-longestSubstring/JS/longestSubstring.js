/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
  //   if (s === "") return 0;
  //   return 1;

  let left = 0;
  let mx = 0;
  const char_map = new Map();

  for (let right = 0; right < s.length; right++) {
    ch = s.at(right);
    lch = char_map.get(ch) ? char_map.get(ch) : 0;
    left = Math.max(lch, left);
    mx = Math.max(mx, right - left + 1);
    char_map.set(ch, right + 1);
  }

  return mx;
};

module.exports = lengthOfLongestSubstring;
