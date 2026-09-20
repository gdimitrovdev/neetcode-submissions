class MyQueue:

    def __init__(self):
        self.stack = []
        

    def push(self, x: int) -> None:
        self.stack.append(x)
        

    def pop(self) -> int:
        stack_temp = []

        length = len(self.stack)
        for i in range(length - 1):
            stack_temp.append(self.stack.pop(-1))

        ans = self.stack.pop(-1)

        for i in range(length - 1):
            self.stack.append(stack_temp.pop(-1))

        return ans
        

    def peek(self) -> int:
        stack_temp = []

        length = len(self.stack)
        for i in range(length):
            stack_temp.append(self.stack.pop(-1))

        ans = stack_temp[-1]

        for i in range(length):
            self.stack.append(stack_temp.pop(-1))

        return ans
        

    def empty(self) -> bool:
        return len(self.stack) == 0
        


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()