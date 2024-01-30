const findMedianSortedArrays = require("./medianTwoSortedArrays");

function expectations(n1, n2, want) {
  let got = findMedianSortedArrays(n1, n2);

  expect(got).toStrictEqual(want);
  expect(got).not.toBeLessThan(want);
  expect(got).not.toBeGreaterThan(want);
  expect(got).not.toBe(null);
}

test("empty", () => {
  let num = [0];
  let want = 0;

  expectations(num, num, want);
});

test("tests", () => {
  let n1 = [1, 2, 3];
  let n2 = [3, 4, 5];
  let want = 3.0;

  expectations(n1, n2, want);
});

test("tests", () => {
  let n1 = [];
  let n2 = [3, 4, 5];
  let want = 4.0;

  expectations(n1, n2, want);
});

test("tests", () => {
  let n1 = [1, 2];
  let n2 = [3, 4, 5];
  let want = 3.0;

  expectations(n1, n2, want);
});

test("tests", () => {
  let n1 = [1, 2, 2];
  let n2 = [3, 4, 5];
  let want = 2.5;

  expectations(n1, n2, want);
});
