/**
 * @param {string} s
 * @return {number}
 */
const lengthOfLongestSubstring = require("./longestSubstring");

function expectations(s, want) {
  let got = lengthOfLongestSubstring(s);

  expect(got).toStrictEqual(want);
  expect(got).not.toBeLessThan(want);
  expect(got).not.toBeGreaterThan(want);
  expect(got).not.toBe(null);
}

test("emtpy", () => {
  let str = "";
  let want = 0;

  expectations(str, want);
});

test("ones", () => {
  let str_arr = ["a", "bb", "ccc", "4444", "     ", "......"];
  let want = 1;

  for (str of str_arr) {
    expectations(str, want);
  }
});

test("twos", () => {
  let str_arr = ["ab", "bab", "bcc", "4434", " ]  ", ".....?"];
  let want = 2;

  for (str of str_arr) {
    expectations(str, want);
  }
});
