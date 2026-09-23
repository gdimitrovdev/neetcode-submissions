class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stack = []

        for token in tokens:
            if token.isnumeric() or token[0] == "-" and token[1:].isnumeric():
                stack.append(int(token))
            elif token == "+":
                y = stack.pop(-1)
                x = stack.pop(-1)
                stack.append(x + y)
            elif token == "-":
                y = stack.pop(-1)
                x = stack.pop(-1)
                stack.append(x - y)
            elif token == "*":
                y = stack.pop(-1)
                x = stack.pop(-1)
                stack.append(x * y)
            elif token == "/":
                y = stack.pop(-1)
                x = stack.pop(-1)
                stack.append(int(x / y))

        return stack[0]
        