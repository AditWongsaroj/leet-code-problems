const { ListNode, addTwoNumbers } = require("./addTwoNumbers.js");

function listnodeconstructor(array) {
  let ln = new ListNode();
  let cur = ln;
  let alen = array.length;
  for (let i = 0; i < alen; i++) {
    cur.val = array[i];
    if (i != alen - 1) {
      cur.next = new ListNode();
      cur = cur.next;
    }
  }
  return ln;
}

test("zeros", () => {
  let l1 = listnodeconstructor([0]);
  let l2 = listnodeconstructor([0]);
  let result = listnodeconstructor([0]);

  expect(addTwoNumbers(l1, l2)).toStrictEqual(result);
});

test("simple carry", () => {
  let l1 = listnodeconstructor([2, 4, 3]);
  let l2 = listnodeconstructor([5, 6, 4]);
  let result = listnodeconstructor([7, 0, 8]);

  expect(addTwoNumbers(l1, l2)).toStrictEqual(result);
});

test("multi carry", () => {
  let l1 = listnodeconstructor([9, 9, 9]);
  let l2 = listnodeconstructor([9, 9, 9]);
  let result = listnodeconstructor([8, 9, 9, 1]);

  expect(addTwoNumbers(l1, l2)).toStrictEqual(result);
});

test("size difference, multi carry", () => {
  let l1 = listnodeconstructor([9, 9, 9, 9, 9, 9, 9]);
  let l2 = listnodeconstructor([9, 9, 9, 9]);
  let result = listnodeconstructor([8, 9, 9, 9, 0, 0, 0, 1]);

  expect(addTwoNumbers(l1, l2)).toStrictEqual(result);
});
