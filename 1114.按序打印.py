#
# @lc app=leetcode.cn id=1114 lang=python3
#
# [1114] 按序打印
#
# https://leetcode-cn.com/problems/print-in-order/description/
#
# concurrency
# Easy (63.08%)
# Likes:    153
# Dislikes: 0
# Total Accepted:    35.6K
# Total Submissions: 56.3K
# Testcase Example:  '[1,2,3]'
#
# 我们提供了一个类：
#
#
# public class Foo {
# public void one() { print("one"); }
# public void two() { print("two"); }
# public void three() { print("three"); }
# }
#
#
# 三个不同的线程将会共用一个 Foo 实例。
#
#
# 线程 A 将会调用 one() 方法
# 线程 B 将会调用 two() 方法
# 线程 C 将会调用 three() 方法
#
#
# 请设计修改程序，以确保 two() 方法在 one() 方法之后被执行，three() 方法在 two() 方法之后被执行。
#
#
#
# 示例 1:
#
#
# 输入: [1,2,3]
# 输出: "onetwothree"
# 解释:
# 有三个线程会被异步启动。
# 输入 [1,2,3] 表示线程 A 将会调用 one() 方法，线程 B 将会调用 two() 方法，线程 C 将会调用 three() 方法。
# 正确的输出是 "onetwothree"。
#
#
# 示例 2:
#
#
# 输入: [1,3,2]
# 输出: "onetwothree"
# 解释:
# 输入 [1,3,2] 表示线程 A 将会调用 one() 方法，线程 B 将会调用 three() 方法，线程 C 将会调用 two() 方法。
# 正确的输出是 "onetwothree"。
#
#
#
# 注意:
#
# 尽管输入中的数字似乎暗示了顺序，但是我们并不保证线程在操作系统中的调度顺序。
#
# 你看到的输入格式主要是为了确保测试的全面性。
#
#

# @lc code=start
import threading
class Foo:
    def __init__(self):
        self.firstDone = threading.Lock()
        self.secondDone = threading.Lock()
        self.firstDone.acquire()
        self.secondDone.acquire()


    def first(self, printFirst: 'Callable[[], None]') -> None:

        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()
        self.firstDone.release()


    def second(self, printSecond: 'Callable[[], None]') -> None:

        # printSecond() outputs "second". Do not change or remove this line.
        with self.firstDone:
            printSecond()
            self.secondDone.release()


    def third(self, printThird: 'Callable[[], None]') -> None:

        # printThird() outputs "third". Do not change or remove this line.
        with self.secondDone:
            printThird()
# @lc code=end
