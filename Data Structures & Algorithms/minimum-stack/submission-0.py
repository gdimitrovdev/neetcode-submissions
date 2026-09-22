from heapq import heappush, heappop

class MinStack:

    def __init__(self):
        self.stack = []
        self.heap = []
        self.counters = {}
        

    def push(self, val: int) -> None:
        self.stack.append(val)
        heappush(self.heap, val)

        if val not in self.counters:
            self.counters[val] = 0
        self.counters[val] += 1
        

    def pop(self) -> None:
        popped = self.stack.pop(-1)
        self.counters[popped] -= 1
        

    def top(self) -> int:
        return self.stack[-1]
        

    def getMin(self) -> int:
        smallest = self.heap[0]
        while self.counters[smallest] == 0:
            heappop(self.heap)
            smallest = self.heap[0]
        
        return smallest
