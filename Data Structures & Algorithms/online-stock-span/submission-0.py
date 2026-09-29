class StockSpanner:

    def __init__(self):
        self.stack = []
        

    def next(self, price: int) -> int:
        counter = 1
        while len(self.stack) > 0 and self.stack[-1][0] <= price:
            _, span = self.stack.pop()
            counter += span
        self.stack.append((price, counter))

        return counter
        


# Your StockSpanner object will be instantiated and called as such:
# obj = StockSpanner()
# param_1 = obj.next(price)