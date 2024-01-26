const { ListNode, addTwoNumbers } = require("./addTwoNumbers.js");

test("zeros", () => {
  let l1 = new ListNode(2, new ListNode(4, new ListNode(3)));
  let l2 = new ListNode(5, new ListNode(6, new ListNode(4)));
  let result = new ListNode(7, new ListNode(0, new ListNode(8)));

  expect(addTwoNumbers(l1, l2)).toStrictEqual(result);
});
