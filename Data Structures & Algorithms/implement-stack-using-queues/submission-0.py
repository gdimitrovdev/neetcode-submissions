class MyStack:

    def __init__(self):
        self.queue = []
        self.l = 0
        

    def push(self, x: int) -> None:
        self.queue.append(x)
        self.l += 1
        

    def pop(self) -> int:
        for i in range(self.l):
            popped = self.queue.pop(0)

            if i == self.l - 1:
                self.l -= 1
                return popped
            else:
                self.queue.append(popped)
        

    def top(self) -> int:
        for i in range(self.l - 1):
            popped = self.queue.pop(0)
            self.queue.append(popped)
        
        ans = self.queue[0]
        self.queue.pop(0)
        self.queue.append(ans)

        return ans


    def empty(self) -> bool:
        return self.l == 0
        


# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()