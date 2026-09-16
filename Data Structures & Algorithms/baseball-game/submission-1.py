class Solution:
    def calPoints(self, operations: List[str]) -> int:
        stack = []

        for op in operations:
            if op.isnumeric() or op[0] == "-" and op[1:].isnumeric():
                stack.append(int(op))
            elif op == "+":
                stack.append(stack[-1] + stack[-2])
            elif op == "D":
                stack.append(stack[-1] * 2)
            elif op == "C":
                stack.pop(-1)
            print(stack)

        return sum(stack)
        