const twoSum = require("./twoSum.js");

test("zeros", () => {
  expect(twoSum([0, 0], 0)).toStrictEqual([0, 1]);
});
