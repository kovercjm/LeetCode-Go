# Double pointers solution

- Make fast pointer `p` and slow pointer `q` start from `head`;
- `p` goes 2 steps a time and `q` goes 1 step a time;

## when no cycle

`p` will reaches the end of linked list

## when has a cycle

`p` and `q` will met in cycle, and now:

- if there are `a + b` nodes in the list, with `a` nodes before the entry of cycle and `b` nodes in the cycle;
- `p` goes `f` (stands for fast) steps, `q` goes `s` (stand for slow) steps, and we got `f = 2s`;
- when `p` met `q` in the cycle, `p` has gone several rounds in cycle more than `q`, so we got `f = s + nb`;
- from last two equations, and then got `s = nb`;
- as `q` has gone `nb` and goes `a` more steps can reach the entry of cycle, so let `p` starts from `head`, and goes 1 step a time together with `q`;
- when `p` and `q` meets again, they will reach the entry of cycle.

## reference

- https://leetcode.cn/problems/linked-list-cycle-ii/solution/linked-list-cycle-ii-kuai-man-zhi-zhen-shuang-zhi-/
- https://leetcode.cn/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode-solution/
