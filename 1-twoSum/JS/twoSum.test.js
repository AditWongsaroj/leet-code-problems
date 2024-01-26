const twoSum = require("./twoSum.js");

test("zeros", () => {
  expect(twoSum([0, 0], 0)).toStrictEqual([0, 1]);
});

test("more tests", () => {
  let inArray = [2, 7, 12, 4];
  let inTargets = [9, 6, 19, 11];
  let wants = [
    [0, 1],
    [0, 3],
    [1, 2],
    [1, 3],
  ];

  for (let i = 0; i < wants.length; i++) {
    expect(twoSum(inArray, inTargets[i])).toStrictEqual(wants[i]);
  }
});

test("negative tests", () => {
  let inArray = [9, -3, 0, -2, 100];
  let inTargets = [6, 9, -3, -5, 7];
  let wants = [
    [0, 1],
    [0, 2],
    [1, 2],
    [1, 3],
    [0, 3],
  ];

  for (let i = 0; i < wants.length; i++) {
    expect(twoSum(inArray, inTargets[i])).toStrictEqual(wants[i]);
  }
});
