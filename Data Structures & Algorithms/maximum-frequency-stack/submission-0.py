class FreqStack:

    def __init__(self):
        self.freqs = []
        self.seen = {}
        

    def push(self, val: int) -> None:
        if val not in self.seen:
            self.seen[val] = -1

        self.seen[val] += 1

        while len(self.freqs) < self.seen[val] + 1:
            self.freqs.append([])

        self.freqs[self.seen[val]].append(val)
        

    def pop(self) -> int:
        ans = self.freqs[-1].pop()
        self.seen[ans] -= 1
        while len(self.freqs) > 0 and not self.freqs[-1]:
            self.freqs.pop()
        return ans
        


# Your FreqStack object will be instantiated and called as such:
# obj = FreqStack()
# obj.push(val)
# param_2 = obj.pop()