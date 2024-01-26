function ListNode(val, next) {
  this.val = val === undefined ? 0 : val;
  this.next = next === undefined ? null : next;
}

var addTwoNumbers = function (l1, l2) {
  let result = new ListNode();
  let result_cur = result;
  let carry = 0;

  while (l1 || l2 || carry) {
    let sum = 0;
    let l1v = l1 ? l1.val : 0;
    let l2v = l2 ? l2.val : 0;

    sum = l1v + l2v + carry;
    carry = 0;
    if (sum > 9) {
      carry = 1;
      sum %= 10;
    }
    l1 = l1 ? l1.next : null;
    l2 = l2 ? l2.next : null;

    result_cur.val = sum;

    if (l1 || l2 || carry) {
      result_cur.next = new ListNode();
      result_cur = result_cur.next;
    }
  }
  return result;
};

module.exports = { ListNode, addTwoNumbers };
