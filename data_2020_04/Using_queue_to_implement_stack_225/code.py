class MyStack:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.q = []


    def push(self, x: int) -> None:
        """
        Push element x onto stack.
        """
        self.q.append(x)
        q_len = len(self.q)
        while q_len > 1:
            self.q.append(self.q.pop(0))
            q_len -= 1

    def pop(self) -> int:
        """
        Removes the element on top of the stack and returns that element.
        """
        return self.q.pop(0)


    def top(self) -> int:
        """
        Get the top element.
        """
        return self.q[0]


    def empty(self) -> bool:
        """
        Returns whether the stack is empty.
        """
        return not bool(self.q)

list1 = ['Google', 'Runoob', 'Taobao']
list_pop=list1.pop(0)
print ("删除的项为 :", list_pop)
print ("列表现在为 : ", list1)